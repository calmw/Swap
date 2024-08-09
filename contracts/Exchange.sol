// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "./RoleAccess.sol";
import "./interface/IERC20.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Exchange is RoleAccess, ERC20 {
    event Log(string func, address sender, uint256 value, bytes data);

    address public tokenAddress; // 该交易所使用的代币

    // 交易所本身设置为LP代币
    // LP 代币基本上是发行给流动性提供者以换取其流动性的 ERC20 代币
    // 事实上，LP 代币是股票：
    // 您可以获得 LP 代币来换取您的流动性。
    // 您获得的代币数量与您在池储备中的流动性份额成正比。
    // 费用按您持有的代币数量按比例分配。
    // LP 代币可以兑换回流动性 + 累积费用。
    constructor(address _token) ERC20("Liquidity TOKEN", "LP") {
        _addAdmin(msg.sender);
        require(_token != address(0), "invalid token address");
        tokenAddress = _token;
    }

    // 添加流动性
    // 代币价格 Px=y/x ; Py=x/y ; Px和Py是ETH和token的价格，x和y是ETH和token储备量
    // LP， 每次流动性存入都会根据存入的以太币在以太币储备中的份额按比例发行LP 代币.
    // 发行数量 amountMinted = totalAmount * ( ethDeposited / ethReserve )
    // 添加流动性时，代币的数量是根据添加的ETH来计算需要添加的代币数量
    function addLiquidity(
        uint256 _tokenAmount
    ) public payable returns (uint256) {
        // 这是一个新的交易所（没有流动性），则当池子为空时允许任意流动性比例
        if (getReserve() == 0) {
            IERC20 token = IERC20(tokenAddress);
            token.transferFrom(msg.sender, address(this), _tokenAmount);
            // 在增加初始流动性时，发行的 LP 代币数量等于存入的以太币数量。
            uint256 liquidity = address(this).balance;
            _mint(msg.sender, liquidity);

            return liquidity;
        } else {
            // 在有一定流动性的情况下执行既定的储备比例
            // 这种情况下，我们不会存入用户提供的所有代币，而只会存入根据当前储备比率计算出的金额。要获得该金额，
            // 我们将比率 ( tokenReserve / ethReserve) 乘以存入的以太币数量。然后，如果用户存入的金额少于此金额，则会引发错误。
            // 当流动性添加到资金池时，这将保留价格
            uint256 ethReserve = address(this).balance - msg.value;
            uint256 tokenReserve = getReserve();
            // 添加前后不影响价格；由tokenAmount/msg.value =tokenReserve / ethReserve，可推出下面公式。
            // 保持价格不变，根据msg.value推出需要添加token的数量
            uint256 tokenAmount = (msg.value * tokenReserve) / ethReserve;
            require(_tokenAmount >= tokenAmount, "insufficient token amount");

            IERC20 token = IERC20(tokenAddress);
            token.transferFrom(msg.sender, address(this), tokenAmount);

            // 在增加初始流动性时，额外的流动性将根据存入的以太币数量按比例铸造 LP 代币：
            uint256 liquidity = (totalSupply() * msg.value) / ethReserve;
            _mint(msg.sender, liquidity);

            return liquidity;
        }
    }

    // 消除流动性
    // 为了消除流动性，我们可以再次使用 LP 代币：
    // 我们不需要记住每个流动性提供者存入的金额，并且可以根据 LP 代币份额计算移除的流动性数量
    // 消除流动性，返回减少的ETH和token数量
    function removeLiquidity(
        uint256 _amount
    ) public returns (uint256, uint256) {
        require(_amount > 0, "invalid amount");

        uint256 ethAmount = (address(this).balance * _amount) / totalSupply();
        uint256 tokenAmount = (getReserve() * _amount) / totalSupply();

        _burn(msg.sender, _amount);
        payable(msg.sender).transfer(ethAmount);
        IERC20(tokenAddress).transfer(msg.sender, tokenAmount);

        return (ethAmount, tokenAmount);
    }

    // 返回交易所代币余额
    function getReserve() public view returns (uint256) {
        return IERC20(tokenAddress).balanceOf(address(this));
    }

    // 获取价格，定价功能
    function getPrice(
        uint256 inputReserve,
        uint256 outputReserve
    ) public pure returns (uint256) {
        require(inputReserve > 0 && outputReserve > 0, "invalid reserves");

        return (inputReserve * 1000) / outputReserve;
    }

    // 从每次交换中收取 1% 的费用
    // amountWithFee = amount∗(100-fee)/100
    function getAmount(
        uint256 inputAmount,
        uint256 inputReserve,
        uint256 outputReserve
    ) private pure returns (uint256) {
        require(inputReserve > 0 && outputReserve > 0, "invalid reserves");

        uint256 inputAmountWithFee = inputAmount * 99;
        uint256 numerator = inputAmountWithFee * outputReserve;
        uint256 denominator = (inputReserve * 100) + inputAmountWithFee;

        return numerator / denominator;
    }

    function getTokenAmount(uint256 _ethSold) public view returns (uint256) {
        require(_ethSold > 0, "ethSold is too small");

        uint256 tokenReserve = getReserve();

        return getAmount(_ethSold, address(this).balance, tokenReserve);
    }

    function getEthAmount(uint256 _tokenSold) public view returns (uint256) {
        require(_tokenSold > 0, "tokenSold is too small");

        uint256 tokenReserve = getReserve();

        return getAmount(_tokenSold, tokenReserve, address(this).balance);
    }

    function ethToTokenSwap(uint256 _minTokens) public payable {
        uint256 tokenReserve = getReserve();
        uint256 tokensBought = getAmount(
            msg.value,
            address(this).balance - msg.value,
            tokenReserve
        );

        require(tokensBought >= _minTokens, "insufficient output amount");

        IERC20(tokenAddress).transfer(msg.sender, tokensBought);
    }

    function tokenToEthSwap(uint256 _tokensSold, uint256 _minEth) public {
        uint256 tokenReserve = getReserve();
        uint256 ethBought = getAmount(
            _tokensSold,
            tokenReserve,
            address(this).balance
        );

        require(ethBought >= _minEth, "insufficient output amount");

        IERC20(tokenAddress).transferFrom(
            msg.sender,
            address(this),
            _tokensSold
        );
        payable(msg.sender).transfer(ethBought);
    }

    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }
}
