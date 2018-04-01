package mock

import "api/app/models"

// ItemService ...
type ItemService struct {
	ItemFn      func(id int) (*models.Item, error)
	ItemInvoked bool

	ItemsFn      func() ([]*models.Item, error)
	ItemsInvoked bool

	CreateItemFn      func(i *models.Item) error
	CreateItemInvoked bool

	DeleteItemFn      func(id int) error
	DeleteItemInvoked bool
}

// Item ...
func (is *ItemService) Item(id int) (*models.Item, error) {
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
func (is *ItemService) DeleteItem(id int) error {
	is.DeleteItemInvoked = true
	return is.DeleteItemFn(id)
}
