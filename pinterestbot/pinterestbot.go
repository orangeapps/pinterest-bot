package pinterestbot

import (
	"github.com/dghubble/sling"
	"github.com/orangeapps/pinterest-bot/pinterestbot/resource"
)

type Client struct {
	sling  *sling.Sling
	Search *SearchService
	Pin    PinService
	User   UserService
}

func NewClient() *Client {
	base := sling.New().Base(resource.UrlBase)
	api := resource.NewApiService(base)
	return &Client{
		sling:  base,
		Search: newSearchService(base),
		Pin:    NewPinService(api),
		User:   NewUserService(api),
	}

}
