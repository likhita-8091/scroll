package main

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
	"math/big"
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
	ScrollToETHOk       = "scroll:scroll_to_eth:send_ok:2"
	ScrollClaimUsdc     = "scroll:claim_usdc_ok:1"
	ScrollSwapUsdc2WEth = "scroll:swap_ok:1"
	ScrollAddLP         = "scroll:add_lp:1"
	ScrollApproveUsdc   = "scroll:approve_usdc_ok:1"
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
	2、主账户将余额的99%转账到该子账户
	3、该账户发起跨链交易：交易的value(0.11) > amount(0.10)，不然会失败
	4、该账户作为主账户，跳到1步骤
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
			log.Println("err", err.Error())
			continue
		}
	}

	return nil
}

// ClaimUsdc 水龙头领取usdc，并且approve
func ClaimUsdc(ctx context.Context) error {
	cli, err := pkg.NewScrollClient(ctx, ScrollAPI)
	if err != nil {
		return err
	}

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
				log.Println("scroll claim usdc", path)

				subAccount, err := pkg.NewAccount(path, Password, cli)
				if err != nil {
					return err
				}

				exist, err := redisCli.SIsMember(ctx, ScrollClaimUsdc, subAccount.Address().String()).Result()
				if err != nil {
					log.Println("redis 调用错误：", err.Error())
					return err
				}

				// 如果不存在，才去claim usdc
				if !exist {
					err = subAccount.ClaimUSDC()
					if err != nil {
						return err
					}

					// 添加到redis
					redisCli.SAdd(ctx, ScrollClaimUsdc, subAccount.Address().String())

					log.Println("claim usdc ok: ", subAccount.Address().String())
				}

				time.Sleep(1 * time.Second)

				exist, err = redisCli.SIsMember(ctx, ScrollApproveUsdc, subAccount.Address().String()).Result()
				if err != nil {
					log.Println("redis 调用错误：", err.Error())
					return err
				}

				if exist {
					log.Println("skip approve：", subAccount.Address().String())
					return nil
				}

				a := big.NewInt(1000000000000000000)
				b := new(big.Int).Mul(a, big.NewInt(10000))
				err = subAccount.ApproveUSDC("0xd9880690bd717189cc3fbe7b9020f27fae7ac76f", b)
				if err != nil {
					return err
				}

				// 添加到redis
				redisCli.SAdd(ctx, ScrollApproveUsdc, subAccount.Address().String())
				log.Println("approve usdc ok: ", subAccount.Address().String())
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

// 交换代币：usdc--weth
func Swap(ctx context.Context) error {
	cli, err := pkg.NewScrollClient(ctx, ScrollAPI)
	if err != nil {
		return err
	}

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
				log.Println("scroll uni swap usdc to weth", path)

				subAccount, err := pkg.NewAccount(path, Password, cli)
				if err != nil {
					return err
				}

				exist, err := redisCli.SIsMember(ctx, ScrollSwapUsdc2WEth, subAccount.Address().String()).Result()
				if err != nil {
					log.Println("redis 调用错误：", err.Error())
					return err
				}

				if exist {
					log.Println("skip swap", subAccount.Address().String())
					return nil
				}

				a := big.NewInt(1000000000000000000)
				b := new(big.Int).Mul(a, big.NewInt(2200))
				err = subAccount.MultiCall(b)
				if err != nil {
					return err
				}

				// 添加到redis
				redisCli.SAdd(ctx, ScrollSwapUsdc2WEth, subAccount.Address().String())
				log.Println("swap usdc to weth ok: ", subAccount.Address().String())
				time.Sleep(2 * time.Second)
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

// AddLP 添加流动性
func AddLP(ctx context.Context) error {
	cli, err := pkg.NewScrollClient(ctx, ScrollAPI)
	if err != nil {
		return err
	}

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
				log.Println("scroll uni add lp", path)

				subAccount, err := pkg.NewAccount(path, Password, cli)
				if err != nil {
					return err
				}

				exist, err := redisCli.SIsMember(ctx, ScrollAddLP, subAccount.Address().String()).Result()
				if err != nil {
					log.Println("redis 调用错误：", err.Error())
					return err
				}

				if exist {
					log.Println("skip add uni lp", subAccount.Address().String())
					return nil
				}

				// 批准：授权人、数量
				toAddress := "0xbd1a5920303f45d628630e88afbaf012ba078f37"
				a := big.NewInt(1000000000000000000)
				b := new(big.Int).Mul(a, big.NewInt(10000))
				// 批准usdc
				err = subAccount.ApproveUSDC(toAddress, b)
				if err != nil {
					return err
				}

				time.Sleep(2 * time.Second)

				// 批准weth
				err = subAccount.ApproveETH(toAddress, b)
				if err != nil {
					return err
				}
				time.Sleep(2 * time.Second)

				err = subAccount.AddLP()
				if err != nil {
					log.Println("add uni lp error", err)
				}

				// 添加到redis
				redisCli.SAdd(ctx, ScrollAddLP, subAccount.Address().String())
				log.Println("add uni lp ok: ", subAccount.Address().String())

				time.Sleep(2 * time.Second)
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

func main() {
	keyStoreDir := flag.String("keystore", "/Users/jiangziya/code/github/scroll/cmd/account/keystore", "密钥文件夹")
	password := flag.String("password", "jw", "密码")
	cmd := flag.String("cmd", "send_to_eth", "cmd：send_to_eth、claim_usdc、swap、add_lp")
	flag.Parse()

	KeyStoreDir = *keyStoreDir
	Password = *password

	log.Println("keystore文件夹", KeyStoreDir)
	log.Println("cmd命令", *cmd)

	before()

	switch *cmd {
	case "send_to_eth":

		log.Println("==================send to weth start==================")

		err := SendToETH(context.Background(), filepath.Join(KeyStoreDir, "0x69b0b7f5079201D4d34C66A8AD7De56607F7dc40"))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("==================send to weth end==================")

	case "claim_usdc":

		log.Println("==================claim usdc start==================")

		err := ClaimUsdc(context.Background())
		if err != nil {
			log.Println(err)
		}

		log.Println("=====================claim usdc end=================")

	case "swap":
		log.Println("==================swap usdc<-->eth start==================")

		err := Swap(context.Background())
		if err != nil {
			log.Println(err)
		}

		log.Println("=====================swap usdc<-->eth end=================")
	case "add_lp":
		log.Println("==================添加流动性start==================")

		err := AddLP(context.Background())
		if err != nil {
			log.Println(err)
		}

		log.Println("=====================添加流动性end=================")
	}
}
