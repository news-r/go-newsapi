package newsapi

const baseurl = "https://newsapi.org/"

// GetHeadlines get top headlines
func GetHeadlines(client *HeadlinesClient) Response {
	preprocHeadlines(client)
	params := headlinesParams(*client)
	url := buildURL(baseurl, params)
	response := callAPI(url)
	return response
}
