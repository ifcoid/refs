package refs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// fetchDOIContent melakukan HTTP GET ke doi.org dan memastikan header Accept ikut terbawa saat redirect.
func fetchDOIContent(doi, acceptHeader string) (string, error) {
	// doi.org URL format
	url := "https://doi.org/" + doi

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("gagal membuat request: %v", err)
	}
	req.Header.Set("Accept", acceptHeader)

	// Buat custom HTTP client untuk handle redirect dan pastikan header Accept tetap ikut
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= 10 {
				return fmt.Errorf("terlalu banyak redirect")
			}
			// Copy header Accept ke request redirect
			req.Header.Set("Accept", acceptHeader)
			return nil
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("gagal melakukan request ke doi.org: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("API mengembalikan status error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("gagal membaca response body: %v", err)
	}

	return string(body), nil
}

// GetDOIMetadataCSLJSON mengembalikan metadata DOI dalam format JSON (CSL-JSON) yang sudah di-unmarshal ke struct CSLJSON.
func GetDOIMetadataCSLJSON(doi string) (*CSLJSON, error) {
	strResp, err := fetchDOIContent(doi, "application/vnd.citationstyles.csl+json")
	if err != nil {
		return nil, err
	}

	var csl CSLJSON
	if err := json.Unmarshal([]byte(strResp), &csl); err != nil {
		return nil, fmt.Errorf("gagal unmarshal CSL-JSON: %v", err)
	}
	return &csl, nil
}

// GetDOIBibTeX mengembalikan metadata DOI dalam format BibTeX.
func GetDOIBibTeX(doi string) (string, error) {
	return fetchDOIContent(doi, "application/x-bibtex")
}

// GetDOIRIS mengembalikan metadata DOI dalam format RIS (Research Information Systems).
func GetDOIRIS(doi string) (string, error) {
	return fetchDOIContent(doi, "application/x-research-info-systems")
}

// GetDOICitationText mengembalikan teks sitasi yang sudah diformat.
// style: contoh "apa", "ieee", "harvard3", "vancouver". Jika kosong, default akan digunakan.
// locale: contoh "en-US", "id-ID". Jika kosong, default akan digunakan.
func GetDOICitationText(doi, style, locale string) (string, error) {
	accept := "text/x-bibliography"
	var options []string
	if style != "" {
		options = append(options, "style="+style)
	}
	if locale != "" {
		options = append(options, "locale="+locale)
	}

	if len(options) > 0 {
		accept += "; " + strings.Join(options, "; ")
	}

	return fetchDOIContent(doi, accept)
}
