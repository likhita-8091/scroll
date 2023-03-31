/*
----------------------------------------------
 File Name: main.go
 Author: jw199@firecloud.ai
 @Company:  Firecloud
 Created Time: 2023/3/27 14:32
-------------------功能说明-------------------
 $
----------------------------------------------
*/

package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
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
	"path/filepath"
	scroll_abi "scroll/abi"
	"strconv"
	"strings"
	"time"
)

var goerliCli *ethclient.Client
var ScrollCli *ethclient.Client
var abi = new(scroll_abi.ScrollABI)
var L2abi = new(scroll_abi.L2ABI)
var chainID *big.Int
var ScrollchainID *big.Int
var KeyStoreDir = ""
var redisCli = new(redis.Client)

const (
	GoerliAPI = "https://rpc.ankr.com/eth_goerli"
	scryptN   = keystore.LightScryptN
	scryptP   = keystore.LightScryptP

	Password = "jw"
)
const dirPerm = 0700

func storeKey(filename string, key *keystore.Key, auth string) error {
	keyJson, err := keystore.EncryptKey(key, auth, scryptN, scryptP)

	if err != nil {
		return err
	}
	return os.WriteFile(filename, keyJson, dirPerm)
}

func newKey(dir string) (*keystore.Key, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf("Could not create random uuid: %v", err))
	}
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}

	prvPath := filepath.Join(dir, key.Address.String()+"_prv")
	err = crypto.SaveECDSA(prvPath, privateKeyECDSA)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// 创建账户
func createAccount(dir string, password string) (keyfile string) {
	key, err := newKey(dir)
	if err != nil {
		panic(err)
	}

	filename := filepath.Join(dir, key.Address.String())
	err = storeKey(filename, key, password)
	if err != nil {
		log.Fatal("create account error：", err.Error())
	}

	fmt.Printf("Public address of the key:   %s\n", key.Address.Hex())
	fmt.Printf("Path of the secret key file: %s\n\n", filename)
	return filename
}

// 获取账户
func getAccount(keyfile string, password string) error {
	keyJSON, err := os.ReadFile(keyfile)
	if err != nil {
		log.Fatalf("Could not read key file: %v", err)
	}
	key, err := keystore.DecryptKey(keyJSON, password)
	if err != nil {
		log.Fatalf("Could not decrypt key file: %v", err)
	}

	buf := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("公钥", key.Address.String())
	fmt.Println("私钥", hex.EncodeToString(buf))
	return nil
}

// 设置客户端
func setGoerliClient(url string) {
	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("获取链ID失败", err.Error())
	}

	chainID = id

	goerliCli = client
}

func setScrollClient() {
	url := "https://alpha-rpc.scroll.io/l2"
	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}

	id, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("获取scroll链ID失败", err.Error())
	}

	ScrollchainID = id
	log.Println("scroll chain id", ScrollchainID)

	ScrollCli = client
}

// 获取余额
func getBalance(ctx context.Context, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := goerliCli.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, err
	}

	log.Println(address+" 余额(ETH)：", balance.String())
	return balance, nil
}

func getScrollBalance(ctx context.Context, address string) (*big.Int, error) {
	account := common.HexToAddress(address)
	balance, err := ScrollCli.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, err
	}

	a := new(big.Int).SetInt64(int64(math.Pow(10, 18)))
	fmt.Println(address+" 余额(ETH)：", balance.Div(balance, a))

	return nil, nil
}

