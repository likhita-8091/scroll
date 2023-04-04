/*
----------------------------------------------
 File Name: eth.go
 Author: jw199@firecloud.ai
 @Company:  Firecloud
 Created Time: 2023/4/4 12:49
-------------------功能说明-------------------
 $
----------------------------------------------
*/

package pkg

import (
	"context"
	"errors"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"log"
	"math/big"
	scroll_abi "scroll/abi"
)

type EthClient struct {
	ctx     context.Context
	Cli     *ethclient.Client
	chainID *big.Int

	ScrollABI *scroll_abi.ScrollABI
}

func NewEthClient(ctx context.Context, url string) (*EthClient, error) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, errors.New("连接eth rpc 失败")
	}

	cli := &EthClient{
		ctx: ctx,
		Cli: client,
	}

	log.Println("rpc client connect ok")

	err = cli.init()
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (e *EthClient) init() error {
	// 设置chainID
	e.SetChainID()

	// 加载abi
	err := e.SetScrollABI("0xe5E30E7c24e4dFcb281A682562E53154C15D3332")
	if err != nil {
		return err
	}

	return nil
}

func (e *EthClient) SetScrollABI(address string) error {

	var err error
	to := common.HexToAddress(address)

	abi, err := scroll_abi.NewScrollABI(to, e.Cli)
	if err != nil {
		return err
	}
	e.ScrollABI = abi
	return nil
}

func (e *EthClient) SetChainID() {
	id, err := e.Cli.ChainID(e.ctx)
	if err != nil {
		log.Fatal("获取eth链ID失败", err.Error())
	}

	log.Println("eth chain id", id)
	e.chainID = id
}

func (e *EthClient) ChainID() *big.Int {
	return e.chainID
}
