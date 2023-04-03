package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/scroll-tech/go-ethereum/accounts/keystore"
	"github.com/scroll-tech/go-ethereum/crypto"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

const (
	scryptN = keystore.LightScryptN
	scryptP = keystore.LightScryptP
	dirPerm = 0700
)

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

	//prvPath := filepath.Join(dir, key.Address.String()+"_prv")
	//err = crypto.SaveECDSA(prvPath, privateKeyECDSA)
	//if err != nil {
	//	return nil, err
	//}

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

	log.Printf("Public address of the key:   %s\n", key.Address.Hex())
	log.Printf("Path of the secret key file: %s\n", filename)
	return filename
}

// 获取账户私钥
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
	log.Println("公钥", key.Address.String())
	log.Println("私钥", hex.EncodeToString(buf))
	return nil
}

func main() {

	// 设置keystore目录
	dir, _ := os.Getwd()
	defaultDir := filepath.Join(dir, "./keystore")

	creates := flag.Bool("creates", false, "批量创建账户，创建完程序退出")
	createOne := flag.Bool("create", false, "创建1个账户，创建完程序退出")
	get := flag.Bool("get", false, "获取一个账户的私钥信息，需指定账户公钥")
	account := flag.String("account", "", "获取一个账户的私钥信息，需指定账户公钥路径")
	_ = flag.Bool("list", false, "查看账户列表")

	keyStoreDir := flag.String("keystore", defaultDir, "密钥文件夹")
	password := flag.String("password", "jw", "密码")

	flag.Parse()

	err := os.MkdirAll(*keyStoreDir, dirPerm)
	if err != nil {
		log.Fatalln("创建keystore文件夹失败")
	}

	if *get {
		if *account == "" {
			log.Fatalln("没有指定账户路径，Example：-account /a/b/c/xxx")
		}

		err = getAccount(*account, *password)
		if err != nil {
			log.Fatalln("获取账户信息失败：", err.Error())
		}
		log.Println("获取账户信息")
	}

	if *createOne && *creates {
		log.Fatal("不能同时指定")
	}

	// 创建一个账户
	if *createOne {
		createAccount(*keyStoreDir, *password)
		log.Println("创建一个账户成功")
		return
	}

	// 批量创建账户
	if *creates {
		num := 0
		for i := 1; i < 11; i++ {
			// 先判断有没有已经存在文件夹
			d := filepath.Join(*keyStoreDir, strconv.Itoa(i))
			_, err = os.Stat(d)
			if !os.IsNotExist(err) {
				readDir, err := os.ReadDir(d)
				if err != nil {
					log.Fatal(err)
					return
				}

				if len(readDir) != 0 {
					log.Fatal("文件夹已经存在：", d)
				}
			}

			err := os.MkdirAll(d, dirPerm)
			if err != nil {
				log.Println("创建文件夹失败：", d)
			}

			for j := 1; j < 11; j++ {
				createAccount(d, *password)
				num += 1
			}
		}
		log.Printf("批量创建%v个账户完成\n", num)
		return
	}
}
