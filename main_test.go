package refs

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestScopusSearch(t *testing.T) {
	apiKey := os.Getenv("SCOPUS_API_KEY")
	if apiKey == "" {
		t.Fatal("Environment variable SCOPUS_API_KEY belum diset. Silakan set terlebih dahulu.")
	}

	// Gunakan keyword statis untuk pengujian
	query := "eegmamba ssm"
	fmt.Printf("Mencari artikel dengan kata kunci: '%s'\n", query)

	// 1. Proses Pencarian di Scopus
	searchResult, err := searchScopusArticles(query, apiKey)
	if err != nil {
		t.Fatalf("Error searching articles: %v", err)
	}

	if len(searchResult.SearchResults.Entry) == 0 {
		fmt.Println("Tidak ada artikel ditemukan.")
		return
	}

	entry := searchResult.SearchResults.Entry[0]
	if entry.DOI == "" {
		fmt.Println("Artikel ditemukan tetapi tidak memiliki DOI, tidak dapat mengambil Full Text.")
		return
	}

	fmt.Println("--------------------------------------------------")
	fmt.Printf("Ditemukan total %s hasil.\n", searchResult.SearchResults.TotalResults)
	fmt.Printf("Mencoba mengambil Full Text untuk Artikel Teratas:\n")
	fmt.Printf("Judul : %s\n", entry.Title)
	fmt.Printf("DOI   : %s\n", entry.DOI)
	fmt.Println("--------------------------------------------------")

	fullTextResult, err := retrieveFullText(entry.DOI, apiKey)
	if err != nil {
		fmt.Printf("\n[WARNING] Tidak dapat mengambil Full Text: %v\n", err)
		return
	}

	cd := fullTextResult.FullTextRetrievalResponse.CoreData

	// Cetak Metadata Tambahan
	fmt.Println("\n[INFO] Metadata Ekstra:")
	fmt.Printf("- Jurnal: %s\n", cd.PublicationName)
	fmt.Printf("- Tanggal Publikasi: %s\n", cd.CoverDate)
	fmt.Printf("- Penerbit: %s\n", cd.Publisher)

	// Buat BibTeX
	var authors []string
	for _, c := range cd.Creators {
		authors = append(authors, c.Name)
	}
	authorString := strings.Join(authors, " and ")
	year := cd.CoverDate
	if len(year) >= 4 {
		year = year[:4]
	}

	// Gunakan kata pertama nama penulis pertama dan tahun untuk ID BibTeX
	bibID := "article"
	if len(authors) > 0 {
		firstAuthorParts := strings.Split(authors[0], ",")
		bibID = strings.TrimSpace(firstAuthorParts[0]) + year
	}

	fmt.Println("\n[INFO] BibTeX:")
	fmt.Printf("@article{%s,\n", bibID)
	fmt.Printf("  title={%s},\n", cd.Title)
	if authorString != "" {
		fmt.Printf("  author={%s},\n", authorString)
	}
	if cd.PublicationName != "" {
		fmt.Printf("  journal={%s},\n", cd.PublicationName)
	}
	if cd.Volume != "" {
		fmt.Printf("  volume={%s},\n", cd.Volume)
	}
	if cd.PageRange != "" {
		fmt.Printf("  pages={%s},\n", cd.PageRange)
	}
	if year != "" {
		fmt.Printf("  year={%s},\n", year)
	}
	if cd.Publisher != "" {
		fmt.Printf("  publisher={%s},\n", cd.Publisher)
	}
	if cd.DOI != "" {
		fmt.Printf("  doi={%s}\n", cd.DOI)
	}
	fmt.Println("}")

	abstract := cd.Description
	if abstract != "" {
		fmt.Println("\n[INFO] Abstrak Artikel:\n", strings.TrimSpace(abstract))
	} else {
		fmt.Println("\n[INFO] Abstrak tidak tersedia untuk artikel ini.")
	}

	originalTextRaw := fullTextResult.FullTextRetrievalResponse.OriginalText
	var originalText string
	if str, ok := originalTextRaw.(string); ok {
		originalText = str
	} else if m, ok := originalTextRaw.(map[string]interface{}); ok {
		if val, ok2 := m["$"].(string); ok2 {
			originalText = val
		}
	}

	if originalText != "" {
		fmt.Println("\n[SUCCESS] Full text berhasil didapatkan!")

		limit := len(originalText)
		if limit > 600 {
			limit = 600
		}
		fmt.Printf("\nPotongan Teks (Tampilkan Maksimal 600 Karakter):\n\n%s...\n", originalText[:limit])
	} else {
		fmt.Println("\n[WARNING] Berhasil menghubungi API, tapi Full Text ('originalText') kosong.")
		fmt.Println("Alasan: Artikel ini mungkin tidak Open Access dan akun API Key Anda tidak memiliki langganan ke jurnal bersangkutan.")
	}
}
