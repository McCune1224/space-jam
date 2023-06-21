package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Key routes for the SpaceTraders API.
const (
	base = "https://api.spacetraders.io"

	users = base + "/users"

	game          = base + "/game"
	gameShips     = game + "/ships"
	gameLoans     = game + "/loans"
	gameStatus    = game + "/status"
	gameSystems   = game + "/systems"
	gameLocations = game + "/locations"
)

type SpaceTradersClient struct {
	token    string
	username string
	client   http.Client
}

func RegisterNewAgent(email string, faction string, symbol string) error {
	url := "https://api.spacetraders.io/v2/register"

	payload := strings.NewReader("{\n  \"faction\": \"" + faction + "\",\n  \"symbol\": \"" + symbol + "\",\n  \"email\": \"" + email + "\"\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
	return nil
}

func New(token string, username string) *SpaceTradersClient {
	return &SpaceTradersClient{
		token:    token,
		username: username,
		client: http.Client{
			Transport: &http.Transport{},
			Timeout:   30 * time.Second,
		},
	}
}

// Internal method for making new requests.
func (stc *SpaceTradersClient) callEndpoint(method string, uri string,
	body string, headers map[string]string,
	urlParams map[string]string,
) (*http.Request, error) {
	bodyReader := strings.NewReader(body)

	req, err := http.NewRequest(method, uri, bodyReader)
	if err != nil {
		return req, err
	}

	for k, v := range headers {
		req.Header.Add(k, v)
	}

	urlQuery := req.URL.Query()
	for k, v := range urlParams {
		urlQuery.Add(k, v)
	}

	req.URL.RawQuery = urlQuery.Encode()

	return req, nil
}

// Internal method for making new requests with expected response.
func (stc *SpaceTradersClient) callEndpointShaped(method string, uri string,
	body string, headers map[string]string,
	urlParams map[string]string, response interface{},
) error {
	req, err := stc.callEndpoint(method, uri, body, headers, urlParams)
	if err != nil {
		return err
	}
	// Unmarshall response into response interface
	err = json.NewDecoder(req.Body).Decode(&response)
	if err != nil {
		return err
	}
	return nil
}

type status struct {
	Status      string `json:"status"`
	Version     string `json:"version"`
	ResetDate   string `json:"resetDate"`
	Description string `json:"description"`
	Stats       struct {
		Agents    int `json:"agents"`
		Ships     int `json:"ships"`
		Systems   int `json:"systems"`
		Waypoints int `json:"waypoints"`
	} `json:"stats"`
	Leaderboards struct {
		MostCredits []struct {
			AgentSymbol string `json:"agentSymbol"`
			Credits     int64  `json:"credits"`
		} `json:"mostCredits"`
		MostSubmittedCharts []struct {
			AgentSymbol string `json:"agentSymbol"`
			ChartCount  int    `json:"chartCount"`
		} `json:"mostSubmittedCharts"`
	} `json:"leaderboards"`
	ServerResets struct {
		Next      string `json:"next"`
		Frequency string `json:"frequency"`
	} `json:"serverResets"`
	Announcements []struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	} `json:"announcements"`
	Links []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"links"`
}

func (stc *SpaceTradersClient) ApiStatus() (status, error) {
	var s status
	err := stc.callEndpointShaped("GET", gameStatus, "", nil, nil, &s)
	if err != nil {
		return status{}, err
	}
	return s, nil
}
