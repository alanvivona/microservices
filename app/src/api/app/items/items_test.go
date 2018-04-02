package items

import (
	"api/app/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"api/app/models"

	"github.com/gin-gonic/gin"
)

func ConfigureRouter() *gin.Engine {
	router := gin.Default()
	Configure(router, nil)
	return router
}

func TestGetItem(t *testing.T) {
	router := ConfigureRouter()

	var is mock.ItemService
	Is = &is

	is.ItemFn = func(id int) (*models.Item, error) {
		return &models.Item{ID: id, Name: "Test result name", Description: "Test result description"}, nil
	}

	testCases := []struct {
		method                string
		url                   string
		shouldServiceBeCalled bool
	}{
		{"GET", "/item/1", true},
		{"GET", "/item/-1", false},
		{"GET", "/item/0", false},
		{"GET", "/item/000001", true},
		{"GET", "/item/99999999999999999999", false},
		{"GET", "/item/a", false},
	}

	for _, testCase := range testCases {
		is.ItemInvoked = false
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(testCase.method, testCase.url, nil)
		router.ServeHTTP(w, r)
		if is.ItemInvoked != testCase.shouldServiceBeCalled {
			t.Fatal("FAILED - Item Controller :: TestGetItem", testCase.method, testCase.url)
		}
	}
}

func TestGetItems(t *testing.T) {
	router := ConfigureRouter()

	var is mock.ItemService
	Is = &is

	is.ItemsFn = func() ([]*models.Item, error) {
		return nil, nil
	}

	testCases := []struct {
		method                string
		url                   string
		shouldServiceBeCalled bool
	}{
		{"GET", "/item", true},
		{"GET", "/item/0", false},
	}

	for _, testCase := range testCases {
		is.ItemsInvoked = false
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(testCase.method, testCase.url, nil)
		router.ServeHTTP(w, r)
		if is.ItemsInvoked != testCase.shouldServiceBeCalled {
			t.Fatal("FAILED - Item Controller :: TestGetItems", testCase.method, testCase.url)
		}
	}
}

func TestPostItem(t *testing.T) {
	router := ConfigureRouter()

	var is mock.ItemService
	Is = &is

	is.CreateItemFn = func(*models.Item) error {
		return nil
	}

	testCases := []struct {
		method                string
		url                   string
		body                  string
		shouldServiceBeCalled bool
	}{
		{"POST", "/item/", "{", false},
		{"POST", "/item/", "{\"fakeproperty\":\"value\"}", false},
		{"POST", "/item/", "{\"name\":\"test name\"}", false},
		{"POST", "/item/", "{\"name\":\"test name\", \"fakeproperty\": \"value\"}", false},
		{"POST", "/item/", "{\"name\":\"test name\", \"description\": \"test desc\"}", true},
		{"POST", "/item/", "{\"name\":\"test name ---- really long name that does not meet the criteria\", \"description\": \"test desc\"}", false},
		{"POST", "/item/", "{\"name\":\"test name\", \"description\": \"test desc ---- really long desc that does not meet the criteria\"}", false},
	}

	for _, testCase := range testCases {
		is.CreateItemInvoked = false
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(testCase.method, testCase.url, strings.NewReader(testCase.body))
		router.ServeHTTP(w, r)
		if is.CreateItemInvoked != testCase.shouldServiceBeCalled {
			t.Fatal("FAILED - Item Controller :: TestPostItem", testCase.method, testCase.url, testCase.body)
		}
	}
}

func TestDeleteItem(t *testing.T) {
	router := ConfigureRouter()

	var is mock.ItemService
	Is = &is

	is.DeleteItemFn = func(int) error {
		return nil
	}

	testCases := []struct {
		method                string
		url                   string
		shouldServiceBeCalled bool
	}{
		{"DELETE", "/item/1", true},
		{"DELETE", "/item/-1", false},
		{"DELETE", "/item/0", false},
		{"DELETE", "/item/99999999999999999999", false},
		{"DELETE", "/item/001", true},
		{"DELETE", "/item/a", false},
		{"DELETE", "/item/", false},
	}

	for _, testCase := range testCases {
		is.DeleteItemInvoked = false
		w := httptest.NewRecorder()
		r, _ := http.NewRequest(testCase.method, testCase.url, nil)
		router.ServeHTTP(w, r)
		if is.DeleteItemInvoked != testCase.shouldServiceBeCalled {
			t.Fatal("FAILED - Item Controller :: TestDeleteItem", testCase.method, testCase.url)
		}
	}
}
