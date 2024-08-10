package main

import (
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"uniswapv2-test/binding/uniswap"
	"uniswapv2-test/eth"
	mlog "uniswapv2-test/log"
)

func main() {
	cli := eth.Client()
	//erc20, err := uniswap.NewErc20(common.HexToAddress("0xa19FD5253844b29986D7ac84867da3C0Eb8d3620"), cli)
	//if err != nil {
	//	mlog.Logger.Sugar().Error(err)
	//	return
	//}
	//err, auth := eth.GetAuth(cli)
	//if err != nil {
	//	mlog.Logger.Sugar().Error(err)
	//	return
	//}
	//big1e18 := big.NewInt(1e18)
	//amount := big1e18.Mul(big1e18, big.NewInt(10000))
	//res, err := erc20.Approve(auth, common.HexToAddress("0xE0e5Bd44D9075F21FBD0cddFB8A73AA34E861bde"), amount)
	//if err != nil {
	//	mlog.Logger.Sugar().Error(err)
	//	return
	//}
	//log.Println(res.Hash())
	exchange, err := uniswap.NewExchange(common.HexToAddress("0xE0e5Bd44D9075F21FBD0cddFB8A73AA34E861bde"), cli)
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return
	}

	//price, err := exchange.GetPrice(nil, big.NewInt(1), big.NewInt(1))
	//if err != nil {
	//	mlog.Logger.Sugar().Error(err)
	//	return
	//}
	//log.Println(price)

	//err, auth := eth.GetAuth(cli)
	//if err != nil {
	//	mlog.Logger.Sugar().Error(err)
	//	return
	//}
	//big1e18 := big.NewInt(1e18)
	//amount := big1e18.Mul(big1e18, big.NewInt(2))
	//liquidity, err := exchange.AddLiquidity(auth, amount)
	//if err != nil {
	//	mlog.Logger.Sugar().Error(err)
	//	return
	//}
	//log.Println(liquidity.Hash())
	tokenAmount, err := exchange.GetTokenAmount(nil, big.NewInt(1e18))
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return
	}
	// 994974874371859296
	// 180163785259326660  100000000000000000
	log.Println(tokenAmount)
}
