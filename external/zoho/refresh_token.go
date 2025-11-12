package zoho

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func RefreshAccessToken(refreshToken, clientID, clientSecret string) (*TokenResponse, error) {
	tokenURL := "https://accounts.zoho.com/oauth/v2/token"

	fmt.Printf("token is: %s, id is: %s, secret is: %s", refreshToken, clientID, clientSecret)
	params := url.Values{}
	params.Add("refresh_token", refreshToken)
	params.Add("client_id", clientID)
	params.Add("client_secret", clientSecret)
	params.Add("redirect_uri", ZOHO_REDIRECT_URL)
	params.Add("grant_type", "refresh_token")

	resp, err := http.Post(tokenURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return nil, err
	}

	if tokenResp.Error != "" {
		return nil, fmt.Errorf("error from Zoho: %s", tokenResp.Error)
	}

	return &tokenResp, nil
}
