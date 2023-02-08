package main

import (
	"fmt"
	"github.com/kiosanim/jwt-client/client"
	"github.com/spf13/viper"
	"log"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	tokenClientInstance := client.GetTokenClientInstance()
	accessToken, err := tokenClientInstance.GetAccessToken()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(accessToken)
}
