package newsapi

// Response A response object from the newsapi.
type Response struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// Article news articles as returned by the API
type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

// Source The source of the article.
type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HeadlinesClient creates a client to query top headlines
type HeadlinesClient struct {
	APIKey   string `url:"apiKey"`
	Query    string `url:"q"`
	Country  string `url:"country,omitempty"`
	Category string `url:"category,omitempty"`
	Sources  string `url:"sources,omitempty"`
	PageSize int    `url:"pageSize,omitempty"`
	Page     int    `url:"page,omitempty"`
}
