package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/McCune1224/space-jam/types"
)

func RegisterNewAgent(email string, faction string, symbol string) error {
	url := "https://api.spacetraders.io/v2/register"

	// TODO: Make this a struct you can pass in to fill in the values, default if struct with neccessary values is not passed in or missing...
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

func (stc *SpaceTradersClient) GetAgent() (ApiDataResponse[types.Agent], error) {
	var agentResponse struct {
		Data types.Agent `json:"data"`
	}
	url := BASE_URL + "/my/agent"
	err := stc.httpClient.Get(url, stc.authHeader, nil, &agentResponse)
	if err != nil {
		return agentResponse, err
	}
	return agentResponse, nil
}
