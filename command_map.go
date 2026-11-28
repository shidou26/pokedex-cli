package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func commandMap(cfg *config, argv []string) error {
	if len(argv) != 0 {
		return fmt.Errorf("usage: map")
	}
	if len(cfg.locationAreaNext) == 0 {
		return fmt.Errorf("you're on the last page")
	}

	var data LocationAreas

	entry, ok := cfg.commandCache.Get(cfg.locationAreaNext)
	if ok {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return err
		}
	} else {
		resp, err := http.Get(cfg.locationAreaNext)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code: %v", resp.StatusCode)
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return err
		}
		cfg.commandCache.Add(cfg.locationAreaNext, bytes)
	}

	for _, c := range data.Results {
		fmt.Println(c.Name)
	}

	cfg.locationAreaPrevious = data.Previous
	cfg.locationAreaNext = data.Next

	return nil
}

func commandMapBack(cfg *config, argv []string) error {
	if len(argv) != 0 {
		return fmt.Errorf("usage: mapb")
	}
	if len(cfg.locationAreaPrevious) == 0 {
		return fmt.Errorf("you're on the first page")
	}

	var data LocationAreas

	entry, ok := cfg.commandCache.Get(cfg.locationAreaPrevious)
	if ok {
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return err
		}
	} else {
		resp, err := http.Get(cfg.locationAreaPrevious)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("status code: %v", resp.StatusCode)
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &data)
		if err != nil {
			return err
		}
		cfg.commandCache.Add(cfg.locationAreaPrevious, bytes)
	}

	for _, c := range data.Results {
		fmt.Println(c.Name)
	}

	cfg.locationAreaNext = data.Next
	cfg.locationAreaPrevious = data.Previous
	return nil
}
