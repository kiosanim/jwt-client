package model

type Credentials struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	URIToken     string `json:"uri_token"`
	URILogout    string `json:"uri_logout"`
}
