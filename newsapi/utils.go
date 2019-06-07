package newsapi

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
)

func preprocHeadlines(client *HeadlinesClient) HeadlinesClient {
	if client.Page == 0 {
		client.Page = 1
	}

	if client.PageSize == 0 {
		client.PageSize = 1
	}

	return *client
}

func headlinesParams(client HeadlinesClient) string {
	values, err := query.Values(client)
	if err != nil {
		log.Fatal(err)
	}
	return values.Encode()
}

func buildURL(baseurl, params string) string {
	return baseurl + "v2/top-headlines?" + params
}

func callAPI(url string) Response {
	data := new(Response)

	// call API
	cl := &http.Client{Timeout: 10 * time.Second}
	resp, err := cl.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close() // close response

	// Unmarshall
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		log.Panic(err)
	}

	return (*data)
}
