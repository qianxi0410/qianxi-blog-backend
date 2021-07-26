package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

type OAuth2Service struct{}

func (o OAuth2Service) GetUserInfo(code string) (string, error) {
	viper.SetConfigFile("./config/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	resp, err := http.Post(fmt.Sprintf("%s?client_id=%s&client_secret=%s&code=%s",
		viper.GetString("github.token_url"), viper.GetString("github.client_id"),
		viper.GetString("github.client_secret"), code), "application/json;charset=utf-8", nil)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return "", nil
	}

	accessToken := strings.Split(strings.Split(string(body), "&")[0], "=")[1]

	req, err := http.NewRequest(http.MethodGet, viper.GetString("github.api_url"), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("token %s", accessToken))

	githubResp, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	defer githubResp.Body.Close()

	info, err := io.ReadAll(githubResp.Body)

	if err != nil {
		return "", nil
	}
	return string(info), nil
}
