package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"time"
)

type Client struct {
	Name         string
	Addr         string
	Money        float64
	ShoppingList []string
}

type ProductsInfo struct {
	Name         string
	Price        float64
	Expires      time.Time
	Qty          int
	ProviderName string
}

type Provider struct {
	Name         string
	Addr         string
	ProductsList []ProductsInfo
}

type Stock struct {
	ProductsInfo
	PurchasePrice float64
}

type service struct {
	db         *database
	transport  ProviderApi
	HostAddr   string
	OrderRoute string
}

// NewStore create a new Store obj
func NewStore(db *database) *service {
	transport := NewProviderApi()
	return &service{
		db:        db,
		transport: transport,
	}
}

func (s *service) VerifyAmount() {
	tick := time.Tick(5 * time.Second)
	for range tick {
		for _, product := range s.db.GetProductsBelowQty() {
			if product.ProviderName != "" {
				order, err := s.orderProduct(product)
				if err != nil {
					fmt.Printf("Order ERROR: %v", err.Error())
				}
				if err == nil {
					s.db.AddStock(order)
				}
			}
		}
	}
}

func (s *service) orderProduct(product ProductsInfo) (ProductsInfo, error) {
	purchaseOrder := ProductsInfo{
		Name: product.Name,
		Qty:  30,
	}
	var order ProductsInfo

	jsonBytes, err := json.Marshal(purchaseOrder)
	if err != nil {
		return order, err
	}

	resp, err := s.transport.SendOrder(jsonBytes, s.db.Providers[product.Name].Addr)
	if err != nil {
		return order, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &order)

	return order, nil
}

func (s *service) VerifyPrices() {
	tick := time.Tick(4 * time.Second)
	for range tick {
		for _, product := range s.db.ProductsLists {
			s.getCheapestPriceFromProviders(product)
			// replace lowest price
		}
	}
}

func (s *service) getCheapestPriceFromProviders(product ProductsInfo) (float64, error) {
	var lowestPrice = math.MaxFloat64
	var err error
	for _, provider := range s.db.Providers {
		if s.IsProductInProvidersStock(product.Name, provider) {
			lowestPrice, err = s.transport.GetPrice(product.Name, provider.Addr)
			if err != nil {
				return 0, err
			}
		}
	}

	return lowestPrice, nil
}

func (s *service) IsProductInProvidersStock(seedProduct string, provider Provider) bool {
	for _, product := range provider.ProductsList {
		if product.Name == seedProduct {
			return true
		}
	}
	return false
}
