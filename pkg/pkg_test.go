package pkg

import (
	"context"
	"math/big"
	"testing"
)

const (
	ScrollAPI = "https://alpha-rpc.scroll.io/l2"
	Password  = "jw"
)

func TestUSDCClaim(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("/Users/jiangziya/code/github/scroll/keystore1/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli)
	if err != nil {
		t.Fatal(err)
	}

	err = account.ClaimUSDC()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("claim usdc ok")
}

func TestSwapUSDC2ETH(t *testing.T) {
	cli, err := NewScrollClient(context.Background(), ScrollAPI)
	if err != nil {
		t.Fatal(err)
	}

	account, err := NewAccount("/Users/jiangziya/code/github/scroll/keystore1/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli)
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

	account, err := NewAccount("/Users/jiangziya/code/github/scroll/keystore1/0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", Password, cli)
	if err != nil {
		t.Fatal(err)
	}

	err = account.SwapEth2WEth(big.NewInt(10000000000000000)) // 0.005 eth
	if err != nil {
		t.Fatal(err)
	}

	t.Log("swap ok")
}
