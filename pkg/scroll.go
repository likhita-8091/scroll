package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"log"
	"math/big"
	scroll_abi "scroll/abi"
	"scroll/abi/swap"
	"scroll/abi/usdc"
	"scroll/abi/weth"
)

const (
	UsdcCoin = "0xA0D71B9877f44C744546D649147E3F1e70a93760"
	WEthCoin = "0xa1EA0B2354F5A344110af2b6AD68e75545009a03"
)

type ScrollClient struct {
	ctx     context.Context
	Cli     *ethclient.Client
	chainID *big.Int

	USDCABI     *usdc.USDCABI
	TestUSDCABI *usdc.TestUsdcABI
	SwapABI     *swap.SwapABI
	WEthABI     *weth.WEthABI
	L2ABI       *scroll_abi.L2ABI
}

func NewScrollClient(ctx context.Context, url string) (*ScrollClient, error) {
	client, err := ethclient.DialContext(ctx, url)
	if err != nil {
		return nil, errors.New("连接scrol rpc 失败")
	}

	cli := &ScrollClient{
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

// 初始化
func (s *ScrollClient) init() error {
	// 设置chainID
	s.SetChainID()

	// 加载abi
	err := s.SetUSDCABI("0xeF71Ddc12Bac8A2ba0b9068b368189FFa2628942")
	if err != nil {
		return fmt.Errorf("new usdc abi error：%v\n", err.Error())
	}

	err = s.SetScrollABI("0x6d79Aa2e4Fbf80CF8543Ad97e294861853fb0649")
	if err != nil {
		return fmt.Errorf("new l2 abi error：%v\n", err.Error())
	}

	err = s.SetTestUsdcABI(UsdcCoin)
	if err != nil {
		return fmt.Errorf("new test usdc abi error：%v\n", err.Error())
	}

	err = s.SetSwapABI("0xD9880690bd717189cC3Fbe7B9020F27fae7Ac76F")
	if err != nil {
		return fmt.Errorf("new test usdc abi error：%v\n", err.Error())
	}

	err = s.SetWEthABI(WEthCoin)
	if err != nil {
		return fmt.Errorf("new test usdc abi error：%v\n", err.Error())
	}

	log.Println("load usdc abi ok")
	log.Println("load test usdc abi ok")
	log.Println("load l2 abi ok")
	log.Println("load swap abi ok")
	log.Println("load weth ok")
	return nil
}

func (s *ScrollClient) ChainID() *big.Int {
	if s.chainID == nil {
		s.SetChainID()
	}
	return s.chainID
}

func (s *ScrollClient) SetChainID() {
	id, err := s.Cli.ChainID(s.ctx)
	if err != nil {
		log.Fatal("获取scroll链ID失败", err.Error())
	}

	log.Println("scroll chain id", id)
	s.chainID = id
}

// SetUSDCABI 0xeF71Ddc12Bac8A2ba0b9068b368189FFa2628942
func (s *ScrollClient) SetUSDCABI(contractAddress string) error {
	to := common.HexToAddress(contractAddress)

	a, err := usdc.NewUSDCABI(to, s.Cli)
	if err != nil {
		return fmt.Errorf("load usdc abi error: %v\n", err.Error())
	}
	s.USDCABI = a
	return nil
}

// SetScrollABI 0x6d79Aa2e4Fbf80CF8543Ad97e294861853fb0649
func (s *ScrollClient) SetScrollABI(contractAddress string) error {
	to := common.HexToAddress(contractAddress)

	a, err := scroll_abi.NewL2ABI(to, s.Cli)
	if err != nil {
		return fmt.Errorf("load usdc abi error: %v\n", err.Error())
	}
	s.L2ABI = a
	return nil
}

// SetTestUsdcABI 0xA0D71B9877f44C744546D649147E3F1e70a93760
func (s *ScrollClient) SetTestUsdcABI(contractAddress string) error {
	to := common.HexToAddress(contractAddress)

	a, err := usdc.NewTestUsdcABI(to, s.Cli)
	if err != nil {
		return fmt.Errorf("load test usdc abi error: %v\n", err.Error())
	}
	s.TestUSDCABI = a
	return nil
}

// SetSwapABI 0xD9880690bd717189cC3Fbe7B9020F27fae7Ac76F
func (s *ScrollClient) SetSwapABI(contractAddress string) error {
	to := common.HexToAddress(contractAddress)

	a, err := swap.NewSwapABI(to, s.Cli)
	if err != nil {
		return fmt.Errorf("load test usdc abi error: %v\n", err.Error())
	}
	s.SwapABI = a
	return nil
}

// SetWEthABI 0xa1EA0B2354F5A344110af2b6AD68e75545009a03
func (s *ScrollClient) SetWEthABI(contractAddress string) error {
	to := common.HexToAddress(contractAddress)

	a, err := weth.NewWEthABI(to, s.Cli)
	if err != nil {
		return fmt.Errorf("load test usdc abi error: %v\n", err.Error())
	}
	s.WEthABI = a
	return nil
}
