package lykke

import (
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

func (s *Client) GetOrders() (*Orders, error) {
	url := fmt.Sprintf(baseURL + "/Orders")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
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

func (s *Client) GetOrder(id string) (*Order, error) {
	url := fmt.Sprintf(baseURL+"/Orders/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
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
