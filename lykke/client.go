package lykke

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://hft-api.lykke.com/api"

// Client to use apiKey
type Client struct {
	APIKey string
}

func NewApiClient(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

type IsAlive struct {
	Version         string            `json:"Version"`
	Env             string            `json:"Env"`
	IsDebug         bool              `json:"IsDebug"`
	IssueIndicators *[]IssueIndicator `json:"IssueIndicators"`
}

type IssueIndicator struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

// GetIsAlive to check if Api is Alive
func (c *Client) GetIsAlive() (*IsAlive, error) {
	url := fmt.Sprintf(baseURL + "/isAlive")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var data IsAlive
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if 200 != resp.StatusCode {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
