package pkg

import (
	"context"
	"math/big"
	"testing"
)

const (
	ScrollAPI = "https://scroll-alphanet.public.blastapi.io"
	GoerliAPI = "https://rpc.ankr.com/eth_goerli"
	Password  = "jw"
)

func TestDeployContract(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("C:\\code\\hxy\\scroll\\keystore\\0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = account.DeployContract()
	if err != nil {
		t.Fatal(err)
	}
}

func TestUSDCClaim(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("/Users/jiangziya/code/github/scroll/keystore1/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = account.ClaimUSDC()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("claim usdc ok")
}

func TestAccount_GetLPTokenID(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	path := "/Users/jiangziya/.scroll/account/keystore/8/0xFbEC5fb1bd2919579930AFAD65e4361d1c9fd127"
	//path := "/Users/jiangziya/.scroll/account/keystore/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B"
	account, err := NewAccount(path, Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	list, err := account.GetLPTokenID()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("get tokenID ok", list)
}

func TestAccount_CancelLP(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	//path := "/Users/jiangziya/.scroll/account/keystore/8/0xFbEC5fb1bd2919579930AFAD65e4361d1c9fd127"
	path := "/Users/jiangziya/.scroll/account/keystore/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B"
	account, err := NewAccount(path, Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = account.CancelLP()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("cancel lp ok")
}

func TestAccount_ScrollWithdrawETH(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("/Users/jiangziya/.scroll/account/keystore3/1/0x7D3280482E9ecb264Fc29C45cE2c4b0d17F7fBfb", Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = account.ScrollWithdrawETH()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("withdraw eth ok")
}

func TestSwapUSDC2ETH(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("/Users/jiangziya/code/github/scroll/keystore1/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	// 2000usdc --> eth
	amount, b := new(big.Int).SetString("1000000000000000000000", 10)
	if !b {
		t.Fatal(err)
	}
	err = account.MultiCall(amount)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("swap ok")
}

func TestAccount_SwapEth2WEth(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("/Users/jiangziya/code/github/scroll/keystore1/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = account.SwapEth2WEth(big.NewInt(10000000000000000)) // 0.005 eth
	if err != nil {
		t.Fatal(err)
	}

	t.Log("swap ok")
}

func TestAccount_AddLP(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("C:\\Users\\jw199\\code\\jw\\scroll\\keystore\\0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli, nil)
	if err != nil {
		t.Fatal(err)
	}

	err = account.AddLP() // 0.005 eth
	if err != nil {
		t.Fatal(err)
	}

	t.Log("add uni lp ok")
}

func TestAccount_EthDepositScroll(t *testing.T) {
	cli, err := NewEthClient(context.Background(), GoerliAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("C:\\Users\\jw199\\code\\jw\\scroll\\keystore\\0x42e11b1f3D41295036E6cA1b0a98453837873Dc4", Password, nil, cli)
	if err != nil {
		t.Fatal(err)
	}

	err = account.EthDepositScroll()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("deposit ok")
}
