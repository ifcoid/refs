package refs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// fetchCrossrefAPI melakukan HTTP GET ke API Crossref dan melakukan unmarshal ke target struct.
func fetchCrossrefAPI(apiURL string, target interface{}) error {
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return fmt.Errorf("gagal membuat request: %v", err)
	}

	// Opsional tapi disarankan oleh Crossref (Polite Pool)
	req.Header.Set("User-Agent", "refs-package/0.2.0")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("gagal menghubungi API Crossref: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("API Crossref mengembalikan status error: %d %s", resp.StatusCode, resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("gagal membaca response body: %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("gagal melakukan parsing JSON: %v", err)
	}

	return nil
}

// GetCrossrefWork mengambil metadata suatu DOI langsung dari API Crossref.
// Mengembalikan data terstruktur dalam bentuk pointer ke CrossrefWork.
func GetCrossrefWork(doi string) (*CrossrefWork, error) {
	urlStr := "https://api.crossref.org/works/" + doi
	var apiResponse CrossrefResponse
	if err := fetchCrossrefAPI(urlStr, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse.Message, nil
}

// SearchCrossrefWorks mencari metadata publikasi berdasarkan kata kunci.
// Jika rows <= 0, default dari Crossref (20 item) akan digunakan. Batas maksimal rows adalah 1000.
// Offset digunakan untuk pagination.
func SearchCrossrefWorks(query string, rows int, offset int) (*CrossrefSearchResponse, error) {
	u, _ := url.Parse("https://api.crossref.org/works")
	q := u.Query()
	if query != "" {
		q.Set("query", query)
	}
	if rows > 0 {
		q.Set("rows", fmt.Sprintf("%d", rows))
	}
	if offset > 0 {
		q.Set("offset", fmt.Sprintf("%d", offset))
	}
	u.RawQuery = q.Encode()

	var apiResponse CrossrefSearchResponse
	if err := fetchCrossrefAPI(u.String(), &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}

// GetCrossrefJournal mengambil profil jurnal berdasarkan ISSN.
func GetCrossrefJournal(issn string) (*CrossrefJournal, error) {
	urlStr := "https://api.crossref.org/journals/" + url.PathEscape(issn)
	var apiResponse CrossrefJournalResponse
	if err := fetchCrossrefAPI(urlStr, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse.Message, nil
}

// GetCrossrefFunder mengambil profil penyandang dana (funder) berdasarkan ID.
func GetCrossrefFunder(id string) (*CrossrefFunder, error) {
	urlStr := "https://api.crossref.org/funders/" + url.PathEscape(id)
	var apiResponse CrossrefFunderResponse
	if err := fetchCrossrefAPI(urlStr, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse.Message, nil
}

// GetCrossrefMember mengambil profil institusi/publisher berdasarkan ID member Crossref.
func GetCrossrefMember(id string) (*CrossrefMember, error) {
	urlStr := "https://api.crossref.org/members/" + url.PathEscape(id)
	var apiResponse CrossrefMemberResponse
	if err := fetchCrossrefAPI(urlStr, &apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse.Message, nil
}

// GetCrossrefTypes mengambil seluruh kamus tipe artikel yang didukung Crossref.
func GetCrossrefTypes() ([]CrossrefType, error) {
	urlStr := "https://api.crossref.org/types"
	var apiResponse CrossrefTypeListResponse
	if err := fetchCrossrefAPI(urlStr, &apiResponse); err != nil {
		return nil, err
	}
	return apiResponse.Message.Items, nil
}

// GetCrossrefLicenses mengambil tipe-tipe lisensi (seperti lisensi Creative Commons).
func GetCrossrefLicenses() ([]CrossrefLicense, error) {
	urlStr := "https://api.crossref.org/licenses"
	var apiResponse CrossrefLicenseListResponse
	if err := fetchCrossrefAPI(urlStr, &apiResponse); err != nil {
		return nil, err
	}
	return apiResponse.Message.Items, nil
}
