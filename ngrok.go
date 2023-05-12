package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getNgrokTunnels() ([]Tunnel, error) {
	// Create a new http client
	client := &http.Client{}

	// Make a request to the ngrok API
	resp, err := client.Get("http://localhost:4040/api/tunnels")
	if err != nil {
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error fetching ngrok tunnels: %d", resp.StatusCode)
	}

	// Decode the response body as JSON
	var tunnelData TunnelsResponse
	err = json.NewDecoder(resp.Body).Decode(&tunnelData)
	if err != nil {
		return nil, err
	}

	// Return the tunnels
	return tunnelData.Tunnels, nil
}

type Tunnel struct {
	Name      string `json:"name"`
	PublicURL string `json:"public_url"`
}

type TunnelsResponse struct {
	Tunnels []Tunnel `json:"tunnels"`
}
