package client

import (
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"github.com/kiosanim/jwt-client/model"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var lock = sync.Mutex{}
var tokenClientInstance *tokenClient

type tokenClient struct {
	CurrentTokens *model.Tokens
	credentials   *model.Credentials
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return
	}
}

func GetInstance(credentials model.Credentials) *tokenClient {
	if tokenClientInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		tokenClientInstance = &tokenClient{credentials: &credentials}
	}
	return tokenClientInstance
}

func (t *tokenClient) authenticate(uri_token string) error {
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", t.credentials.Username)
	form.Add("password", t.credentials.Password)
	form.Add("client_id", t.credentials.ClientID)
	form.Add("client_secret", t.credentials.ClientSecret)
	//uri := viper.GetString("TOKEN_ROUTE")
	req, err := http.NewRequest("POST", uri_token, strings.NewReader(form.Encode()))
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return err
	}
	var tokens model.Tokens
	err = json.Unmarshal([]byte(bodyBytes), &tokens)
	if err != nil {
		log.Println(err)
		return err
	}
	t.CurrentTokens = &tokens
	return nil
}

func (t *tokenClient) refreshTokens(uri_token string) error {
	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", t.credentials.ClientID)
	form.Add("client_secret", t.credentials.ClientSecret)
	form.Add("refresh_token", t.CurrentTokens.RefreshToken)
	req, err := http.NewRequest("POST", uri_token, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var tokens model.Tokens
	err = json.Unmarshal([]byte(bodyBytes), &tokens)
	if err != nil {
		log.Fatal(err)
		return err
	}
	t.CurrentTokens = &tokens
	return nil
}

func (t *tokenClient) Logout(uri_logout string) error {
	form := url.Values{}
	form.Add("client_id", t.credentials.ClientID)
	form.Add("client_secret", t.credentials.ClientSecret)
	form.Add("refresh_token", t.CurrentTokens.RefreshToken)
	req, err := http.NewRequest("POST", uri_logout, strings.NewReader(form.Encode()))
	if err != nil {
		log.Fatal(err)
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer res.Body.Close()
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var tokens model.Tokens
	err = json.Unmarshal([]byte(bodyBytes), &tokens)
	if err != nil {
		log.Fatal(err)
		return err
	}
	t.CurrentTokens = &tokens
	return nil
}

func (t *tokenClient) isTokenValid(token string) error {
	_, err := jwt.Parse(token, nil)
	if err != nil {
		return err
	}
	return nil
}

func (t tokenClient) GetAccessToken() (string, error) {
	if t.CurrentTokens == nil {
		err := t.authenticate(t.credentials.URIToken)
		if err != nil {
			return "", err
		}
		return t.CurrentTokens.AccessToken, nil
	}
	err := t.isTokenValid(t.CurrentTokens.AccessToken)
	if err == nil {
		return t.CurrentTokens.AccessToken, nil
	}
	err = t.isTokenValid(t.CurrentTokens.RefreshToken)
	if err == nil {
		return t.CurrentTokens.AccessToken, nil
	}
	err = t.refreshTokens(t.credentials.URIToken)
	if err == nil {
		return t.CurrentTokens.AccessToken, nil
	}
	err = t.authenticate(t.credentials.URIToken)
	if err != nil {
		return "", err
	}
	return t.CurrentTokens.AccessToken, nil
}
