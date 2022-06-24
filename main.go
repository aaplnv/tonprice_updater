package main

import (
	"fmt"
	"main/markets"
)

func main() {
	result, err := markets.GetAllWithCoinGecko()

	if err != nil {
		fmt.Println("Can't make request: ", err)
	}
	fmt.Println("Current ton price: ", result)

}
