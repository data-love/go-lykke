package main

import (
	"fmt"

	"github.com/data-love/go-lykke-trading/lykke"
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

	pairs, err := client.GetAssetPairs()
	if err != nil {
		panic(err)
	}
	fmt.Println(pairs)

	pair, err := client.GetAssetPair("AUDUSD")
	if err != nil {
		panic(err)
	}
	fmt.Println(pair)

	pairBooks, err := client.GetOrderBooksByAssetPair("AUDUSD")
	if err != nil {
		panic(err)
	}
	fmt.Println(pairBooks)

	orders, err := client.GetOrders()
	if err != nil {
		panic(err)
	}
	fmt.Println(orders)

	order, err := client.GetOrders("someID")
	if err != nil {
		panic(err)
	}
	fmt.Println(order)

}
