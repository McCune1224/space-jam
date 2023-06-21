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
	// err := godotenv.auto(".test.env")
	// if err != nil {
	// 	panic(err)
	// }
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
	resp, err := testTrader.GetStatus()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(resp)
}

// func TestGetUser(t *testing.T) {
// 	resp, err := testTrader.GetUser()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(resp)
// }
