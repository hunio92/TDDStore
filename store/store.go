package store

import (
	"time"
)

type Client struct {
	Name         string
	Addr         string
	Money        float64
	ShoppingList []string
}

type ProductsInfo struct {
	Name    string
	Price   float64
	Expires time.Time
	Qty     int
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
	s *database
}

func main() {

}
