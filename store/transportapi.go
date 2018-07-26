package store

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ProviderApi interface { // rename ProviderApi
	SendOrder([]byte, string) (*http.Response, error)
	GetPrice(product, providerAddr string) (float64, error)
}

type transport struct {
	HostAddr   string
	OrderRoute string
}

func NewProviderApi() *transport {
	return &transport{
		HostAddr:   "http://localhost:",
		OrderRoute: "/orders",
	}
}

func (trans *transport) SendOrder(rawData []byte, providerAddr string) (*http.Response, error) {
	jsonReader := bytes.NewReader(rawData)
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(trans.HostAddr+providerAddr+trans.OrderRoute, "json", jsonReader)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp, nil
}

func (trans *transport) GetPrice(product, providerAddr string) (float64, error) {
	purchaseOrder := ProductsInfo{
		Name: product,
	}
	var order ProductsInfo

	jsonBytes, err := json.Marshal(purchaseOrder)
	if err != nil {
		return 0, err
	}

	resp, err := trans.SendOrder(jsonBytes, providerAddr)
	if err != nil {
		return 0, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &order)

	return order.Price, nil // hmmm ?! marshall or not ?!
}
