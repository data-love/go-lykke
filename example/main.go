package main

import (
	"fmt"

	"github.com/data-love/go-lykke-trading/lykke"
)

func main() {
	client := lykke.New("some id")

	resp, err := client.Ping()
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

	order, err := client.GetOrder("someID")
	if err != nil {
		panic(err)
	}
	fmt.Println(order)

	marketOrder := lykke.MarketOrder{
		AssetPairID: "BTCEUR",
		Asset:       "BTC",
		OrderAction: "Sell",
		Volume:      1,
	}

	err = client.AddMarketOrder(marketOrder)
	if err != nil {
		panic(err)
	}

	limitOrder := lykke.LimitOrder{
		AssetPairID: "BTCEUR",
		OrderAction: "Buy",
		Volume:      1,
		Price:       1.394543,
	}

	err = client.AddLimitOrder(limitOrder)
	if err != nil {
		panic(err)
	}

	wallets, err := client.GetWallets()
	if err != nil {
		panic(err)
	}
	fmt.Println(wallets)
}
