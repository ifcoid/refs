package refs

import (
	"strings"
	"testing"
)

func TestDOIFetching(t *testing.T) {
	// DOI referensi yang stabil untuk testing
	doi := "10.1109/OJCOMS.2026.3666740"

	t.Run("CSL-JSON", func(t *testing.T) {
		csl, err := GetDOIMetadataCSLJSON(doi)
		if err != nil {
			t.Fatalf("Gagal mengambil CSL-JSON: %v", err)
		}
		if csl == nil {
			t.Fatalf("CSL-JSON nil")
		}
		if !strings.EqualFold(csl.DOI, "10.1109/OJCOMS.2026.3666740") {
			t.Errorf("DOI tidak cocok. Didapat: %s", csl.DOI)
		}
	})

	t.Run("BibTeX", func(t *testing.T) {
		bibtex, err := GetDOIBibTeX(doi)
		if err != nil {
			t.Fatalf("Gagal mengambil BibTeX: %v", err)
		}
		if !strings.Contains(bibtex, "@article") {
			t.Errorf("BibTeX tidak mengandung @article. Hasil: %s", bibtex)
		}
	})

	t.Run("RIS", func(t *testing.T) {
		ris, err := GetDOIRIS(doi)
		if err != nil {
			t.Fatalf("Gagal mengambil RIS: %v", err)
		}
		// RIS format typically starts with "TY  - "
		if !strings.Contains(ris, "TY  - ") {
			t.Errorf("RIS tidak mengandung TY tag. Hasil: %s", ris)
		}
	})

	t.Run("Formatted Citation (APA)", func(t *testing.T) {
		citation, err := GetDOICitationText(doi, "apa", "en-US")
		if err != nil {
			t.Fatalf("Gagal mengambil Formatted Citation: %v", err)
		}
		if len(citation) < 10 {
			t.Errorf("Citation string terlalu pendek. Hasil: %s", citation)
		}
	})
}
