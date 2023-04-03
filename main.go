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
	"encoding/hex"
	"flag"
	"fmt"
	redis "github.com/go-redis/redis/v8"
	"github.com/scroll-tech/go-ethereum/accounts/abi/bind"
	"github.com/scroll-tech/go-ethereum/accounts/keystore"
	"github.com/scroll-tech/go-ethereum/common"
	"github.com/scroll-tech/go-ethereum/core/types"
	"github.com/scroll-tech/go-ethereum/crypto"
	"github.com/scroll-tech/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
	"path/filepath"
	scroll_abi "scroll/abi"
	"strconv"
	"strings"
	"time"
)

var (
	goerliCli *ethclient.Client
	ScrollCli *ethclient.Client

	abi   = new(scroll_abi.ScrollABI)
	L2abi = new(scroll_abi.L2ABI)

	chainID       *big.Int
	ScrollchainID *big.Int

	KeyStoreDir = ""
	redisCli    = new(redis.Client)
)
var (
	M2S     bool
	Bridge1 bool
	S2M     bool

	// 模式
	Model string
)

const (
	GoerliAPI = "https://rpc.ankr.com/eth_goerli"
	ScrollAPI = "https://alpha-rpc.scroll.io/l2"
	scryptN   = keystore.LightScryptN
	scryptP   = keystore.LightScryptP

	Password = "jw"
)

const dirPerm = 0700

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

func setScrollClient(url string) {
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

	log.Println(address+" 余额(ETH)：", balance.String())
	return balance, nil
}

func sendToScroll(ctx context.Context, mainAccountPath, subAccountPath string) error {
	mainAccountAddress := filepath.Base(mainAccountPath)
	subAccountAddress := filepath.Base(subAccountPath)
	log.Printf("==========[%v]===========\n", subAccountAddress)

	a, err := redisCli.SIsMember(ctx, ETHToScrollOk, subAccountAddress).Result()
	if err != nil {
		log.Println("skip subAccountAddress address", subAccountAddress)
		return nil
	}

	if a {
		log.Println("skip to address", subAccountAddress)
		return nil
	}

	//发送跨链交易之前，如果子账户余额金额太小，比如0.1eth，会报错。所以先用主账户转到子账户。
	if M2S {
		num1 := big.NewInt(1300000000000000000)
		log.Printf("%v send %v eth to %v\n", mainAccountAddress, num1, subAccountAddress)
		sendETH(ctx, mainAccountPath, []string{subAccountAddress}, num1)

		log.Println("主账户转账到子账户交易完成", mainAccountAddress, "----", subAccountAddress)

		time.Sleep(7 * time.Second)
	}

	//转完之后，查看一下子账户的余额
	getBalance(ctx, subAccountAddress)

	var nonce uint64

	if Bridge1 {
		// 打开keyfile
		keyJson, err := os.ReadFile(subAccountPath)
		if err != nil {
			log.Printf("Could not read key file: %v\n", err)
			return err
		}

		txIDOpts, err := bind.NewTransactorWithChainID(strings.NewReader(string(keyJson)), Password, chainID)
		if err != nil {
			log.Println("new key store tx error", err.Error())
			return err
		}

		nonce, err = goerliCli.PendingNonceAt(ctx, txIDOpts.From)
		if err != nil {
			log.Println("get nonce error", err.Error(), subAccountAddress)
			return err
		}

		log.Println("nonce", nonce)

		gasPrice, err := goerliCli.SuggestGasPrice(ctx)
		if err != nil {
			log.Printf("[%v]get price error：%v\n", subAccountAddress, err)
			return err
		}

		log.Println("gas price", gasPrice.String())

		// 设置gas limit
		txIDOpts.GasLimit = uint64(300000)
		// 设置value
		txIDOpts.Value = big.NewInt(120000000000000000)

		// 设置nonce
		txIDOpts.Nonce = big.NewInt(int64(nonce))

		res, err := abi.DepositETH(txIDOpts, big.NewInt(110000000000000000), new(big.Int).SetUint64(uint64(40000)))
		if err != nil {
			log.Println("call Deposit Eth error: ", err.Error())
			return err
		}

		data, err := res.MarshalJSON()
		if err != nil {
			log.Println("marshal json error: ", err.Error())
			return err
		}
		log.Println("跨链交易的响应数据：", string(data))

		_, err = bind.WaitMined(ctx, goerliCli, res)
		if err != nil {
			log.Println("wait error", err.Error())
			return err
		}

		log.Println("wait 跨链交易完成")
		time.Sleep(10 * time.Second)
	}

	if S2M {
		// 子账户将余额的90%转回到主账户
		balance, _ := getBalance(ctx, subAccountAddress)

		var a1 *big.Int
		var ok bool

		// 循环五次
		for i := 0; i < 5; i++ {
			if i == 0 {
				//a1 = new(big.Int).Mul(balance, big.NewInt(97))
				if new(big.Int).Div(balance, big.NewInt(125000000000)).Int64() >= 1 {
					a1 = new(big.Int).Mul(balance, big.NewInt(98))
				} else {
					a1 = new(big.Int).Mul(balance, big.NewInt(95))
				}
			} else if i == 1 {
				a1 = new(big.Int).Mul(balance, big.NewInt(94))
			} else if i == 2 {
				a1 = new(big.Int).Mul(balance, big.NewInt(93))
			} else if i == 3 {
				a1 = new(big.Int).Mul(balance, big.NewInt(92))
			} else if i == 4 {
				a1 = new(big.Int).Mul(balance, big.NewInt(91))
			}

			b1 := new(big.Int).Div(a1, big.NewInt(100))
			err := sendETH(ctx, subAccountPath, []string{mainAccountAddress}, b1, nonce)
			if err != nil {
				time.Sleep(1 * time.Second)
				continue
			} else {
				ok = true
				break
			}
		}

		if !ok {
			log.Fatal("子账户转账到主账户交易失败")
		}

		err = redisCli.SAdd(ctx, ETHToScrollOk, subAccountAddress).Err()
		if err != nil {
			log.Fatal("save redis error", err.Error())
			return nil
		}

		log.Println("子账户转账到主账户交易完成")
		time.Sleep(8 * time.Second)
	}

	return nil
}

