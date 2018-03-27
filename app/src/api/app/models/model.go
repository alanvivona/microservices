package models

// Item ...
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ItemServiceInterface ...
type ItemServiceInterface interface {
	Item(id int) (*Item, error)
	Items() ([]*Item, error)
	CreateItem(i *Item) error
	DeleteItem(id int) error
}

// GdriveServiceInterface ...
type GdriveServiceInterface interface {
	SearchInDoc(id string, word string) error
	SearchDoc(id string) error
	CreateFile(i *Item) error
}
