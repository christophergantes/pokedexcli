package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	fullURL    string
	httpClient http.Client
}

func (c *Client) getRequest(fullURL, urlPath string) ([]byte, error) {
	url := fullURL + urlPath
	res, err := c.httpClient.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("error running request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("error getting location area: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading body: %w", err)
	}
	return body, nil
}

func (c *Client) GetLocationAreas(pageURL *string) (NamedAPIResourceList, error) {
	url := c.fullURL + "location-area/"
	if pageURL != nil {
		url = *pageURL
	}
	body, err := c.getRequest(url, "")
	if err != nil {
		return NamedAPIResourceList{}, err
	}

	namedAPIResourceList := NamedAPIResourceList{}
	err = json.Unmarshal(body, &namedAPIResourceList)
	if err != nil {
		return NamedAPIResourceList{}, fmt.Errorf("error unmarshalling body: %w", err)
	}
	return namedAPIResourceList, nil
}

func (c *Client) GetLocationArea(identifier string) (LocationArea, error) {
	urlPath := "location-area/" + identifier
	body, err := c.getRequest(c.fullURL, urlPath)
	if err != nil {
		return LocationArea{}, err
	}

	locationArea := LocationArea{}
	err = json.Unmarshal(body, &locationArea)
	return locationArea, nil
}

func NewClient() *Client {
	return &Client{
		fullURL: "https://pokeapi.co/api/v2/",
		httpClient: http.Client{
			Timeout: 5 * time.Second,
		},
	}
}
