package refs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetCrossrefWork mengambil metadata suatu DOI langsung dari API Crossref.
// Mengembalikan data terstruktur dalam bentuk pointer ke CrossrefWork.
func GetCrossrefWork(doi string) (*CrossrefWork, error) {
	url := "https://api.crossref.org/works/" + doi

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("gagal membuat request: %v", err)
	}
	
	// Opsional tapi disarankan oleh Crossref (Polite Pool)
	req.Header.Set("User-Agent", "refs-package/0.1.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gagal menghubungi API Crossref: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API Crossref mengembalikan status error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("gagal membaca response body: %v", err)
	}

	var apiResponse CrossrefResponse
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		return nil, fmt.Errorf("gagal melakukan parsing JSON: %v", err)
	}

	return &apiResponse.Message, nil
}
