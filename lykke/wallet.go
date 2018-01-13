package lykke

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Wallets []Wallet

type Wallet struct {
	AssetID  string  `json:"AssetId"`
	Balance  float32 `json:"Balance"`
	Reserved float32 `json:"Reserved"`
}

func (c *Client) GetWallets() (*Wallets, error) {
	url := fmt.Sprintf(baseURL + "/Wallets")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Wallets
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
