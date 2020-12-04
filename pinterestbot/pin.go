package pinterestbot

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"github.com/orangeapps/pinterest-bot/pinterestbot/pin"
	"github.com/orangeapps/pinterest-bot/pinterestbot/resource"
)

type PinService interface {
	Search(SearchParameters) (*pin.Response, error)
	SearchResource(SearchParameters) (*pin.SearchResponse, error)
	Info(pinID string) (*pin.InfoPin, error)
}

type pinService struct {
	http *resource.Api
}

func (s pinService) Info(pinID string) (*pin.InfoPin, error) {

	var info interface{}
	request := new(resource.GetParameters)
	response := new(pin.InfoPin)
	request.Scope = "pins"
	request.ID = pinID
	request.FieldSetKey = "detailed"

	_, _, err := s.http.Get(resource.UrlPinInfo, *request, &info)
	if err != nil {
		return nil, err
	}

	holder := info.(map[string]interface{})["resource_response"]
	infoHolder := holder.(map[string]interface{})["data"]
	mapstructure.Decode(infoHolder, &response)
	return response, nil
}

//Search
func (s pinService) Search(param SearchParameters) (*pin.Response, error) {

	request := resource.GetParameters(param)
	response := new(pin.Response)

	url := ""
	if len(param.Bookmarks) == 0 {
		url = resource.UrlSearch
	} else {
		url = resource.UrlSearchWithPagination
	}

	request.Scope = "pins"
	_, _, err := s.http.Get(url, request, &response)

	b, _ := json.Marshal(&response)

	// fmt.Println("[PINTEREST RESPONSE]: ", string(b))

	if err != nil {
		return response, nil
	}

	return response, nil
}

//Search Resource
func (s pinService) SearchResource(param SearchParameters) (*pin.SearchResponse, error) {

	request := resource.GetParameters(param)
	response := new(pin.SearchResponse)

	url := ""
	if len(param.Bookmarks) == 0 {
		url = resource.UrlSearch
	} else {
		url = resource.UrlSearchWithPagination
	}

	request.Scope = "pins"
	_, _, err := s.http.Get(url, request, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewPinService(http *resource.Api) PinService {
	return pinService{http: http}
}
