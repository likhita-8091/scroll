package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/accounts/keystore"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
	"os"
	"scroll/abi/swap"
	"strings"
	"time"
)

type Account struct {
	keyJson  string // keystore文件，包含公钥、私钥
	password string // keystore文件的密码
	key      *keystore.Key

	scrollCli *ScrollClient
}

// NewAccount new账户
func NewAccount(addressPath string, password string, scrollCli *ScrollClient) (*Account, error) {
	// 设置cli
	if scrollCli == nil {
		return nil, errors.New("scroll client is nil")
	}

	// 加载keystore
	keyJson, err := os.ReadFile(addressPath)
	if err != nil {
		return nil, fmt.Errorf("Could not read key file: %v\n", err)
	}

	key, err := keystore.DecryptKey(keyJson, password)
	if err != nil {
		return nil, err
	}

	return &Account{
		keyJson:   string(keyJson),
		password:  password,
		key:       key,
		scrollCli: scrollCli,
	}, nil
}

func (a *Account) GetScrollCli() *ethclient.Client {
	return a.scrollCli.Cli
}

func (a *Account) GetScrollChainID() *big.Int {
	return a.scrollCli.ChainID()
}

// makeScrollChainTxOpts 构造交易选项
func (a *Account) makeScrollChainTxOpts() (*bind.TransactOpts, error) {
	opts, err := bind.NewTransactorWithChainID(strings.NewReader(a.keyJson), a.password, a.scrollCli.ChainID())
	if err != nil {
		return nil, err
	}

	// 设置交易发起人
	opts.From = a.Address()
	return opts, nil
}

// PrintTx 打印交易数据
func (a *Account) PrintTx(tx *types.Transaction, msg string) {
	log.Printf("%v [%v] \n", msg, a.Address().String())
	log.Printf("交易哈希: %v\n", tx.Hash().String())
}

// PrintReceipt 打印交易收据
func (a *Account) PrintReceipt(receipt *types.Receipt) {
	msg := ""
	switch receipt.Status {
	case types.ReceiptStatusFailed:
		msg = "失败"
	case types.ReceiptStatusSuccessful:
		msg = "成功"
	}

	j, _ := receipt.MarshalJSON()
	fmt.Println(string(j))

	if len(receipt.ReturnValue) != 0 {
		log.Printf("[%v] 交易结果(%v): %v\n", a.Address().String(), msg, receipt.ReturnValue)
	} else {
		log.Printf("[%v] 交易结果(%v)\n", a.Address().String(), msg)
	}
}

// waitMined 等待交易被打包区块并挖掘
func (a *Account) waitMined(tx *types.Transaction) (receipt *types.Receipt, err error) {
	return bind.WaitMined(a.scrollCli.ctx, a.scrollCli.Cli, tx)
}

func (a *Account) Ctx() context.Context {
	return a.scrollCli.ctx
}

func (a *Account) Address() common.Address {
	return a.key.Address
}

// GetGasPrice 获取gas价格
func (a *Account) GetGasPrice() (*big.Int, error) {
	gasPrice, err := a.scrollCli.Cli.SuggestGasPrice(a.Ctx())
	if err != nil {
		return nil, fmt.Errorf("get gas price error: %v\n", err.Error())
	}

	log.Println("gas price ok ", gasPrice.String())
	return gasPrice, nil
}

// GetNonce 获取nonce
func (a *Account) GetNonce() (*big.Int, error) {

	time.Sleep(1200 * time.Millisecond)

	nonce, err := a.scrollCli.Cli.PendingNonceAt(a.Ctx(), a.Address())
	if err != nil {
		return nil, fmt.Errorf("get nonce error: %v\n", err.Error())
	}

	log.Println("get nonce ok  ", nonce)
	return new(big.Int).SetUint64(nonce), nil
}

