package main

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"uniswapv2-test/binding/uniswap"
	"uniswapv2-test/eth"
	mlog "uniswapv2-test/log"
)

func main() {
	cli := eth.Client()
	exchange, err := uniswap.NewExchange(common.HexToAddress("0x08cA183377e450D9660dA29c3B496B5c6F853308"), cli)
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return
	}

	err, auth := eth.GetAuth(cli)
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return
	}
	liquidity, err := exchange.AddLiquidity(auth, big.NewInt(2e18))
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return
	}
}
