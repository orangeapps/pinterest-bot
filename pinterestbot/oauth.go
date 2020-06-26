package pinterestbot

import (
	"context"
	"golang.org/x/oauth2"
	"log"
)

func GetToken() {

	config := &oauth2.Config{
		ClientID:     "5083412461135121305",
		ClientSecret: "66b53d84fd10101db866d45b0b20163ddd57c9f60cd305251ef6dc062f2a42a0",
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://api.pinterest.com/oauth/",
			TokenURL:  "https://api.pinterest.com/v1/oauth/token",
		},
		RedirectURL: "https://75518b1b.ngrok.io/callback",
		//Scopes:      []string{"read_public", "write_public", "read_relationships", "write_relationships"},
		Scopes: []string{"read_public,write_public,read_relationships,write_relationships"},
	}


	//code, err := config.AuthCodeURL("state-token", "")

	_, err := config.Exchange(context.Background(), "55587b8d06c6ceb3", grant)
	if err != nil {
		log.Fatal(err)
	}
}

var grant oauth2.AuthCodeOption = oauth2.SetAuthURLParam("grant_type", "authorization_code")