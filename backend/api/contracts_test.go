package api

import (
	"fmt"
	"testing"
)

func TestListContracts(t *testing.T) {
	spaceTraderClientSetup()
	got, err := testTrader.ListContracts()
	if err != nil {
		t.Error(err)
	}
	if len(got.Data) == 0 {
		t.Error("Got nil, expected non-nil")
	}
	fmt.Println(got)
}

func TestGetContract(t *testing.T) {}

func TestAcceptContract(t *testing.T) {}

func TestDeliverCargoToContract(t *testing.T) {}

func TestFufillContract(t *testing.T) {}
