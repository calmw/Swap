package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	erc20 "uniswapv2-test/binding/uniswap"
	mlog "uniswapv2-test/log"
)

func Client() *ethclient.Client {
	client, err := ethclient.Dial("http://35.78.192.225:8545")
	if err != nil {
		panic(err)
	}
	return client
}

func GetAuth(cli *ethclient.Client) (error, *bind.TransactOpts) {

	key := os.Getenv("META_ACCOUNT")
	privateKeyEcdsa, err := crypto.HexToECDSA(key)
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return err, nil
	}
	publicKey := crypto.PubkeyToAddress(privateKeyEcdsa.PublicKey)
	nonce, err := cli.PendingNonceAt(context.Background(), publicKey)
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return err, nil
	}
	gasPrice, err := cli.SuggestGasPrice(context.Background())
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return err, nil
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(698))
	if err != nil {
		mlog.Logger.Sugar().Error(err)
		return err, nil
	}

	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     big.NewInt(int64(nonce)),
		Value:     big.NewInt(1e18),
		GasPrice:  gasPrice,
		Signer:    auth.Signer,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false,
	}
}

func GetTxByHash(hash string) {
	cli := Client()
	byHash, b, err := cli.TransactionByHash(context.Background(), common.HexToHash(hash))
	json, _ := byHash.MarshalJSON()

	fmt.Println(string(json), b, err)
}

func Transfer() {
	cli := Client()
	erc20, err := erc20.NewErc20(common.HexToAddress("0x816644F8bc4633D268842628EB10ffC0AdcB6099"), cli)
	if err != nil {
		log.Fatalln(err)
	}
	err, auth := GetAuth(cli)
	if err != nil {
		log.Fatalln(err)
		return
	}
	transfer, err := erc20.Transfer(auth, common.HexToAddress(""), big.NewInt(1))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(transfer.Hash(), err)
}
