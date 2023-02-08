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

type tokenClientSingleton struct {
	currentTokens *model.Tokens
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

var lock = &sync.Mutex{}
var singleInstance *tokenClientSingleton

func GetTokenClientInstance() *tokenClientSingleton {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		singleInstance = &tokenClientSingleton{}
	}
	return singleInstance
}

func (t *tokenClientSingleton) authenticate() error {
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", viper.GetString("USERNAME"))
	form.Add("password", viper.GetString("PASSWORD"))
	form.Add("client_id", viper.GetString("CLIENT_ID"))
	form.Add("client_secret", viper.GetString("CLIENT_SECRET"))
	uri := viper.GetString("SERVER_ADDRESS") + ":" + viper.GetString("PORT") + viper.GetString("TOKEN_ROUTE")
	//uri := viper.GetString("TOKEN_ROUTE")
	req, err := http.NewRequest("POST", uri, strings.NewReader(form.Encode()))
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
	t.currentTokens = &tokens
	return nil
}

func (t *tokenClientSingleton) refreshTokens() error {
	form := url.Values{}
	form.Add("grant_type", "refresh_token")
	form.Add("client_id", viper.GetString("CLIENT_ID"))
	form.Add("client_secret", viper.GetString("CLIENT_SECRET"))
	form.Add("refresh_token", viper.GetString(t.currentTokens.RefreshToken))
	req, err := http.NewRequest("POST", viper.GetString("BASE_IAM_ENDPOINT")+viper.GetString("TOKEN_ENDPOINT"), strings.NewReader(form.Encode()))
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
	t.currentTokens = &tokens
	return nil
}

func (t *tokenClientSingleton) Logout() error {
	form := url.Values{}
	form.Add("client_id", viper.GetString("CLIENT_ID"))
	form.Add("client_secret", viper.GetString("CLIENT_SECRET"))
	form.Add("refresh_token", viper.GetString(t.currentTokens.RefreshToken))
	req, err := http.NewRequest("POST", viper.GetString("BASE_IAM_ENDPOINT")+viper.GetString("LOGOUT_ENDPOINT"), strings.NewReader(form.Encode()))
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
	t.currentTokens = &tokens
	return nil
}

func (t *tokenClientSingleton) isTokenValid(token string) error {
	_, err := jwt.Parse(token, nil)
	if err != nil {
		return err
	}
	return nil
}

func (t *tokenClientSingleton) GetAccessToken() (string, error) {
	if t.currentTokens == nil {
		err := t.authenticate()
		if err != nil {
			return "", err
		}
		return t.currentTokens.AccessToken, nil
	}
	err := t.isTokenValid(t.currentTokens.AccessToken)
	if err == nil {
		return t.currentTokens.AccessToken, nil
	}
	err = t.isTokenValid(t.currentTokens.RefreshToken)
	if err == nil {
		return t.currentTokens.AccessToken, nil
	}
	err = t.refreshTokens()
	if err == nil {
		return t.currentTokens.AccessToken, nil
	}
	err = t.authenticate()
	if err != nil {
		return "", err
	}
	return t.currentTokens.AccessToken, nil
}

var httpGet = func(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(resp.Body)
}
