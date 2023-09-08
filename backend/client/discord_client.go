package client

import (
	"encoding/json"
	"errors"
	"net/http"

	"backend/api/types"
	"backend/env"
)

type DiscordClient struct {
	httpClient *http.Client
}

type DiscordApiRoute string

const (
	DiscordUserRoute  DiscordApiRoute = "/users/@me"
	DiscordAuthRoute  DiscordApiRoute = "/oauth2/authorize"
	DiscordTokenRoute DiscordApiRoute = "/oauth2/token"
)

func (route DiscordApiRoute) URL() string {
	return env.DISCORD_API_URL.Value() + string(route)
}

func NewDiscordClient() *DiscordClient {
	return &DiscordClient{
		httpClient: NewHttpClient(),
	}
}

func (client *DiscordClient) FetchUser(token string) (*types.DiscordUser, error) {
	var res *types.DiscordUserResponse
	if err := client.fetch(token, DiscordUserRoute, &res); err != nil {
		return nil, err
	}

	return res.DiscordUser(), nil
}

// fetch fetches data from the DiscordClient API and stores the response body in the variable pointed to by dest.
func (client *DiscordClient) fetch(token string, route DiscordApiRoute, dest interface{}) error {
	req, err := http.NewRequest("GET", route.URL(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("request to discord api failed")
	}

	return json.NewDecoder(res.Body).Decode(dest)
}
