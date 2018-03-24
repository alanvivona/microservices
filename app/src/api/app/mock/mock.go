package mock

import "api/app/models"

// ItemService ...
type ItemService struct {
	ItemFn      func(id string) (*models.Item, error)
	ItemInvoked bool

	ItemsFn      func() ([]*models.Item, error)
	ItemsInvoked bool

	CreateItemFn      func(i *models.Item) error
	CreateItemInvoked bool

	DeleteItemFn      func(id string) error
	DeleteItemInvoked bool
}

// Item ...
func (is *ItemService) Item(id string) (*models.Item, error) {
	is.ItemInvoked = true
	return is.ItemFn(id)
}

// Items ...
func (is *ItemService) Items() ([]*models.Item, error) {
	is.ItemsInvoked = true
	return is.ItemsFn()
}

// CreateItem ...
func (is *ItemService) CreateItem(i *models.Item) error {
	is.CreateItemInvoked = true
	return is.CreateItemFn(i)
}

// DeleteItem ...
func (is *ItemService) DeleteItem(id string) error {
	is.DeleteItemInvoked = true
	return is.DeleteItemFn(id)
}
