package tests

import (
	"bytes"
	"net/http"
	"testing"
)

func TestIntegration(t *testing.T) {
	t.Log("EXECUTING Test cases")

	testCases := []struct {
		method       string
		url          string
		body         string
		expectedCode int
		expectedBody string
	}{
		// Testing Item URLs availability
		{"GET", "http://localhost:8080/", "", http.StatusNotFound, ""},
		{"GET", "http://localhost:8080/item", "", http.StatusOK, ""},
		{"GET", "http://localhost:8080/item/1", "", http.StatusOK, ""},

		// Testing Gdrive URLs availability
		{"GET", "http://localhost:8080/gdrive", "", http.StatusNotFound, ""},
		{"GET", "http://localhost:8080/gdrive/auth", "", http.StatusOK, ""},
		{"GET", "http://localhost:8080/gdrive/search-in-doc", "", http.StatusNotFound, ""},

		// Testing Item URLs functionality
		{"GET", "http://localhost:8080/item/1", "", http.StatusOK, "{\"id\":1,\"name\":\"new item name\",\"description\":\"new item description\"}"},
		{"POST", "http://localhost:8080/item/", "{\"name\":\"integration test item name\",\"description\":\"integration test item description\"}", http.StatusCreated, "{\"id\":22,\"name\":\"integration test item name\",\"description\":\"integration test item description\"}"},
		{"GET", "http://localhost:8080/item/22", "", http.StatusOK, "{\"id\":22,\"name\":\"integration test item name\",\"description\":\"integration test item description\"}"},
	}

	t.Log("TOTAL = ", len(testCases))

	for i, testCase := range testCases {
		t.Log("====================")
		t.Log("EXECUTING Test case ", i+1, " of ", len(testCases), " :", testCase.url, testCase.method)
		if testCase.method != "" {
			var (
				resp *http.Response
				err  error
			)
			switch testCase.method {
			case "GET":
				resp, err = http.Get(testCase.url)
			case "POST":
				body := []byte(testCase.body)
				resp, err = http.Post(testCase.url, "application/json", bytes.NewBuffer(body))
			}

			if err != nil {
				t.Error("Test case FAILED: Request Error", testCase.url, testCase.method, testCase.body, testCase.expectedCode, testCase.expectedBody, err.Error())
				t.Fail()
			} else {
				if resp.StatusCode != testCase.expectedCode {
					t.Error("Test case FAILED - Expected Code: ", testCase.url, testCase.method, testCase.body, testCase.expectedCode, resp.StatusCode, testCase.expectedBody)
					t.Fail()
				} else {
					t.Log("Test case PASSED", testCase.url, testCase.method, testCase.body, testCase.expectedCode, resp.StatusCode, testCase.expectedBody)
				}
			}
		} else {
			t.Error("Test case FAILED - Configuration Issue: ", testCase.url, testCase.method, testCase.body, testCase.expectedCode, testCase.expectedBody)
			t.Fail()
		}
		t.Log("====================")
	}
}
