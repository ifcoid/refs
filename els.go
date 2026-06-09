package refs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func searchScopusArticles(query string, apiKey string) (SearchResponse, error) {
	searchURL := fmt.Sprintf("https://api.elsevier.com/content/search/scopus?query=all(%s)&count=1", url.QueryEscape(query))

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("X-ELS-APIKey", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("error request Search API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return SearchResponse{}, fmt.Errorf("search API Error: %s\nResponse: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchResponse{}, fmt.Errorf("error reading response body: %w", err)
	}
	var searchResult SearchResponse
	if err := json.Unmarshal(body, &searchResult); err != nil {
		return SearchResponse{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return searchResult, nil
}

func retrieveFullText(doi string, apiKey string) (FullTextResponse, error) {
	retrieveURL := fmt.Sprintf("https://api.elsevier.com/content/article/doi/%s", doi)

	req, err := http.NewRequest("GET", retrieveURL, nil)
	if err != nil {
		return FullTextResponse{}, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("X-ELS-APIKey", apiKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return FullTextResponse{}, fmt.Errorf("error request Full Text API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return FullTextResponse{}, fmt.Errorf("article Retrieval API Error (Status: %s).\nResponse: %s", resp.Status, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return FullTextResponse{}, fmt.Errorf("error reading response body: %w", err)
	}
	var fullTextResult FullTextResponse
	if err := json.Unmarshal(body, &fullTextResult); err != nil {
		return FullTextResponse{}, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return fullTextResult, nil
}
