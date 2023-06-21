package api

import (
	"fmt"
	"testing"
)

var testCrudClient *crudClient

func crudClientSetup() {
	testCrudClient = newCrudClient()
}

func crudClientTeardown() {}

func TestGet(t *testing.T) {
	crudClientSetup()
	testCases := []struct {
		url         string
		headers     map[string]string
		queryParams map[string]string
		bodyDump    interface{}
	}{
		// Endpoint with only url
		{
			"https://api.chucknorris.io/jokes/random",
			nil,
			nil,
			nil,
		},
		// Endpoint with url and query params
		{
			"https://api.chucknorris.io/search",
			nil,
			map[string]string{"query": "dev"},
			nil,
		},
		// Status Endpoint for SpaceTraders
		{
			BASE_URL,
			map[string]string{"Authorization": fmt.Sprintf("Bearer %s", testToken)},
			nil,
			&struct {
				Status  string `json:"status"`
				Message string `json:"message"`
			}{},
		},
	}
	for _, testCase := range testCases {
		if testCase.bodyDump != nil {
			err := testCrudClient.Get(testCase.url, testCase.headers, testCase.queryParams, testCase.bodyDump)
            fmt.Println(testCase.bodyDump)
			if err != nil {
				t.Errorf("Error getting %s: %v", testCase.url, err)
			}
			continue
		}
		err := testCrudClient.Get(testCase.url, nil, testCase.queryParams)
		if err != nil {
			t.Errorf("Error getting %s: %v", testCase.url, err)
		}
	}
}

func TestPost(t *testing.T) {
}

// TODO: Implement tests for PUT and DELETE.
// func TestPut(t *testing.T) {
// }

// func TestDelete(t *testing.T) {
//
// }
