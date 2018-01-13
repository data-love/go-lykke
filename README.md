# go-lykke-trading

This is a lightweight API Wrapper for the Lykke Trading API written in __GO__.

[Swagger API Description](https://hft-api.lykke.com/swagger/ui/#/)

# Action Examples

In general most of the calls are public and do not need the API key. POST Actions of Orders do have a API key required.

- initialise the the lykke trading client as follows.

```
package main

import (
	"fmt"

	"github.com/data-love/go-lykke-trading/lykke"
)

func main() {
	client := lykke.NewApiClient("apikey")
}
```

- Check API Status

```
resp, err := client.GetIsAlive()
if err != nil {
  panic(err)
}
fmt.Println(resp)
```

- GetAssetPairs

```
pairs, err := client.GetAssetPairs()
if err != nil {
  panic(err)
}
fmt.Println(pairs)

```

- GetAssetPair

```
pair, err := client.GetAssetPair("AUDUSD")
if err != nil {
  panic(err)
}
fmt.Println(pair)

```

- GetOrderBooksByAssetPair

```
pair, err := client.GetAssetPair("AUDUSD")
if err != nil {
  panic(err)
}
fmt.Println(pair)

```

- GetOrders

```
orders, err := client.GetOrders()
if err != nil {
  panic(err)
}
fmt.Println(orders)

```

- GetOrder

```
order, err := client.GetOrders("someID")
if err != nil {
  panic(err)
}
fmt.Println(order)

```

- AddLimitOrder

```
limitOrder := lykke.LimitOrder{
  AssetPairID: "BTCEUR",
  OrderAction: "buy",
  Volume:      1,
  Price:       1.394543,
}

err = client.AddLimitOrder(limitOrder)
if err != nil {
  panic(err)
}

```

- AddMarketOrder

```
marketOrder := lykke.MarketOrder{
  AssetPairID: "BTCEUR",
  Asset:       "BTC",
  OrderAction: "sell",
  Volume:      1,
}

err = client.AddMarketOrder(marketOrder)
if err != nil {
  panic(err)
}


```

- CancelOrder

```
err := client.CancelOrder("someID")
if err != nil {
  panic(err)
}


```

- GetWallets

```
wallets, err := client.GetWallets()
if err != nil {
  panic(err)
}
fmt.Println(wallets)
```
