package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area"
	URL := baseURL + endpoint

	if pageURL != nil {
		URL = *pageURL
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	areas := LocationAreas{}
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return LocationAreas{}, err
	}

	return areas, nil
}
