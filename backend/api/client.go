package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/McCune1224/space-jam/types"
)

// Wrapper struct to encapsulate a generic struct into the data object in the response.
type ApiDataResponse[ResponseType any] struct {
	Data ResponseType `json:"data"`
}

// Wrapper struct to encapsulate a generic struct into the data object in the response.
type PaginatedApiDataResponse[ResponseType any] struct {
	Data []ResponseType `json:"data"`
	Meta types.Meta     `json:"meta"`
}

// Struct that is used as parameters for requests that use pagination i.e ListContracts.
type PaginationOptions struct {
	Limit int
	Page  int
}

// Convert struct into kv pair map for query params.
func (po *PaginationOptions) ToQueryParams() map[string]string {
	queryParams := make(map[string]string)
	if po.Limit != 0 {
		queryParams["limit"] = fmt.Sprintf("%d", po.Limit)
	}
	if po.Page != 0 {
		queryParams["page"] = fmt.Sprintf("%d", po.Page)
	}

	if len(queryParams) == 0 {
		return nil
	}
	return queryParams
}

// Helper CRUD http client for the SpaceTraders API.
type crudClient struct {
	client http.Client
}

// Create a new CRUD client.
func newCrudClient() *crudClient {
	return &crudClient{
		client: http.Client{
			Transport: &http.Transport{},
			Timeout:   30 * time.Second,
		},
	}
}

// Get a resource from the SpaceTraders API. If no response type is provided,
// the response will be printed to std out as a string.
func (cc *crudClient) Get(url string, headers map[string]string, queryParams map[string]string, responseBody ...interface{}) error {
	requestURL := url
	if queryParams != nil {
		requestURL += "?"
		for key, value := range queryParams {
			requestURL += fmt.Sprintf("%s=%s&", key, value)
		}
		requestURL = requestURL[:len(requestURL)-1]
	}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return err
	}

	// Headers to add
	// Web API so these are just defaults
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	// Add any additional headers (e.g. Authiriization)
	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := cc.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// if no response type is provided, print the response
	if len(responseBody) == 0 {
		fmt.Println(string(body))
		return nil
	}

	// check for non-200 response
	if resp.StatusCode != 200 {
		return fmt.Errorf("Non-200 response code: %d", resp.StatusCode)
	}

	err = json.Unmarshal(body, responseBody[0])
	if err != nil {
		return err
	}
	return nil
}

func (cc *crudClient) Post() {
}
func (cc *crudClient) Put()    {}
func (cc *crudClient) Delete() {}
