package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("Empty campaignID")
	}

	if balance <= 0.0 {
		return nil, fmt.Errorf("balance must be greater than zero")
	}

	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("Time can not be less than present time")
	}

	return &Budget{campaignID, balance, expires}, nil
}

func main() {
	expires := time.Now().Add(7 * 24 * time.Hour)

	B1, err := NewBudget("Cars", 234.89, expires)

	if err != nil {
		fmt.Println("ERROR:", err)
	} else {
		fmt.Printf("%#v", B1)
	}
}
