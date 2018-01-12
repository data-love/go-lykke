package lykke

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AssetPairs []AssetPair

type AssetPair struct {
	ID               string  `json:"Id"`
	Name             string  `json:"IsBuy"`
	Accuracy         float32 `json:"Accuracy"`
	InvertedAccuracy int     `json:"InvertedAccuracy"`
	BaseAssetID      string  `json:"BaseAssetId"`
	QuotingAssetID   string  `json:"QuotingAssetId"`
}

func (s *Client) GetAssetPairs() (*AssetPairs, error) {
	url := fmt.Sprintf(baseURL + "/AssetPairs")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data AssetPairs
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (s *Client) GetAssetPair(id string) (*AssetPair, error) {
	url := fmt.Sprintf(baseURL+"/AssetPairs/%s", id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := s.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data AssetPair
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