// ClaimUSDC 水龙头获取usdc
func (a *Account) ClaimUSDC() error {
	opts, err := a.makeScrollChainTxOpts()
	if err != nil {
		return fmt.Errorf("make usdc tx opts error：%v", err.Error())
	}

	nonce, err := a.GetNonce()
	if err != nil {
		return err
	}

	price, err := a.GetGasPrice()
	if err != nil {
		return err
	}

	opts.GasLimit = 100000
	opts.Nonce = nonce
	opts.GasPrice = price

	txRes, err := a.scrollCli.USDCABI.Claim(opts)
	if err != nil {
		return fmt.Errorf("claim usdc error: %v\n", err.Error())
	}

	a.PrintTx(txRes, "claim usdc 交易数据: ")

	receipt, err := a.waitMined(txRes)
	if err != nil {
		return fmt.Errorf("wait usdc claim error: %v\n", err.Error())
	}

	a.PrintReceipt(receipt)
	log.Println("claim usdc 交易完成")

	time.Sleep(3 * time.Second)

	return nil
}

// ApproveUSDC 批准usdc
func (a *Account) ApproveUSDC(amount *big.Int) error {

	opts, err := a.makeScrollChainTxOpts()
	if err != nil {
		return err
	}

	nonce, err := a.GetNonce()
	if err != nil {
		return err
	}

	price, err := a.GetGasPrice()
	if err != nil {
		return err
	}

	opts.GasLimit = 250000
	opts.Nonce = nonce
	opts.GasPrice = price
	approve, err := a.scrollCli.TestUSDCABI.Approve(opts, common.HexToAddress("0xd9880690bd717189cc3fbe7b9020f27fae7ac76f"), amount)
	if err != nil {
		return err
	}

	a.PrintTx(approve, "approve usdc")

	mined, err := a.waitMined(approve)
	if err != nil {
		return err
	}

	a.PrintReceipt(mined)
	log.Println("approve usdc ok")

	return nil
}

func (a *Account) SwapEth2WEth(amount *big.Int) error {
	opts, err := a.makeScrollChainTxOpts()
	if err != nil {
		return err
	}

	nonce, err := a.GetNonce()
	if err != nil {
		return err
	}

	price, err := a.GetGasPrice()
	if err != nil {
		return err
	}

	opts.GasLimit = 80000
	opts.Nonce = nonce
	opts.GasPrice = price
	opts.Value = amount

	deposit, err := a.scrollCli.WEthABI.Deposit(opts)
	if err != nil {
		return err
	}

	a.PrintTx(deposit, "swap eth to weth")

	mined, err := a.waitMined(deposit)
	if err != nil {
		return err
	}

	a.PrintReceipt(mined)
	log.Println("swap eth to weth ok")
	return nil
}

func (a *Account) SwapUSDC2ETH(amount *big.Int) ([]byte, error) {
	param := swap.IV3SwapRouterExactInputSingleParams{
		TokenIn:           common.HexToAddress(UsdcCoin),
		TokenOut:          common.HexToAddress(WEthCoin),
		Fee:               big.NewInt(500),
		Recipient:         a.Address(),
		AmountIn:          amount,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	}

	abi, _ := swap.SwapABIMetaData.GetAbi()
	return abi.Pack("exactInputSingle", param)
}

func (a *Account) MultiCall(amount *big.Int) error {
	//err := a.ApproveUSDC(amount)
	//if err != nil {
	//	return fmt.Errorf("approve usdc error: %v\n", err.Error())
	//}

	param, err := a.SwapUSDC2ETH(amount)
	if err != nil {
		return err
	}

	log.Println("构造swap交易param成功")

	opts, err := a.makeScrollChainTxOpts()
	if err != nil {
		return err
	}

	nonce, err := a.GetNonce()
	if err != nil {
		return err
	}

	//price, err := a.GetGasPrice()
	//if err != nil {
	//	return err
	//}

	opts.GasLimit = 150000
	opts.Nonce = nonce
	//opts.Value = big.NewInt(10000000000)
	//opts.GasPrice = price

	deadLine := time.Now().Add(5 * time.Minute).Unix()

	txData, err := a.scrollCli.SwapABI.Multicall0(opts, big.NewInt(deadLine), [][]byte{param})
	if err != nil {
		return err
	}

	a.PrintTx(txData, "swap usdc to weth")

	mined, err := a.waitMined(txData)
	if err != nil {
		return err
	}

	a.PrintReceipt(mined)
	time.Sleep(2 * time.Second)
	log.Println("swap usdc to weth ok")
	return nil
}