func sendToScroll(ctx context.Context, mainAccountPath, subAccountPath string) error {
	mainAccountAddress := filepath.Base(mainAccountPath)
	subAccountAddress := filepath.Base(subAccountPath)
	log.Printf("==========[%v]===========\n", subAccountAddress)

	//a, err := redisCli.SIsMember(ctx, ETHToScrollOk, subAccountAddress).Result()
	//if err != nil {
	//	log.Println("skip subAccountAddress address", subAccountAddress)
	//	return nil
	//}
	//
	//if a {
	//	log.Println("skip to address", subAccountAddress)
	//	return nil
	//}

	//发送跨链交易之前，如果子账户余额金额太小，比如0.1eth，会报错。所以先用主账户转到子账户。
	num1 := big.NewInt(1300000000000000000)
	log.Printf("%v send %v eth to %v\n", mainAccountAddress, num1, subAccountAddress)
	sendETH(ctx, mainAccountPath, []string{subAccountAddress}, num1)

	log.Println("主账户转账到子账户交易完成", mainAccountAddress, "----", subAccountAddress)

	time.Sleep(10 * time.Second)

	//转完之后，查看一下子账户的余额
	getBalance(ctx, subAccountAddress)

	// 打开keyfile
	keyJson, err := os.ReadFile(subAccountPath)
	if err != nil {
		log.Printf("Could not read key file: %v\n", err)
		return nil
	}

	txIDOpts, err := bind.NewTransactorWithChainID(strings.NewReader(string(keyJson)), Password, chainID)
	if err != nil {
		log.Println("new key store tx error", err.Error())
		return nil
	}

	nonce, err := goerliCli.PendingNonceAt(ctx, txIDOpts.From)
	if err != nil {
		log.Println("get nonce error", err.Error(), subAccountAddress)
		return nil
	}

	log.Println("nonce", nonce)

	gasPrice, err := goerliCli.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("[%v]get price error：%v\n", subAccountAddress, err)
		return nil
	}

	log.Println("gas price", gasPrice.String())

	// 设置gas limit
	txIDOpts.GasLimit = uint64(300000)
	// 设置value
	txIDOpts.Value = big.NewInt(120000000000000000)
	//txIDOpts.GasPrice = big.NewInt(250059291519)
	//txIDOpts.GasTipCap = tip
	//txIDOpts.GasFeeCap = gasFeeCap

	// 设置nonce
	txIDOpts.Nonce = big.NewInt(int64(nonce))

	res, err := abi.DepositETH(txIDOpts, big.NewInt(110000000000000000), new(big.Int).SetUint64(uint64(40000)))
	if err != nil {
		log.Println("call Deposit Eth error: ", err.Error())
		time.Sleep(3 * time.Second)
		return nil
	}

	data, err := res.MarshalJSON()
	if err != nil {
		log.Println("marshal json error: ", err.Error())
		return nil
	}
	log.Println("跨链交易的响应数据：", string(data))

	_, err = bind.WaitMined(ctx, goerliCli, res)
	if err != nil {
		log.Println("wait error", err.Error())
		return nil
	}

	err = redisCli.SAdd(ctx, ETHToScrollOk, subAccountAddress).Err()
	if err != nil {
		log.Println("save redis error", err.Error())
		return nil
	}

	log.Println("wait 跨链交易完成")
	time.Sleep(10 * time.Second)

	// 子账户将余额的90%转回到主账户
	balance, _ := getBalance(ctx, subAccountAddress)
	a := new(big.Int).Mul(balance, big.NewInt(95))
	b := new(big.Int).Div(a, big.NewInt(100))
	sendETH(ctx, subAccountPath, []string{mainAccountAddress}, b)
	log.Println("子账户转账到主账户交易完成")
	time.Sleep(10 * time.Second)

	return nil
}

/*
SendToScroll 发送交易

	前提：主账户eth余额大于3
	1、遍历keystore目录下的1、2、3...文件夹下的账户
	2、主账户转账到该账户：转1.3枚eth
	3、该账户发起跨链交易：交易的value(0.12) > amount(0.11)，不然会失败
	4、该账户将主账户发来的余额转回到主账户，1.3-0.12=1.11
*/
func SendToScroll(ctx context.Context, mainAccountPath string) {

	for i := 5; i < 11; i++ {

		dir1 := filepath.Join(KeyStoreDir, strconv.Itoa(i))
		log.Println("遍历文件夹：", dir1)
		err := filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				log.Println(path)
				return nil
			}
			if info.IsDir() {
				log.Println("skip")
				return nil
			}

			if !strings.Contains(path, "_prv") {
				sendToScroll(ctx, mainAccountPath, path)
			}
			return nil
		})
		if err != nil {
			return
		}
	}

}

// 加载abi合约
func setABI(toAddress string) {
	var err error
	to := common.HexToAddress(toAddress)

	abi, err = scroll_abi.NewScrollABI(to, goerliCli)
	if err != nil {
		log.Println("new scroll abi error: ", err.Error())
		return
	}
}

func setL2(toAddress string) {
	var err error
	to := common.HexToAddress(toAddress)

	L2abi, err = scroll_abi.NewL2ABI(to, ScrollCli)
	if err != nil {
		log.Println("new scroll abi error: ", err.Error())
		return
	}
}

