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
	GetAuthURL() (string, error)
	CreateClient() error
	SearchInDoc(id string, word string) error
	CreateFile(id string) error
}
