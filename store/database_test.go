package store_test

import (
	"TDD_Store/store"
	"reflect"
	"testing"
	"time"
)

func TestAddClient(t *testing.T) {
	tests := []struct {
		in, out store.Client
	}{
		{store.Client{
			Name:         "Pista",
			Addr:         "Principala38",
			Money:        15.01,
			ShoppingList: []string{"leffe", "banane"},
		}, store.Client{
			Name:         "Pista",
			Addr:         "Principala38",
			Money:        15.01,
			ShoppingList: []string{"leffe", "banane"},
		}},
	}

	db := store.NewDatabase()

	for i, test := range tests {
		db.AddClient(test.in)
		if !reflect.DeepEqual(db.Client[i], test.out) {
			t.Errorf("Expected %v, but got %v", db.Client[i], test.out)
		}
	}
}

func TestAddProvider(t *testing.T) {
	tests := []struct {
		in, out store.Provider
	}{
		{store.Provider{
			Name: "LaDoiPasi",
			Addr: "Principala",
			ProductsList: []store.ProductsInfo{
				{
					Name:    "Ciuc",
					Price:   3.50,
					Expires: time.Now(),
					Qty:     10,
				},
			},
		}, store.Provider{
			Name: "LaDoiPasi",
			Addr: "Principala",
			ProductsList: []store.ProductsInfo{
				{
					Name:    "Ciuc",
					Price:   3.50,
					Expires: time.Now(),
					Qty:     10,
				},
			},
		}},
	}

	db := store.NewDatabase()

	for _, test := range tests {
		db.AddProvider(test.in)
		if !reflect.DeepEqual(db.Provider[test.in.Name], test.out) {
			t.Errorf("Expected %v, but got %v", db.Provider[test.in.Name], test.out)
		}
	}
}
