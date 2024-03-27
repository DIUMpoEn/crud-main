package test

import (
	"crud/models"
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"testing"
)

var url1 = "http://127.0.0.1:5555"

func TestGetUser(t *testing.T) {
	output, err := GetUser("3")
	if err != nil {
		t.Fatalf("failed getting user: %v", err)
	}
	log.Printf("Output: %v", output)
}

func GetUser(id string) ([]models.Account, error) {
	output := []models.Account{}

	url := fmt.Sprintf("%s/get/%s", url1, id)
	log.Printf("URL: %v", url)
	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetResult(&output).
		Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("status code wrong. status: %d. body: %s", resp.StatusCode(), resp.String())
	}

	return output, nil
}
