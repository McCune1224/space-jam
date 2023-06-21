package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/McCune1224/space-jam/types"
)

// Key routes for the SpaceTraders API.
const (
	BASE_URL = "https://api.spacetraders.io/v2"

	USERS_URL = BASE_URL + "/users"

	GAME_URL           = BASE_URL + "/game"
	GAME_SHIPS_URL     = GAME_URL + "/ships"
	GAME_LOANS_URL     = GAME_URL + "/loans"
	GAME_SYSTEMS_URL   = GAME_URL + "/systems"
	GAME_LOCATIONS_URL = GAME_URL + "/locations"
)

// API client for the SpaceTraders API.
// Provides methods for interacting with the API.
type SpaceTradersClient struct {
	token      string
	username   string
	httpClient *crudClient
	authHeader map[string]string
}

// Create a new SpaceTradersClient.
func NewSpaceTraderClient(token string, username string) *SpaceTradersClient {
	return &SpaceTradersClient{
		token:      token,
		username:   username,
		httpClient: newCrudClient(),
		authHeader: map[string]string{
			"Authorization": "Bearer " + token,
		},
	}
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

func (stc *SpaceTradersClient) GetStatus() (*types.Status, error) {
	statusResponse := types.Status{}
	err := stc.httpClient.Get(BASE_URL, nil, nil, &statusResponse)
	if err != nil {
		return nil, err
	}
	return &statusResponse, nil
}
