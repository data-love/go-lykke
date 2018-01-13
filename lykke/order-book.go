package lykke

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderBooks []OrderBook

type OrderBook struct {
	AssetPair string   `json:"AssetPair"`
	IsBuy     bool     `json:"IsBuy"`
	Timestamp string   `json:"Timestamp"`
	Prices    *[]Price `json:"Prices"`
}

type Price struct {
	Volume float32 `json:"Volume"`
	Price  float32 `json:"Price"`
}

func (c *Client) GetOrderBooks() (*OrderBooks, error) {
	url := fmt.Sprintf(baseURL + "/OrderBooks")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data OrderBooks
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetOrderBooksByAssetPair(assetPairID string) (*OrderBooks, error) {
	url := fmt.Sprintf(baseURL+"/OrderBooks/%s", assetPairID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data OrderBooks
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
