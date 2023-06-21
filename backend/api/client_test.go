package api

import (
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

var (
	username = ""
	token    = ""
	trader   *SpaceTradersClient
)

func setup() {
	// err := godotenv.auto(".test.env")
	// if err != nil {
	// 	panic(err)
	// }
	username = os.Getenv("ST_USERNAME")
	token = os.Getenv("ST_TOKEN")
	if username == "" || token == "" {
		panic("failed to load env variables")
	}
}

func teardown() {
}

func TestMain(m *testing.M) {
	setup()
	trader = New(token, username)
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestApiStatus(t *testing.T) {
	resp, err := trader.ApiStatus()
	if err != nil {
		t.Error(err)
	}
	if resp.Status != "spacetraders is currently online and available to play" {
		t.Error("Unexpected status message")
	}
}

// func TestRegisterNewAgent(t *testing.T) {
// 	err := RegisterNewAgent("alexmccune1224@gmail.com", "COSMIC", "KUSA")
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }
