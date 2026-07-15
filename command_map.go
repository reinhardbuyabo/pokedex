package main

import (
	"fmt"
	"log"
)

func commandMap(cfg *config) error {
	url := ""

	if cfg.nextLocationsURL != nil {
		url = *cfg.nextLocationsURL
	}

	data, err := getNamesofLocations(url)
	if err != nil {
		log.Fatal(err)
	}
	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}

	cfg.nextLocationsURL = data.Next
	cfg.prevLocationsURL = data.Previous

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationsURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	data, err := getNamesofLocations(*cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	for _, locationArea := range data.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
