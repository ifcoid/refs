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

func TestSearchCrossrefWorks(t *testing.T) {
	result, err := SearchCrossrefWorks("machine learning", 2, 0)
	if err != nil {
		t.Fatalf("Gagal mencari data di Crossref: %v", err)
	}
	if result == nil || len(result.Message.Items) == 0 {
		t.Fatalf("Hasil pencarian kosong atau nil")
	}
	if len(result.Message.Items) > 2 {
		t.Errorf("Limit tidak berjalan dengan baik. Didapatkan %d item, diharapkan maksimal 2", len(result.Message.Items))
	}
}

func TestGetCrossrefJournal(t *testing.T) {
	// Contoh ISSN jurnal yang valid
	issn := "2644-125X"
	result, err := GetCrossrefJournal(issn)
	if err != nil {
		t.Fatalf("Gagal mengambil data jurnal: %v", err)
	}
	if result == nil || result.Title == "" {
		t.Fatalf("Data jurnal kosong")
	}
}

func TestGetCrossrefFunder(t *testing.T) {
	// 100000001 adalah ID untuk National Science Foundation
	id := "100000001"
	result, err := GetCrossrefFunder(id)
	if err != nil {
		t.Fatalf("Gagal mengambil data funder: %v", err)
	}
	if result == nil || result.Name == "" {
		t.Fatalf("Data funder kosong")
	}
}

func TestGetCrossrefMember(t *testing.T) {
	// 98 adalah ID untuk Hindawi
	id := "98"
	result, err := GetCrossrefMember(id)
	if err != nil {
		t.Fatalf("Gagal mengambil data member: %v", err)
	}
	if result == nil || result.PrimaryName == "" {
		t.Fatalf("Data member kosong")
	}
}

func TestGetCrossrefTypes(t *testing.T) {
	result, err := GetCrossrefTypes()
	if err != nil {
		t.Fatalf("Gagal mengambil daftar tipe: %v", err)
	}
	if len(result) == 0 {
		t.Fatalf("Daftar tipe kosong")
	}
}

func TestGetCrossrefLicenses(t *testing.T) {
	result, err := GetCrossrefLicenses()
	if err != nil {
		t.Fatalf("Gagal mengambil daftar lisensi: %v", err)
	}
	if len(result) == 0 {
		t.Fatalf("Daftar lisensi kosong")
	}
}