func init() {
	// 设置keystore目录
	dir, _ := os.Getwd()
	KeyStoreDir = filepath.Join(dir, "./keystore")

	//for i := 1; i < 11; i++ {
	//	// 创建文件夹
	//	err := os.MkdirAll(filepath.Join(KeyStoreDir, strconv.Itoa(i)), os.ModePerm)
	//	if err != nil {
	//		log.Fatal("创建文件夹失败", err.Error())
	//	}
	//}

	// 设置客户端
	setGoerliClient(GoerliAPI)
	setScrollClient()

	// 设置redis
	redisCli = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	// 设置scroll客户端
	setABI("0xe5E30E7c24e4dFcb281A682562E53154C15D3332")
	setL2("0x6d79Aa2e4Fbf80CF8543Ad97e294861853fb0649")

}

// scroll转账到以太坊
func ScrollSendToETH(ctx context.Context, path string) error {
	from := filepath.Base(path)
	log.Printf("==========[%v]===========\n", from)

	//a, err := redisCli.SIsMember(ctx, ETHToScrollOk, from).Result()
	//if err != nil {
	//	log.Println("skip from address", from)
	//	return nil
	//}
	//
	//if a {
	//	log.Println("skip to address", from)
	//	return nil
	//}

	log.Println("send eth 0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B to", from)
	// 发送之前，如果账户余额金额太小，比如0.1,eth，会报错。所以先用其余账户转到该账户。
	//sendETH(ctx, "C:\\Users\\jw199\\code\\jw\\scroll\\keystore\\0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", "0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B", []string{from}, big.NewInt(1200000000000000000))

	//time.Sleep(10 * time.Second)

	getScrollBalance(ctx, from)

	// 打开keyfile
	keyJson, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Could not read key file: %v\n", err)
		return nil
	}

	txIDOpts, err := bind.NewTransactorWithChainID(strings.NewReader(string(keyJson)), Password, ScrollchainID)
	if err != nil {
		log.Println("new key store tx error", err.Error())
		return nil
	}

	nonce, err := ScrollCli.PendingNonceAt(ctx, txIDOpts.From)
	if err != nil {
		log.Println("get nonce error", err.Error(), from)
		return nil
	}

	log.Println("nonce", nonce)

	gasPrice, err := ScrollCli.SuggestGasPrice(ctx)
	if err != nil {
		log.Printf("[%v]get price error：%v\n", from, err)
		return nil
	}

	log.Println("gas price", gasPrice.String())

	//tip, err := goerliCli.SuggestGasTipCap(ctx)
	//if err != nil {
	//	log.Println("get gas tip cap error", err.Error())
	//	return nil
	//}

	//log.Println("get gas tip cap", tip)
	//
	//head, err := goerliCli.HeaderByNumber(ctx, nil)
	//if err != nil {
	//	log.Println("get header head error", err.Error())
	//	return err
	//}

	//gasFeeCap := new(big.Int).Add(
	//	tip,
	//	new(big.Int).Mul(head.BaseFee, big.NewInt(2)),
	//)
	//
	//log.Println("gas fee cap", gasFeeCap)

	// 设置gas
	txIDOpts.GasLimit = uint64(307936)
	// 设置value
	txIDOpts.Value = big.NewInt(40000000000000000)
	//txIDOpts.GasPrice = big.NewInt(250059291519)
	//txIDOpts.GasTipCap = tip
	//txIDOpts.GasFeeCap = gasFeeCap

	// 设置nonce
	txIDOpts.Nonce = big.NewInt(int64(nonce))

	res, err := L2abi.WithdrawETH0(txIDOpts, big.NewInt(1000000000000000), new(big.Int).SetUint64(uint64(160000)))
	if err != nil {
		log.Println("call Deposit Eth error: ", err.Error())
		time.Sleep(3 * time.Second)
		return nil
	}

	log.Println("res value", res.Value().String())

	data, err := res.MarshalJSON()
	if err != nil {
		log.Println("marshal json error: ", err.Error())
		return nil
	}
	log.Println("res data", string(data))
	_, err = bind.WaitMined(ctx, ScrollCli, res)
	if err != nil {
		log.Println("wait error", err.Error())
		return nil
	}

	//err = redisCli.SAdd(ctx, ETHToScrollOk, from).Err()
	//if err != nil {
	//	log.Println("save redis error", err.Error())
	//	return nil
	//}

	log.Println("wait ok")
	time.Sleep(10 * time.Second)

	// 再发送回去
	//sendETH(ctx, path, from, []string{"0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B"}, big.NewInt(1239990000000000000))

	time.Sleep(10 * time.Second)
	return nil
}

