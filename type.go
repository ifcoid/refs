package refs

type SearchResponse struct {
	SearchResults struct {
		Entry []struct {
			Title string `json:"dc:title"`
			DOI   string `json:"prism:doi"`
		} `json:"entry"`
		TotalResults string `json:"opensearch:totalResults"`
	} `json:"search-results"`
}

type FullTextResponse struct {
	FullTextRetrievalResponse struct {
		CoreData struct {
			URL              string `json:"prism:url"`
			Identifier       string `json:"dc:identifier"`
			EID              string `json:"eid"`
			DOI              string `json:"prism:doi"`
			PII              string `json:"pii"`
			Title            string `json:"dc:title"`
			PublicationName  string `json:"prism:publicationName"`
			AggregationType  string `json:"prism:aggregationType"`
			PubType          string `json:"pubType"`
			ISSN             string `json:"prism:issn"`
			Volume           string `json:"prism:volume"`
			StartingPage     string `json:"prism:startingPage"`
			PageRange        string `json:"prism:pageRange"`
			ArticleNumber    string `json:"articleNumber"`
			Format           string `json:"dc:format"`
			CoverDate        string `json:"prism:coverDate"`
			CoverDisplayDate string `json:"prism:coverDisplayDate"`
			Copyright        string `json:"prism:copyright"`
			Publisher        string `json:"prism:publisher"`
			Description      string `json:"dc:description"`
			Creators         []struct {
				Name string `json:"$"`
			} `json:"dc:creator"`
			Subjects []struct {
				Name string `json:"$"`
			} `json:"dcterms:subject"`
			OpenAccess         string `json:"openaccess"`
			OpenAccessArticle  bool   `json:"openaccessArticle"`
			OpenArchiveArticle bool   `json:"openArchiveArticle"`
			Links              []struct {
				Href string `json:"@href"`
				Rel  string `json:"@rel"`
			} `json:"link"`
		} `json:"coredata"`
		ScopusID     string      `json:"scopus-id"`
		ScopusEID    string      `json:"scopus-eid"`
		OriginalText interface{} `json:"originalText"`
	} `json:"full-text-retrieval-response"`
}
