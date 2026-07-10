package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type pokeResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func getNamesofLocations(url string) (*pokeResponse, error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?limit=20&offset=0"
	}
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("something went wrong")
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("couldn't read response body")
	}

	var response pokeResponse
	if err := json.Unmarshal(data, &response); err != nil {
		log.Fatal(err)
	}

	return &response, nil
}
