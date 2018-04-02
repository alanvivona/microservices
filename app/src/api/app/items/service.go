package items

import (
	"api/app/models"
	"database/sql"
	"html"
)

// ItemService ...
type ItemService struct {
	DB *sql.DB
}

// Item ...
func (s *ItemService) Item(id int) (*models.Item, error) {
	var i models.Item
	row := s.DB.QueryRow(`SELECT id, name, description FROM items WHERE id = ?`, id)
	if err := row.Scan(&i.ID, &i.Name, &i.Description); err != nil {
		return nil, err
	}
	return &i, nil
}

// Items ...
func (s *ItemService) Items() ([]*models.Item, error) {
	var items []*models.Item
	rows, err := s.DB.Query(`SELECT id, name, description FROM items`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var item models.Item
		err := rows.Scan(&item.ID, &item.Name, &item.Description)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return items, err
}

// CreateItem ...
func (s *ItemService) CreateItem(i *models.Item) error {
	stmt, err := s.DB.Prepare(`INSERT INTO items(name,description) values(?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(html.EscapeString(i.Name), html.EscapeString(i.Description))
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	i.ID = int(id)
	return nil
}

// DeleteItem ...
func (s *ItemService) DeleteItem(id int) error {
	stmt, err := s.DB.Prepare(`DELETE FROM items WHERE id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
