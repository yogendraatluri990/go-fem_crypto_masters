package main

import (
	"fmt"

	cryptoApi "github.com/yogendraatluri990/go-fem_crypto_masters/api"
)

func main() {
	rate, err := cryptoApi.GetRateViaPost("BTC")

	fmt.Println(rate, err)
}
