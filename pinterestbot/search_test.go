package pinterestbot

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"testing"
)

func TestSearch(t *testing.T) {

	pbot := NewClient()

	opts := new(SearchParameters)
	var q int
	limit := 10

	var bk []string

	for {

		opts.Query = "owl"
		opts.IsPrefetch = false
		opts.PageSize = 5
		opts.Bookmarks = bk

		pins, err := pbot.Pin.Search(*opts)
		if err != nil {
			t.Error(err)
		}

		bk = pins.Resource.Options.Bookmarks
		fmt.Println(bk)

		for _, v := range pins.ResourceResponse.Data.Results {
			user, _ := pbot.Pin.Info(v.ID)
			fmt.Println(user.Pinner.ID)
		}

		//file, _ := json.MarshalIndent(pins.ResourceResponse.Data.Results, "", " ")
		//
		//_ = ioutil.WriteFile("test.json", file, 0644)

		if len(pins.Resource.Options.Bookmarks) == 0 {
			t.Log("NO BOOKMARK", q)
		}

		q++
		if q == limit {
			break
		}
	}
}

func TestGetToken(t *testing.T) {
	//GetToken()

	url := "https://www.worldometers.info/coronavirus/"

	resp, _ := http.Get(url)

	body := html.NewTokenizer(resp.Body)
	for {
		doc := body.Next()
		switch {
		case doc == html.ErrorToken:
			// End of the document, we're done
			return
		case doc == html.StartTagToken:
			tag := body.Token()
			isRow := tag.Data == "tr"
			if isRow {

				fmt.Println(tag.Type)
			}
		}
	}

	resp.Body.Close()
}