func sendETH(ctx context.Context, fromAddressPath string, sendToAddress []string, amount *big.Int) {

	fromAddress := filepath.Base(fromAddressPath)

	// 获取余额
	_, err := getBalance(ctx, fromAddress)
	if err != nil {
		log.Println("获取余额失败")
		return
	}

	// 打开keyfile
	keyJson, err := os.ReadFile(fromAddressPath)
	if err != nil {
		log.Printf("Could not read key file: %v\n", err.Error())
		return
	}

	// 获取公钥、私钥
	key, err := keystore.DecryptKey(keyJson, Password)
	if err != nil {
		log.Printf("Could not decrypt key file: %v\n", err.Error())
		return
	}

	for _, to := range sendToAddress {
		//a, err := redisCli.SIsMember(ctx, scrollAccountSendOk, to).Result()
		//if err != nil {
		//	log.Println("skip to address", to)
		//	continue
		//}
		//
		//if a {
		//	log.Println("skip to address", to)
		//	continue
		//}

		// 获取nonce
		time.Sleep(1 * time.Second)
		nonce, err := goerliCli.PendingNonceAt(ctx, key.Address)
		if err != nil {
			log.Println("get nonce error", err.Error())
			return
		}

		// 获取gas价格
		gasPrice, err := goerliCli.SuggestGasPrice(ctx)
		if err != nil {
			log.Println("get gas price error")
			return
		}

		log.Println("gas price", gasPrice.String())

		toAddress := common.HexToAddress(to)

		// 构造交易
		baseTx := &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: big.NewInt(300584513913),
			Gas:      uint64(40000),
			To:       &toAddress,
			Value:    amount, // 0.1 eth
		}
		tx := types.NewTx(baseTx)

		// 交易签名
		signer := types.LatestSignerForChainID(chainID)
		signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key.PrivateKey)
		if err != nil {
			log.Println("生成signature失败", err.Error())
			return
		}
		signTx, err := tx.WithSignature(signer, signature)
		if err != nil {
			log.Println("签名失败", err.Error())
			return
		}
		log.Println("start send sign tx")

		// 发送交易
		err = goerliCli.SendTransaction(ctx, signTx)
		if err != nil {
			log.Fatal("send signTx error", err.Error())
			return
		}

		data, err := signTx.MarshalJSON()
		if err != nil {
			log.Println("marshal json error: ", err.Error())
			return
		}
		log.Println("Tx data", string(data))

		_, err = bind.WaitMined(ctx, goerliCli, signTx)
		if err != nil {
			log.Println("wait error", err.Error())
			return
		}

		//err = redisCli.SAdd(ctx, scrollAccountSendOk, to).Err()
		//if err != nil {
		//	log.Println("save redis error", err.Error())
		//}

		log.Println("wait 转账 ok", fromAddress, to)

		time.Sleep(3 * time.Second)
	}
}

