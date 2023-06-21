package api

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

var (
	testUsername = ""
	testToken    = ""
	testTrader   *SpaceTradersClient
)

func spaceTraderClientSetup() {
	testUsername = os.Getenv("ST_USERNAME")
	testToken = os.Getenv("ST_TOKEN")
	if testUsername == "" || testToken == "" {
		panic("failed to load env variables")
	}
	testTrader = NewSpaceTraderClient(testToken, testUsername)
}

func spaceTraderClientTeardown() {
}

func TestGetStatus(t *testing.T) {
	spaceTraderClientSetup()
	_, err := testTrader.GetStatus()
	if err != nil {
		t.Error(err)
	}
}

func TestGetAgent(t *testing.T) {
	spaceTraderClientSetup()
	agentResponse, err := testTrader.GetAgent()
	if err != nil {
		t.Error(err)
	}
    fmt.Println(agentResponse)
}
