package main

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"path/filepath"
	"scroll/pkg"
	"strconv"
	"strings"
	"time"
)

var (
	KeyStoreDir = ""
	GoerliAPI   = "https://rpc.ankr.com/eth_goerli"
	Password    = "jw"
	redisCli    = new(redis.Client)
)

const (
	ETHSendToScrollOK = "scroll:eth_to_scroll:send_ok:2"
)

func before() {
	// 设置redis
	redisCli = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}

func SendToScroll(ctx context.Context, mainAccountPath string) error {

	cli, err := pkg.NewEthClient(ctx, GoerliAPI)
	if err != nil {
		return err
	}

	// 主账户
	mainAccount, err := pkg.NewAccount(mainAccountPath, Password, nil, cli)
	if err != nil {
		return err
	}

	log.Println("new main account ok: ", mainAccount.Address().String())

	for i := 1; i < 11; i++ {

		dir1 := filepath.Join(KeyStoreDir, strconv.Itoa(i))
		log.Println("==========遍历文件夹==========: ", dir1)
		err := filepath.Walk(dir1, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)
				log.Println(path)
				return nil
			}
			if info.IsDir() {
				log.Println("skip dir")
				return nil
			}

			if !strings.Contains(path, "_prv") {
				log.Println("scroll send to eth", path)

				// 子账户
				subAccount, err := pkg.NewAccount(path, Password, nil, cli)
				if err != nil {
					return err
				}

				log.Println("new sub account ok: ", subAccount.Address().String())

				account, err := ETHSendToScroll(ctx, mainAccount, subAccount)
				if err != nil {
					return err
				}

				// 交易完成后，将子账户设置为主账户
				mainAccount = account
			}
			return nil
		})
		if err != nil {
			log.Println("err", err.Error())
			continue
		}
	}

	return nil
}

func ETHSendToScroll(ctx context.Context, mainAccount, subAccount *pkg.Account) (*pkg.Account, error) {
	a, err := redisCli.SIsMember(ctx, ETHSendToScrollOK, subAccount.Address().String()).Result()
	if err != nil {
		return nil, err
	}

	if a {
		log.Println("skip to address", subAccount.Address().String())
		return subAccount, nil
	}

	// 查询主账户的余额，将余额的98%转到子账户
	log.Printf("主 [%v] 转账余额的98%%到子账户 [%v] \n", mainAccount.Address().String(), subAccount.Address().String())
	err = mainAccount.SendEth(subAccount.Address().String())
	if err != nil {
		return nil, err
	}

	log.Printf("主 [%v] 转账余额的98%%到子账户 [%v] : 成功\n", mainAccount.Address().String(), subAccount.Address().String())

	time.Sleep(2 * time.Second)

	//转完之后，查看一下子账户的余额
	_, err = subAccount.GetEthBalance()
	if err != nil {
		return nil, err
	}

	// 跨链交易，eth转账到scroll
	err = subAccount.EthDepositScroll()
	if err != nil {
		return nil, err
	}

	// 添加到redis
	redisCli.SAdd(ctx, ETHSendToScrollOK, subAccount.Address().String())

	time.Sleep(5 * time.Second)
	return subAccount, nil
}

func main() {
	keyStoreDir := flag.String("keystore", "/Users/jiangziya/code/github/scroll/cmd/account/keystore", "密钥文件夹")
	password := flag.String("password", "jw", "密码")
	cmd := flag.String("cmd", "send_to_scroll", "cmd：send_to_scroll")
	flag.Parse()

	KeyStoreDir = *keyStoreDir
	Password = *password

	log.Println("keystore文件夹", KeyStoreDir)
	log.Println("cmd命令", *cmd)

	before()

	switch *cmd {
	case "send_to_scroll":

		log.Println("==================send to eth to scroll start==================")

		err := SendToScroll(context.Background(), filepath.Join(KeyStoreDir, "8", "0x709e7814635bf0eeA85820ED2470B90b93f51A1B"))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("==================send to eth to scroll end==================")
	}
}
