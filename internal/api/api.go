package api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Result struct {
	Name string
	Url string
}

type LocationArea struct {
	Count int `json:"count"`
	Next *string
	Previous *string
	Results []Result
}

func fetchLocationAreas(config Config) error {
	apiUrl := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(apiUrl)
	if err != nil {
		return errors.New("Unable to fetch location areas")
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode >= 300 {
		errorMessage := fmt.Sprintf("Response failed with status code: %d", res.StatusCode)
		return errors.New(errorMessage)
	}

	if err != nil {
		return err
	}

	// parse body from json to go and update config
	// find out why structs need to have "json:" tag
	// import encoded/json
	fmt.Println(body, config.Next, config.Previous)

	return nil
}

func GetNextTwentyLocations()

func GetPreviousTwentyLocations()