package main

import (
	"crypto/ecdsa"
	"crypto/rand"
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

func main() {

	// 设置keystore目录
	dir, _ := os.Getwd()
	defaultDir := filepath.Join(dir, "./keystore")

	create := flag.Bool("create", false, "是否仅仅创建账户，创建完程序退出")
	_ = flag.Bool("list", false, "查看账户列表")

	keyStoreDir := flag.String("keystore", defaultDir, "密钥文件夹")
	password := flag.String("password", "jw", "密码")

	flag.Parse()

	// 创建账户
	if *create {
		for i := 1; i < 11; i++ {
			for j := 1; j < 11; j++ {
				createAccount(filepath.Join(*keyStoreDir, strconv.Itoa(i)), *password)
			}
		}
		log.Println("创建账户完成")
		return
	}
}
