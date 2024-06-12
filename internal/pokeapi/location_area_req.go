package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area?offset=0&limit=20"
	URL := baseURL + endpoint

	if pageURL != nil {
		URL = *pageURL
	}

	data, ok := c.cache.Get(URL)
	if ok {
		areas := LocationAreas{}
		err := json.Unmarshal(data, &areas)
		if err != nil {
			return LocationAreas{}, err
		}
		return areas, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreas{}, fmt.Errorf("bad status code; %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	areas := LocationAreas{}
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return LocationAreas{}, err
	}
	c.cache.Add(URL, data)

	return areas, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	URL := baseURL + endpoint

	data, ok := c.cache.Get(URL)
	if ok {
		area := LocationArea{}
		err := json.Unmarshal(data, &area)
		if err != nil {
			return LocationArea{}, err
		}
		return area, nil
	}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationArea{}, fmt.Errorf("bad status code; %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	area := LocationArea{}
	err = json.Unmarshal(data, &area)
	if err != nil {
		return LocationArea{}, err
	}
	c.cache.Add(URL, data)

	return area, nil
}
