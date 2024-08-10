// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.18;

import "./RoleAccess.sol";
import "./interface/IERC20.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

/// V1版本仅允许ETH和token之间的兑换，可通过链式兑换来实现token之间的兑换
/// Uniswap底层是一种算法，可以创建池子，或者代币对，并向其中填充流动性，让用户利用这些流动性来交换代币。这种算法被称为自动做市商或自动流动性提供者
/// DEX 必须拥有足够（或大量）的流动性才能正常运作，并成为中心化交易所的替代品
/// 比起让开发者或投资者成为做市商，更好的解决方案是允许任何人成为做市商，这就是 Uniswap 成为自动化做市商的原因：任何用户都可以将他们的资金存入交易对（并从中受益）
/// Uniswap 仍然可以充当价格预言机
/// Uniswap 充当二级市场，吸引套利者，他们利用 Uniswap 和中心化交易所之间的价格差异获利。这使得 Uniswap 池中的价格尽可能接近大型交易所的价格
/// Uniswap 的核心是恒定乘积函数：
///    x∗y=k  // x是ETH储备，y是token储备，这是一个基础公式，实际用的是在这个基础上发展而来的 (x+Δx)(y−Δy)=xy，所以我们用来计算价格的常数乘积公式实际上是一条双曲线，此时理论上x,y永远不为0，这使得储备无限
/// 价格函数会导致价格滑点。相对于储备量，交易的代币数量越大，价格就越高。

contract Exchange is RoleAccess, ERC20 {
    event Log(string func, address sender, uint256 value, bytes data);

    address public tokenAddress; // 该交易所使用的代币，由于每个交易所只允许一种代币的兑换，我们需要用代币地址连接交易所；这里每个exchange合约就相当于一种代币的交易所

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
    // 我们需要msg.value从合约余额中减去，因为在调用该函数时，发送的以太币已经添加到其余额中。 ？？
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

    // 获取价格(ETH/token兑token/ETH的价格)，定价功能；也就是每个ETH兑几个token或者每个token兑几个ETH
    // 价格函数会导致价格滑点。相对于储备量，交易的代币数量越大，价格就越高。
    // 价格大致是储备量的关系，真实可成交的数量，后面的方法获取可实际兑换的数量（会用新的公式并且扣除交易费用）
    // 我们得到的数量比预期根据此价格计算的要少一些。这可能被视为恒定产品做市商的缺点（因为每笔交易都有滑点），但这是保护资金池不被耗尽的相同机制
    function getPrice(
        uint256 inputReserve, // ETH储备量/token储备量
        uint256 outputReserve // token储备量/ETH储备量
    ) public pure returns (uint256) {
        require(inputReserve > 0 && outputReserve > 0, "invalid reserves");

//        return (inputReserve * 1000) / outputReserve; // 正常该使用这种计算逻辑，但是solidity不支持浮点数，就用下面的方式来提高精度，在使用的时候，记得放大了1000倍
        return (inputReserve * 1000) / outputReserve;
    }

    // 根据数量交易，计算数量
    // 从每次交换中收取 1% 的费用
    // amountWithFee = amount∗(100-fee)/100
    // (x+Δx)(y−Δy)=xy ;该公式是在xy=k的基础上发展而来，xy=k是直线，在与xy轴相交时，有一个值会为0（比如全部兑换走ETH或者token），这样交易所就没有了流动性。x是ETH储备量，y是token储备量,Δx是我们交易的ETH/token的数量，Δy是我们获得的token/ETH的数量
    //            yΔx
    //   Δy = -----------
    //            x+Δx
    // 根据上面公式计算出可兑换ETH或者token的数量，因为要扣除1%费用，所以下面在计算的时候分子分母同时乘以100（因为扣除1%所以交易的ETH或token数量只乘了99）
    function getAmount(
        uint256 inputAmount, // 要交易的ETH/token的数量,相当于Δx或Δy
        uint256 inputReserve, // ETH/token的储备量,相当于x或y
        uint256 outputReserve // 可兑换成 token/ETH的储备量,相当于y或x
    ) private pure returns (uint256) { // 返回可兑换成token/ETH的数量
        require(inputReserve > 0 && outputReserve > 0, "invalid reserves");
        // (inputAmount + inputReserve)(returnValue + outputReserve) = inputReserve * outputReserve
        uint256 inputAmountWithFee = inputAmount * 99;
        uint256 numerator = inputAmountWithFee * outputReserve;
        uint256 denominator = (inputReserve * 100) + inputAmountWithFee;

        return numerator / denominator;
    }

    // 计算ETH可兑换token的数量
    function getTokenAmount(uint256 _ethSold) public view returns (uint256) {
        require(_ethSold > 0, "ethSold is too small");

        uint256 tokenReserve = getReserve();

        return getAmount(_ethSold, address(this).balance, tokenReserve);
    }

    // 计算token可兑换ETH的数量
    function getEthAmount(uint256 _tokenSold) public view returns (uint256) {
        require(_tokenSold > 0, "tokenSold is too small");

        uint256 tokenReserve = getReserve();

        return getAmount(_tokenSold, tokenReserve, address(this).balance);
    }

    // ETH兑换token
    // _minTokens，这是用户希望用以太币换取的最小代币数量。此数量包括滑点容忍度；用户同意至少获得这么多，但不能更少。这是一个非常重要的机制，
    // 可以保护用户免受抢先交易机器人的攻击，这些机器人试图拦截他们的交易并修改池余额以获取利润。
    // http://35.78.192.225:4000/tx/0x62c29700cc390669159037747070c168a99e1ef3fa014731e51019f904aa4a4e
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

    // token兑换ETH
    // _minEth，这是用户希望用token换取的最小ETH数量。此数量包括滑点容忍度；用户同意至少获得这么多，但不能更少。这是一个非常重要的机制，
    // 可以保护用户免受抢先交易机器人的攻击，这些机器人试图拦截他们的交易并修改池余额以获取利润。
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
