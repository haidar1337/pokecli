package main

import (
	"fmt"

	"github.com/haidar1337/pokecli/internal/api"
)

func commandMap(args ...string) error {
	config := api.GetConfig()
	results, err := api.FetchNextTwentyLocations(&config)

	if err != nil {
		return fmt.Errorf("%w: you are currently in the last page of the world, you may not go forward further", err)
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}