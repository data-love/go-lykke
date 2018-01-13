package lykke

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Orders []Order

type Order struct {
	ID              string  `json:"Id"`
	ClientID        string  `json:"ClientId"`
	Status          string  `json:"Status"`
	AssetPairID     string  `json:"AssetPairId"`
	Volume          float32 `json:"Volume"`
	Price           float32 `json:"Price"`
	RemainingVolume float32 `json:"RemainingVolume"`
	LastMatchTime   string  `json:"LastMatchTime"`
	CreatedAt       string  `json:"CreatedAt"`
	Registered      string  `json:"Registered"`
}

type MarketOrder struct {
	AssetPairID string  `json:"AssetPairId"`
	Asset       string  `json:"Asset"`
	OrderAction string  `json:"OrderAction"`
	Volume      float32 `json:"Volume"`
}

type LimitOrder struct {
	AssetPairID string  `json:"AssetPairId"`
	OrderAction string  `json:"OrderAction"`
	Volume      float32 `json:"Volume"`
	Price       float32 `json:"Price"`
}

func (c *Client) GetOrders() (*Orders, error) {
	url := fmt.Sprintf(baseURL + "/Orders")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("api-key", c.APIKey)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Orders
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) GetOrder(id string) (*Order, error) {
	url := fmt.Sprintf(baseURL+"/Orders/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("api-key", c.APIKey)
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data Order
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) AddMarketOrder(order MarketOrder) error {
	url := fmt.Sprintf(baseURL + "/Orders/market")
	j, err := json.Marshal(order)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Set("api-key", c.APIKey)
	req.Header.Set("Content-Type", "application/json-patch+json")
	_, err = c.doRequest(req)
	return err
}

func (c *Client) AddLimitOrder(order LimitOrder) error {
	url := fmt.Sprintf(baseURL + "/Orders/limit")
	j, err := json.Marshal(order)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return err
	}
	req.Header.Set("api-key", c.APIKey)
	req.Header.Set("Content-Type", "application/json-patch+json")
	_, err = c.doRequest(req)
	return err
}

func (c *Client) CancelOrder(id string) error {
	url := fmt.Sprintf(baseURL+"/Orders/%s/Cancel", id)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("api-key", c.APIKey)
	_, err = c.doRequest(req)
	return err
}
