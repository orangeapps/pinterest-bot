package pinterestbot

import (
	"encoding/json"
	"github.com/dghubble/sling"
	"github.com/orangeapps/pinterest-bot/pinterestbot/resource"
	"net/http"
	"time"
)

type RequestOptions struct {
	Options *SearchParameters `json:"options" url:"options"`
	Context string            `json:"context" url:"context,omitempty"`
}

type RequestParams struct {
	SourceUrl string `json:"source_url" url:"source_url"`
	Data      string `json:"data" url:"data"`
	TS        int64  `json:"_" url:"_"`
}
type SearchParameters struct {
	Bookmarks   []string `json:"bookmarks,omitempty" url:"bookmarks,omitempty"`
	PageSize    int64    `json:"page_size,omitempty" url:"page_size,omitempty"`
	IsPrefetch  bool     `json:"isPrefetch,omitempty" url:"isPrefetch,omitempty"`
	Query       string   `json:"query,omitempty" url:"query,omitempty"`
	Scope       string   `json:"scope,omitempty" url:"scope,omitempty"`
	Count       string   `json:"count,omitempty" url:"count,omitempty"`
	FieldSetKey string   `json:"field_set_key,omitempty" url:"field_set_key,omitempty"`
	UserID      string   `json:"user_id,omitempty" url:"user_id,omitempty"`
	ID          string   `json:"id,omitempty" url:"id,omitempty"`
}

type SearchService struct {
	http *sling.Sling
}

func newSearchService(sling *sling.Sling) *SearchService {
	return &SearchService{
		http: sling,
	}
}

func (s *SearchService) Do(params *SearchParameters) (*resource.Response, *http.Response, error) {

	request := new(RequestParams)
	options := new(RequestOptions)
	response := new(resource.Response)

	request.SourceUrl = resource.BuildSourceUrl(params.Query, params.Scope)
	request.TS = time.Now().UnixNano()

	options.Options = params

	parsedOptions, err := json.Marshal(options)
	if err != nil {
		return nil, nil, err
	}

	request.Data = string(parsedOptions)
	resp, errr := s.http.New().Get(resource.UrlSearch).QueryStruct(request).Receive(response, err)
	if errr != nil {
		return nil, nil, errr
	}

	return response, resp, nil
}
