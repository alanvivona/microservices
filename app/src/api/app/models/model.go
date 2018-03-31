package models

import (
	"github.com/gin-gonic/gin"
	drive "google.golang.org/api/drive/v3"
)

// Item ...
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// File ...
type File struct {
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
	HasClient() bool
	GetAuthURL() (string, error)
	CreateClient(*gin.Context, string) error
	SearchInDoc(id string, word string) (bool, error)
	CreateFile(f *File) (*drive.File, error)
}
