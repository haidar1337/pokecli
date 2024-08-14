package main

import (
	"fmt"

	"github.com/haidar1337/pokecli/internal/api"
)

func commandMapb(args ...string) error {
	config := api.GetConfig()
	results, err := api.FetchPreviousTwentyLocations(&config)


	if err != nil {
		return fmt.Errorf("%w: you are currently in the first page of the world, you may not go back further", err)
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}

	return nil
}