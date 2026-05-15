package main

import (
	"fmt"
)

type Bank struct {
	Name          string
	AccountNumber []string
	Balance       float64
}

func main() {

	bank := Bank{
		AccountNumber: []string{
			"01234567892",
			"22234521388",
			"41944783524",
		},
	}

	for _, acc := range bank.AccountNumber {
		fmt.Println(acc)
	}
}