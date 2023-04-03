package main

import (
	"context"
	"flag"
	"github.com/go-redis/redis/v8"
	"log"
	"path/filepath"
)

var (
	KeyStoreDir = ""
	ScrollAPI   = "https://alpha-rpc.scroll.io/l2"
	Password    = "jw"
	redisCli    = new(redis.Client)
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

		log.Println("==================send to eth start==================")

		err := SendToETH(context.Background(), filepath.Join(KeyStoreDir, "0x69b0b7f5079201D4d34C66A8AD7De56607F7dc40"))
		if err != nil {
			log.Println(err)
			return
		}

		log.Println("==================send to eth end==================")
	}
}
