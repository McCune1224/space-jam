package api

import (
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

func (stc *SpaceTradersClient) GetStatus() (*types.Status, error) {
	// This is the only endpoint that doesn't wrap the response in a data object
	statusResponse := types.Status{}
	err := stc.httpClient.Get(BASE_URL, nil, nil, &statusResponse)
	if err != nil {
		return nil, err
	}
	return &statusResponse, nil
}
