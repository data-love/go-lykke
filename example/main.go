package main

import (
	"fmt"

	"github.com/data-love/go-lykke/lykke"
)

func main() {
	client := lykke.NewApiClient("apikey")

	resp, err := client.GetIsAlive()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)

	books, err := client.GetOrderBooks()
	if err != nil {
		panic(err)
	}
	fmt.Println(books)
}
