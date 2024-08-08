#### 合约

- 用户跟router合约交互即可
- router 合约地址： 0xD736c91d85755b503Ac1Ceb79867Dd73C177b696
- TOXA： 0x65FCEc1a5A5E803877C788b494d9adF2a955e95d
- TOXB： 0xc787ac4A98B974cdB3d4a2D5E1ca31802f9063EA

#### 添加流动性
```shell
function addLiquidity(
        address tokenA,
        address tokenB,
        uint amountADesired,
        uint amountBDesired,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) external virtual override ensure(deadline) returns (uint amountA, uint amountB, uint liquidity) {}
      
```
#### 移除流动性
```shell
    function removeLiquidity(
        address tokenA,
        address tokenB,
        uint liquidity,
        uint amountAMin,
        uint amountBMin,
        address to,
        uint deadline
    ) public virtual override ensure(deadline) returns (uint amountA, uint amountB) {}
```
#### 兑换

```shell
# 想要转入的代币数量、最少转出的数量、路径、接收者、交易失效时间
function swapExactTokensForTokens(
        uint amountIn,
        uint amountOutMin,
        address[] calldata path,
        address to,
        uint deadline
    ) external virtual override ensure(deadline) returns (uint[] memory amounts) {}
```