/*
SendToScroll 转账到scroll

	前提：主账户eth余额大于3
	1、遍历keystore目录下的1、2、3...文件夹下的账户
	2、主账户转账到该账户：转1.3枚eth
	3、该账户发起跨链交易：交易的value(0.12) > amount(0.11)，不然会失败
	4、该账户的余额的90%+转回到主账户
*/
func SendToScroll(ctx context.Context, mainAccountPath string) {

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
				log.Println("send to scroll", path)
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
func setETHABI(toAddress string) {
	var err error
	to := common.HexToAddress(toAddress)

	abi, err = scroll_abi.NewScrollABI(to, goerliCli)
	if err != nil {
		log.Println("new scroll abi error: ", err.Error())
		return
	}
}

func setScrollABI(toAddress string) {
	var err error
	to := common.HexToAddress(toAddress)

	L2abi, err = scroll_abi.NewL2ABI(to, ScrollCli)
	if err != nil {
		log.Println("new scroll abi error: ", err.Error())
		return
	}
}

func Before() {

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

	// 设置redis
	redisCli = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}

	// 设置scroll客户端
	if Model == "ETH" {
		setGoerliClient(GoerliAPI)
		setETHABI("0xe5E30E7c24e4dFcb281A682562E53154C15D3332")

	} else if Model == "Scroll" {
		setScrollClient(ScrollAPI)
		setScrollABI("0x6d79Aa2e4Fbf80CF8543Ad97e294861853fb0649")
	}
}

