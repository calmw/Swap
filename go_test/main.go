package main

import (
	"github.com/ethereum/go-ethereum/common"
	"uniswapv2-test/binding/uniswap"
	"uniswapv2-test/eth"
	mlog "uniswapv2-test/log"
)

func main() {
	cli := eth.Client()
	router, err := uniswap.NewRouter(common.HexToAddress(""), cli)
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return
	}
	//router.AddLiquidity()
}
