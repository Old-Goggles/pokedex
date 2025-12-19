package main

import "fmt"

func commandMapB(cfg *config, args ...string) error {
	// 1. If there is no previous page, print message and return
	if cfg.previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	// 2. Use cfg.previous as the page URL
	pageURL := cfg.previous

	resp, err := cfg.pokeClient.ListLocations(&pageURL)
	if err != nil {
		return err
	}

	// 3. Print the locations
	for _, loc := range resp.Results {
		fmt.Println(loc.Name)
	}

	// 4. Update next/previous from the response
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
