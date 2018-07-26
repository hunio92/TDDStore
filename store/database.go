package store

import "sync"

type database struct {
	Clients       []Client
	Providers     map[string]Provider
	ProductsLists []ProductsInfo
	Stocks        map[string]Stock
	m             sync.Mutex
	minQty        int
}

// NewDatabase create new Database
func NewDatabase() *database {
	db := &database{
		Clients:       make([]Client, 0),
		Providers:     make(map[string]Provider),
		ProductsLists: make([]ProductsInfo, 0),
		Stocks:        make(map[string]Stock),
	}

	db.minQty = 100

	return db
}

func (db *database) AddClient(client Client) {
	db.Clients = append(db.Clients, client)
}

func (db *database) AddProvider(provider Provider) {
	if _, ok := db.Providers[provider.Name]; !ok {
		db.Providers[provider.Name] = provider
	}
}

func (db *database) AddProduct(product ProductsInfo) {
	db.m.Lock()
	defer db.m.Unlock()
	db.ProductsLists = append(db.ProductsLists, product)
}

func (db *database) AddStock(product ProductsInfo) {
	db.m.Lock()
	defer db.m.Unlock()
	if _, ok := db.Stocks[product.Name]; !ok {
		db.Stocks[product.Name] = Stock{ProductsInfo: product, PurchasePrice: product.Price}
	}
}

func (db *database) GetProductsBelowQty() []ProductsInfo {
	var belowStock = make([]ProductsInfo, 0)
	for _, product := range db.ProductsLists {
		if product.Qty < db.minQty {
			belowStock = append(belowStock, product)
		}
	}
	return belowStock
}
