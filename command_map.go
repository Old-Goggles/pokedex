package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	var pageURL *string
	if cfg.next != "" {
		pageURL = &cfg.next
	}

	resp, err := pokeClient.ListLocations(pageURL)
	if err != nil {
		return err
	}

	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	if resp.Next != nil {
		cfg.next = *resp.Next
	} else {
		cfg.next = ""
	}

	if resp.Previous != nil {
		cfg.previous = *resp.Previous
	} else {
		cfg.previous = ""
	}

	return nil
}
