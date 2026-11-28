package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LocationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	if len(cfg.locationAreaNext) == 0 {
		return fmt.Errorf("you're on the last page")
	}

	resp, err := http.Get(cfg.locationAreaNext)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %v", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	var areas LocationAreas
	if err := dec.Decode(&areas); err != nil {
		return err
	}

	for _, c := range areas.Results {
		fmt.Println(c.Name)
	}

	cfg.locationAreaPrevious = cfg.locationAreaNext
	cfg.locationAreaNext = areas.Next
	return nil
}

func commandMapBack(cfg *config) error {
	if len(cfg.locationAreaPrevious) == 0 {
		return fmt.Errorf("you're on the first page")
	}

	resp, err := http.Get(cfg.locationAreaPrevious)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %v", resp.StatusCode)
	}

	dec := json.NewDecoder(resp.Body)
	var areas LocationAreas
	if err := dec.Decode(&areas); err != nil {
		return err
	}

	for _, c := range areas.Results {
		fmt.Println(c.Name)
	}

	cfg.locationAreaNext = cfg.locationAreaPrevious
	cfg.locationAreaPrevious = areas.Previous
	return nil
}
