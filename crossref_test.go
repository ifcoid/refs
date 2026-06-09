package refs

import (
	"strings"
	"testing"
)

func TestGetCrossrefWork(t *testing.T) {
	doi := "10.1109/OJCOMS.2026.3666740"
	
	result, err := GetCrossrefWork(doi)
	if err != nil {
		t.Fatalf("Gagal mengambil data dari Crossref: %v", err)
	}

	if result == nil {
		t.Fatalf("Hasil pointer nil")
	}

	if !strings.EqualFold(result.DOI, doi) {
		t.Errorf("DOI tidak cocok. Diharapkan: %s, Didapat: %s", doi, result.DOI)
	}

	if len(result.Title) == 0 {
		t.Errorf("Title kosong")
	}
}
