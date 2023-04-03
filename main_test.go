/*
----------------------------------------------
 File Name: main_test.go
 Author: jw199@firecloud.ai
 @Company:  Firecloud
 Created Time: 2023/3/31 17:56
-------------------功能说明-------------------
 $
----------------------------------------------
*/

package main

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func Test1(t *testing.T) {
	balance := big.NewInt(19999999999)
	//balance := big.NewInt(120000000000)

	fmt.Println(new(big.Int).Div(balance, big.NewInt(120000000000)).Int64())

	a := new(big.Int).Mul(balance, big.NewInt(9))
	b := new(big.Int).Div(a, big.NewInt(10))
	fmt.Println(a)
	fmt.Println(b)
}

func Test2(t *testing.T) {
	balance, _ := new(big.Int).SetString("120000000000", 10)
	b1 := new(big.Float)
	b1.SetString(balance.String())
	ethValue := new(big.Float).Quo(b1, big.NewFloat(math.Pow(10, 18)))

	fmt.Println(ethValue)
}

func Test3(t *testing.T) {
}