// SendScroll scroll账户转账scroll到scroll账户
func (a *Account) SendScroll(toAddress string) error {
	balance, err := a.GetScrollBalance()
	if err != nil {
		return err
	}

	var a1 *big.Int
	var ok bool

	// 循环五次
	for i := 0; i < 5; i++ {
		if i == 0 {
			a1 = new(big.Int).Mul(balance, big.NewInt(99))
		} else if i == 1 {
			a1 = new(big.Int).Mul(balance, big.NewInt(97))
		} else if i == 2 {
			a1 = new(big.Int).Mul(balance, big.NewInt(96))
		} else if i == 3 {
			a1 = new(big.Int).Mul(balance, big.NewInt(95))
		} else if i == 4 {
			a1 = new(big.Int).Mul(balance, big.NewInt(94))
		}

		b1 := new(big.Int).Div(a1, big.NewInt(100))

		nonce, err := a.GetNonce()
		if err != nil {
			continue
		}

		price, err := a.GetGasPrice()
		if err != nil {
			continue
		}

		to := common.HexToAddress(toAddress)

		// 构造交易
		baseTx := &types.LegacyTx{
			Nonce:    nonce.Uint64(),
			GasPrice: price,
			Gas:      uint64(21000),
			To:       &to,
			Value:    b1,
		}
		tx := types.NewTx(baseTx)

		// 交易签名
		signer := types.LatestSignerForChainID(a.GetScrollChainID())
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), a.key.PrivateKey)
		if err != nil {
			log.Println("生成signature失败", err.Error())
			continue
		}
		signTx, err := tx.WithSignature(signer, signature)
		if err != nil {
			log.Println("签名失败", err.Error())
			continue
		}

		// 发送交易
		err = a.GetScrollCli().SendTransaction(a.Ctx(), signTx)
		if err != nil {
			log.Println("send signTx error", err.Error())
			continue
		}

		a.PrintTx(signTx, "交易发送成功")

		mined, err := a.waitMined(signTx)
		if err != nil {
			continue
		}

		a.PrintReceipt(mined)

		if mined.Status == types.ReceiptStatusSuccessful {
			ok = true
			break
		}
	}

	if ok {
		msg := fmt.Sprintf("========[%v] 账户交易[%v]成功========", a.Address(), toAddress)
		log.Println(msg)
		return nil
	} else {
		msg := fmt.Sprintf("========[%v]账户[%v]失败========", a.Address(), toAddress)
		log.Println(msg)
		return errors.New(msg)
	}
}

// GetScrollBalance 获取scroll账户余额
func (a *Account) GetScrollBalance() (*big.Int, error) {

	balance, err := a.GetScrollCli().BalanceAt(a.Ctx(), a.Address(), nil)
	if err != nil {
		return nil, fmt.Errorf("[%v] 获取余额失败: %v", a.Address().String(), err.Error())
	}

	newBalance, _ := new(big.Int).SetString(balance.String(), 10)
	b1 := new(big.Float)
	b1.SetString(newBalance.String())
	ethValue := new(big.Float).Quo(b1, big.NewFloat(math.Pow(10, 18)))

	log.Printf("余额(ETH) [%v]：%v\n", a.Address().String(), ethValue)
	return balance, nil
}

// ScrollWithdrawETH 将scroll链上的eth提现到eth的测试网
func (a *Account) ScrollWithdrawETH() error {
	txOpts, err := a.makeScrollChainTxOpts()
	if err != nil {
		return err
	}

	// 设置gas limit
	txOpts.GasLimit = uint64(300000)

	// 设置value
	txOpts.Value = big.NewInt(20000000000000000)

	// 获取nonce
	nonce, err := a.GetNonce()
	if err != nil {
		return err
	}

	// 设置nonce
	txOpts.Nonce = nonce

	res, err := a.scrollCli.L2ABI.WithdrawETH0(txOpts, big.NewInt(18002100000000000), new(big.Int).SetUint64(uint64(160000)))
	if err != nil {
		return fmt.Errorf("call WithdrawETH error: %v\nf", err.Error())
	}

	a.PrintTx(res, "WithdrawETH")

	mined, err := a.waitMined(res)
	if err != nil {
		return err
	}

	a.PrintReceipt(mined)

	if mined.Status == types.ReceiptStatusFailed {
		log.Println("跨链交易[WithdrawETH]: ", "失败")
		return errors.New("交易失败")
	}

	log.Println("跨链交易[WithdrawETH]: ", "成功")
	return nil
}