func sendETH1(ctx context.Context, keyfile string, fromAddress string, to string) {

	// 获取余额
	_, err := getBalance(ctx, fromAddress)
	if err != nil {
		log.Println("获取余额失败")
		return
	}

	// 打开keyfile
	keyJson, err := os.ReadFile(keyfile)
	if err != nil {
		log.Printf("Could not read key file: %v\n", err)
		return
	}

	// 获取公钥、私钥
	key, err := keystore.DecryptKey(keyJson, Password)
	if err != nil {
		log.Printf("Could not decrypt key file: %v\n", err)
		return
	}

	a, err := redisCli.SIsMember(ctx, scrollAccountSendOk1, fromAddress).Result()
	if err != nil {
		log.Println("skip to address", to)
		return
	}

	if a {
		log.Println("skip to address", to)
		return
	}

	// 获取nonce
	nonce, err := goerliCli.PendingNonceAt(ctx, key.Address)
	if err != nil {
		log.Println("get nonce error", err.Error())
		return
	}
	log.Println("nonce", nonce)

	// 获取gas价格
	gasPrice, err := goerliCli.SuggestGasPrice(ctx)
	if err != nil {
		log.Println("get gas price error")
		return
	}
	log.Println("gas price", gasPrice.String())

	time.Sleep(1 * time.Second)

	toAddress := common.HexToAddress(to)

	// 构造交易
	baseTx := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: big.NewInt(300584513913),
		Gas:      uint64(40000),
		To:       &toAddress,
		Value:    big.NewInt(int64(91999999999999999)), // 0.1 eth
	}
	tx := types.NewTx(baseTx)

	// 交易签名
	signer := types.LatestSignerForChainID(chainID)
	signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key.PrivateKey)
	if err != nil {
		log.Println("生成signature失败", err.Error())
		return
	}
	signTx, err := tx.WithSignature(signer, signature)
	if err != nil {
		log.Println("签名失败", err.Error())
		return
	}
	log.Println("start send sign tx")

	// 发送交易
	err = goerliCli.SendTransaction(ctx, signTx)
	if err != nil {
		log.Println("send signTx error", err.Error())
		return
	}

	log.Println("转账成功")
	log.Println("tx value", signTx.Value().String())

	data, err := signTx.MarshalJSON()
	if err != nil {
		log.Println("marshal json error: ", err.Error())
		return
	}
	log.Println("Tx data", string(data))

	_, err = bind.WaitMined(ctx, goerliCli, signTx)
	if err != nil {
		log.Println("wait error", err.Error())
		return
	}

	err = redisCli.SAdd(ctx, scrollAccountSendOk1, fromAddress).Err()
	if err != nil {
		log.Println("save redis error", err.Error())
	}

	log.Println("wait ok")

	time.Sleep(2 * time.Second)

}

const (
	scrollAccountSendOk  = "scroll:account:send_ok"
	scrollAccountSendOk1 = "scroll:account1:send_ok"
	ETHToScrollOk        = "scroll:eth_to_scroll:send_ok"
)

func getAccounts() []string {
	var accountList []string
	for i := 1; i < 11; i++ {
		dir1 := filepath.Join(KeyStoreDir, strconv.Itoa(i))
		log.Println("遍历文件夹：", dir1)
		err := filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				log.Println(path)
				return nil
			}
			if info.IsDir() {
				log.Println("skip")
				return nil
			}
			if !strings.Contains(path, "_prv") {
				accountList = append(accountList, filepath.Base(path))
			}
			return nil
		})
		if err != nil {
			return nil
		}
	}
	return accountList
}

func main() {

	// 创建账户
	//for i := 1; i < 11; i++ {
	//	for j := 1; j < 11; j++ {
	//		createAccount(filepath.Join(KeyStoreDir, strconv.Itoa(i)), Password)
	//	}
	//}

	// 获取账户列表
	//accountList := getAccounts()
	//log.Println("账户个数", len(accountList))

	// 批量转账
	//sendETH(context.Background(), "0x719C4b1559Ec3f8BBceFE099c9F6Cd6d58d15988", sendToAddressList)

	//for i := 1; i < 11; i++ {
	//	dir1 := filepath.Join(KeyStoreDir, strconv.Itoa(i))
	//	log.Println("遍历文件夹：", dir1)
	//	err := filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			log.Println(err)
	//			log.Println(path)
	//			return nil
	//		}
	//		if info.IsDir() {
	//			log.Println("skip")
	//			return nil
	//		}
	//		if !strings.Contains(path, "_prv") {
	//			sendETH1(context.Background(), path, filepath.Base(path), "0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B")
	//		}
	//		return nil
	//	})
	//	if err != nil {
	//		return
	//	}
	//}

	// 导入账户
	//err := getAccount(keyfile, Password)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	// 查询余额
	//abs, _ := filepath.Abs(keyfile)
	//balance, err := getBalance(context.Background(), "0x69b0b7f5079201D4d34C66A8AD7De56607F7dc40")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(balance.String())

	// 批量转账到scroll
	//SendToScroll(context.Background(), accountList)
	SendToScroll(context.Background(), filepath.Join(KeyStoreDir, "0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B"))

	//ScrollSendToETH(context.Background(), "C:\\Users\\jw199\\code\\jw\\scroll\\keystore\\0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B")
}
