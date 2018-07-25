package store

import "sync"

type database struct {
	Client       []Client
	Provider     map[string]Provider
	ProductsList []ProductsInfo
	Stock        map[string]Stock
	m            sync.Mutex
}

// NewDatabase create new Database
func NewDatabase() *database {
	db := &database{
		Client:       make([]Client, 0),
		Provider:     make(map[string]Provider),
		ProductsList: make([]ProductsInfo, 0),
		Stock:        make(map[string]Stock),
	}

	return db
}

func (db *database) AddClient(client Client) {
	db.Client = append(db.Client, client)
}
