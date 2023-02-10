package client

import (
	"fmt"
	"github.com/kiosanim/jwt-client/model"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func TestClient(t *testing.T) {
	seeds, err := os.ReadFile("./seeds/token.json")
	if err != nil {
		t.Errorf(err.Error())
	}
	expected := string(seeds)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, strings.TrimSpace(expected))
	}))
	defer srv.Close()
	tokenClientInstance := GetInstance(model.Credentials{
		Username:     "juvenal",
		Password:     "123456",
		ClientID:     "fake_client_id",
		ClientSecret: "fake_client_secret",
		URIToken:     srv.URL + "/token",
		URILogout:    srv.URL + "/logout",
	})
	res, err := tokenClientInstance.GetAccessToken()
	if err != nil {
		t.Errorf(err.Error())
	}
	status := strings.Contains(expected, res)
	if status == false {
		t.Errorf("expected res to be %s got %s", expected, res)
	}
}
