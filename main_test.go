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
	"math/big"
	"testing"
)

func Test1(t *testing.T) {
	balance := big.NewInt(1183650979312180731)

	a := new(big.Int).Mul(balance, big.NewInt(9.1))
	b := new(big.Int).Div(a, big.NewInt(10))
	fmt.Println(a)
	fmt.Println(b)
}
