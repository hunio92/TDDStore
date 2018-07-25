package store_test

import (
	"TDD_Store/store"
	"reflect"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	actual := store.NewDatabase()
	if actual == nil {
		t.Errorf("Expected to NOT to be nil: %v", actual)
	}
}

func TestAddClient(t *testing.T) {
	tests := []struct {
		in, out store.Client
	}{
		{store.Client{
			Name:         "LaDoiPasi",
			Addr:         "Principala",
			Money:        15.01,
			ShoppingList: []string{"leffe", "banane"},
		}, store.Client{
			Name:         "LaDoiPasi",
			Addr:         "Principala",
			Money:        15.01,
			ShoppingList: []string{"leffe", "banane"},
		}},
	}

	db := store.NewDatabase()

	for i, test := range tests {
		db.AddClient(test.in)
		if !reflect.DeepEqual(db.Client[i], test.out) {
			t.Errorf("Expected %v got %v", db.Client[i], test.out)
		}
	}
}
