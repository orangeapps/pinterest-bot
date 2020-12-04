package resource

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dghubble/sling"
)

type Api struct {
	http *sling.Sling
}

func NewApiService(sling *sling.Sling) *Api {
	return &Api{
		http: sling,
	}
}

type GetParameters struct {
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

type requestParams struct {
	SourceUrl string `json:"source_url" url:"source_url"`
	Data      string `json:"data" url:"data"`
	TS        int64  `json:"_" url:"_"`
}

type requestOptions struct {
	Options GetParameters `json:"options" url:"options"`
	Context string        `json:"context" url:"context,omitempty"`
}

func (a *Api) Get(resource string, request GetParameters, response interface{}) (interface{}, *http.Response, error) {
	requestParam := new(requestParams)
	options := new(requestOptions)
	requestParam.SourceUrl = "" //BuildSourceUrl(request.Query, request.Scope)
	requestParam.TS = time.Now().UnixNano()

	if request.PageSize == 0 {
		request.PageSize = 250
	}

	options.Options = request

	parsedOptions, err := json.Marshal(options)
	if err != nil {
		return nil, nil, err
	}

	requestParam.Data = string(parsedOptions)

	httpResponse, err := a.http.New().Get(resource).QueryStruct(requestParam).Receive(response, err)
	if err != nil {
		fmt.Println("[PINTEREST ERROR]: ", err)
		fmt.Println("[PINTEREST HTTP BODY RESPONSE]: ", &httpResponse.Body)

		body := strings.NewReader("raw body")
		a, err := a.http.New().Get(resource).Body(body).Request()
		fmt.Println("[TEST RESOURCE]: ", a)
		return nil, nil, err
	}
	fmt.Println("[PINTEREST HTTP BODY RESPONSE]: ", &httpResponse)
	fmt.Println("[PINTEREST HTTP RESPONSE]: ", response)
	return response, httpResponse, nil
}
