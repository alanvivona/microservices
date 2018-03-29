package items

import (
	"api/app/models"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func executeBeforeTest() (*sql.DB, bool) {
	// Connect to the DB
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", "user", "userpwd", "db", "3306", "db"))
	if err != nil {
		panic("Could not connect to the db")
		return db, true
	}
	s := &ItemService{
		DB: db,
	}
	res, err := s.DB.Exec(`TRUNCATE items`)
	if err != nil {
		panic("Could not truncate items table")
		return db, true
	}
	defer stmt.Close()
	return db, false
}

func TestItemService_Items(t *testing.T) {

	db, err := executeBeforeTest()
	if err == true {
		panic("executeBeforeTest execution failed")
		return
	}

	// Set some items
	s := &ItemService{
		DB: db,
	}
	for i := 0; i < 3; i++ {
		err := s.CreateItem()
		if err != nil {
			panic("Could not create item: ", err)
		}
	}

	type fields struct {
		DB *sql.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*models.Item
		wantErr bool
	}{
		{
			name:    "Should return a list of items from the DB",
			fields:  {DB: db},
			want:    {},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ItemService{
				DB: tt.fields.DB,
			}
			got, err := s.Items()
			if (err != nil) != tt.wantErr {
				t.Errorf("ItemService.Items() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItemService.Items() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestItemService_CreateItem(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		i *models.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ItemService{
				DB: tt.fields.DB,
			}
			if err := s.CreateItem(tt.args.i); (err != nil) != tt.wantErr {
				t.Errorf("ItemService.CreateItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestItemService_DeleteItem(t *testing.T) {
	type fields struct {
		DB *sql.DB
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ItemService{
				DB: tt.fields.DB,
			}
			if err := s.DeleteItem(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("ItemService.DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
