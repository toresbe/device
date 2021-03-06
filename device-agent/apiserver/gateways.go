package apiserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Gateway struct {
	PublicKey string   `json:"publicKey"`
	Endpoint  string   `json:"endpoint"`
	IP        string   `json:"ip"`
	Routes    []string `json:"routes"`
}

type UnauthorizedError struct{}

func (u *UnauthorizedError) Error() string {
	return "unauthorized"
}

func GetGateways(sessionKey, apiServerURL string, ctx context.Context) ([]Gateway, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	deviceConfigAPI := fmt.Sprintf("%s/deviceconfig", apiServerURL)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, deviceConfigAPI, nil)
	if err != nil {
		return nil, fmt.Errorf("creating get request: %w", err)
	}
	req.Header.Add("x-naisdevice-session-key", sessionKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("getting device config: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return nil, &UnauthorizedError{}
	}

	var gateways []Gateway
	if err := json.NewDecoder(resp.Body).Decode(&gateways); err != nil {
		return nil, fmt.Errorf("unmarshalling response body into gateways: %w", err)
	}

	return gateways, nil
}
