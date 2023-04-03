package main

import (
	"context"
	"flag"
	"fmt"
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
	ScrollAPI   = "https://alpha-rpc.scroll.io/l2"
	Password    = "jw"
	redisCli    = new(redis.Client)
)

const (
	ScrollToETHOk = "scroll:scroll_to_eth:send_ok:2"
)

func init() {
	// 设置redis
	redisCli = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		log.Fatal(err)
	}
}

func scrollSendToETH(ctx context.Context, mainAccount, subAccount *pkg.Account) (*pkg.Account, error) {
	a, err := redisCli.SIsMember(ctx, ScrollToETHOk, subAccount.Address().String()).Result()
	if err != nil {
		return nil, err
	}

	if a {
		log.Println("skip to address", subAccount.Address().String())
		return subAccount, nil
	}

	// 查询主账户的余额，将余额的98%转到子账户
	log.Printf("主 [%v] 转账余额的98%%到子账户 [%v] \n", mainAccount.Address().String(), subAccount.Address().String())
	err = mainAccount.SendScroll(subAccount.Address().String())
	if err != nil {
		return nil, err
	}

	log.Printf("主 [%v] 转账余额的98%%到子账户 [%v] : 成功\n", mainAccount.Address().String(), subAccount.Address().String())

	time.Sleep(2 * time.Second)

	//转完之后，查看一下子账户的余额
	_, err = subAccount.GetScrollBalance()
	if err != nil {
		return nil, err
	}

	// 跨链交易，scroll转账到eth
	err = subAccount.ScrollWithdrawETH()
	if err != nil {
		return nil, err
	}

	// 添加到redis
	redisCli.SAdd(ctx, ScrollToETHOk, subAccount.Address().String())

	time.Sleep(5 * time.Second)
	return subAccount, nil
}

/*
SendToETH scroll转账到以太坊

	前提：主账户eth余额大于3
	1、遍历keystore目录下的1、2、3...文件夹下的账户
	2、主账户转账到该账户：转1.3枚eth
	3、该账户发起跨链交易：交易的value(0.11) > amount(0.10)，不然会失败
	4、该账户的余额的98%转回到主账户
*/
func SendToETH(ctx context.Context, mainAccountPath string) error {

	cli, err := pkg.NewScrollClient(ctx, ScrollAPI)
	if err != nil {
		return err
	}

	// 主账户
	mainAccount, err := pkg.NewAccount(mainAccountPath, Password, cli)
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
				subAccount, err := pkg.NewAccount(path, Password, cli)
				if err != nil {
					return err
				}

				log.Println("new sub account ok: ", subAccount.Address().String())

				account, err := scrollSendToETH(ctx, mainAccount, subAccount)
				if err != nil {
					return err
				}

				// 交易完成后，将子账户设置为主账户
				mainAccount = account
			}
			return nil
		})
		if err != nil {
			fmt.Println("err", err.Error())
			continue
		}
	}

	return nil
}

func main() {
	keyStoreDir := flag.String("keystore", "/Users/jiangziya/code/github/scroll/cmd/account/keystore", "密钥文件夹")
	password := flag.String("password", "jw", "密码")
	flag.Parse()

	KeyStoreDir = *keyStoreDir
	Password = *password

	log.Println("==================send to eth start==================")
	log.Println("keystore文件夹", KeyStoreDir)

	err := SendToETH(context.Background(), filepath.Join(KeyStoreDir, "0x69b0b7f5079201D4d34C66A8AD7De56607F7dc40"))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("==================send to eth end==================")
}