func sendETH(ctx context.Context, fromAddressPath string, sendToAddress []string, amount *big.Int, opt ...uint64) error {
	var nonce1 uint64
	if len(opt) != 0 {
		nonce1 = opt[0]
	}

	fromAddress := filepath.Base(fromAddressPath)

	// 获取余额
	_, err := getBalance(ctx, fromAddress)
	if err != nil {
		log.Println("获取余额失败")
		return err
	}

	// 打开keyfile
	keyJson, err := os.ReadFile(fromAddressPath)
	if err != nil {
		log.Printf("Could not read key file: %v\n", err.Error())
		return err
	}

	// 获取公钥、私钥
	key, err := keystore.DecryptKey(keyJson, Password)
	if err != nil {
		log.Printf("Could not decrypt key file: %v\n", err.Error())
		return err
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
		nonce, err := goerliCli.PendingNonceAt(ctx, key.Address)
		if err != nil {
			log.Println("get nonce error", err.Error())
			return err
		}

		log.Println("get nonce--- ", nonce)
		if nonce == nonce1 {
			nonce += 1
		}

		// 获取gas价格
		gasPrice, err := goerliCli.SuggestGasPrice(ctx)
		if err != nil {
			log.Println("get gas price error")
			return err
		}

		log.Println("gas price", gasPrice.String())

		toAddress := common.HexToAddress(to)

		// 构造交易
		baseTx := &types.LegacyTx{
			Nonce:    nonce,
			GasPrice: new(big.Int).Add(gasPrice, big.NewInt(88120347580)),
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
			return err
		}
		signTx, err := tx.WithSignature(signer, signature)
		if err != nil {
			log.Println("签名失败", err.Error())
			return err
		}
		log.Println("start send sign tx")

		// 发送交易
		err = goerliCli.SendTransaction(ctx, signTx)
		if err != nil {
			log.Println("send signTx error", err.Error())
			return err
		}

		data, err := signTx.MarshalJSON()
		if err != nil {
			log.Println("marshal json error: ", err.Error())
			return err
		}
		log.Println("转账交易的数据(Tx data)", string(data))

		_, err = bind.WaitMined(ctx, goerliCli, signTx)
		if err != nil {
			log.Println("wait error", err.Error())
			return err
		}

		//err = redisCli.SAdd(ctx, scrollAccountSendOk, to).Err()
		//if err != nil {
		//	log.Println("save redis error", err.Error())
		//}

		log.Println("wait 转账 ok", fromAddress, to)

		time.Sleep(2 * time.Second)
	}
	return nil
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

	log.Println("wait ok")

	time.Sleep(2 * time.Second)

}

const (
	scrollAccountSendOk  = "scroll:account:send_ok"
	scrollAccountSendOk1 = "scroll:account1:send_ok"
	ETHToScrollOk        = "scroll:eth_to_scroll:send_ok:1"
	ScrollToETHOk        = "scroll:scroll_to_eth:send_ok:1"
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

/*
UniSwap uni swap3 交互
	1、水龙头领取usdc稳定币
		0x6ad3cfc30298fce6b448643e917755dc1e32b7c1183a74ba10adfd7a627d7eda
	2、批准usdc稳定币：0xa0d71b9877f44c744546d649147e3f1e70a93760
		https://blockscout.scroll.io/tx/0x302c5f7341099f68912f3fdbfa6b94b40acc68cf45304ee90c3e72a44f1573ea
	3、usdc兑换一些weth
		0x15231ddd679b2a613552c39e8b632bdfaf73fd6c43993786f68819488641694d
	4、批准weth
		0x77430f8d58b53285da616bd70c00c9a0e2d0b08626c0371d9d1518a58c65e480
	5、批准usdc
		0x38aee49e61539b6cfc9861f6cbe0ec556207d32665355000d2bb61dd3f6630e7
	6、添加流动性
		0x848bc3634e957e0c5fa8f79a1933045bd8cb76e5a67484920516e4b79d6d915d
	7、移除流动性
		0xe470e88120aa8e3b33fe845cc5a30bdad56dfd7f1b446d9ec4b176d68358074c
*/

func UniSwap() {

}

func main() {

	model := flag.String("model", "ETH", "model: eth or scroll")

	m := flag.Bool("m2s", true, "main account send to sub account")
	s := flag.Bool("s2m", true, "sub account send to main account")
	b1 := flag.Bool("b1", true, "send eth to scroll")

	flag.Parse()

	M2S = *m
	S2M = *s
	Bridge1 = *b1
	Model = *model

	Before()

	// 获取账户列表
	//accountList := getAccounts()
	//log.Println("账户个数", len(accountList))

	// 批量转账
	//sendETH(context.Background(), "0x719C4b1559Ec3f8BBceFE099c9F6Cd6d58d15988", sendToAddressList)

	// 导入账户
	//err := getAccount(keyfile, password)
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

	if Model == "ETH" {
		SendToScroll(context.Background(), filepath.Join(KeyStoreDir, "0x1b2dE9662dF9983D7E87B9C064c0F6568516eC6B"))
		return
	}

}
