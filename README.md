# refs
Reference Lookup Golang Package

`refs` adalah package utilitas berbahasa Go (Golang) yang dirancang untuk memudahkan developer dalam mencari dan mengekstrak metadata dari berbagai pustaka penyedia referensi akademik.

## Instalasi
Gunakan perintah berikut untuk mengunduh package:
```bash
go get github.com/ifcoid/refs@latest
```

## Fitur & Penggunaan

### 1. Ekstraksi Artikel via Scopus API
Anda dapat melakukan pencarian berdasar kata kunci dan langsung menarik _Full-Text_ dari database Scopus Elsevier.

```go
import "github.com/ifcoid/refs"

// Melakukan pencarian menggunakan API Key Scopus
searchResult, err := refs.SearchScopusArticles("machine learning", "SCOPUS_API_KEY_ANDA")

// Mengambil detail metadata dan Full Text (bila Open Access)
fullText, err := refs.RetrieveFullText("10.1109/OJCOMS.2026.3666740", "SCOPUS_API_KEY_ANDA")
```

### 2. Ekstraksi Sitasi via doi.org (Content Negotiation)
Metode ini sangat fleksibel untuk mendapatkan sitasi dalam berbagai *format standar* tanpa parsing manual.

```go
import "github.com/ifcoid/refs"

// Mendapatkan metadata format JSON CSL
jsonCSL, err := refs.GetDOIMetadataCSLJSON("10.1109/OJCOMS.2026.3666740")

// Mendapatkan metadata format BibTeX
bibtex, err := refs.GetDOIBibTeX("10.1109/OJCOMS.2026.3666740")

// Mendapatkan Formatted Citation Text (misal: format APA, IEEE)
citation, err := refs.GetDOICitationText("10.1109/OJCOMS.2026.3666740", "apa", "en-US")
```

### 3. Ekstraksi Profil Lengkap via Crossref API
Kini tersedia beragam fungsi *endpoint* tambahan dari Crossref yang memudahkan Anda mengekstrak profil entitas akademik dan memetakan responnya (Unmarshal) secara otomatis ke **Struct Go**.

```go
import "github.com/ifcoid/refs"

// 1. Mencari Jurnal berdasarkan ISSN
journal, _ := refs.GetCrossrefJournal("2644-125X")

// 2. Mencari Profil Institusi / Funder
funder, _ := refs.GetCrossrefFunder("100000001")
member, _ := refs.GetCrossrefMember("98")

// 3. Melihat kamus tipe dan lisensi
types, _ := refs.GetCrossrefTypes()
licenses, _ := refs.GetCrossrefLicenses()

// 4. Pencarian Artikel Bebas
// (Parameter: Query, Limit/Rows, Offset)
results, _ := refs.SearchCrossrefWorks("machine learning", 10, 0)
```

> [!WARNING]
> **Tata Cara Penggunaan Limitasi Crossref API**: 
> - Secara bawaan (*default*), API Crossref membatasi balikan array (seperti pada fungsi *SearchCrossrefWorks*) ke **20 item** apabila Anda memasukkan angka limit `0`.
> - Batas maksimum (*hard limit*) penarikan data per *request* adalah **1000 item**.
> - Jika Anda membutuhkan lebih dari 1000 item untuk keperluan analisis data besar (*big data*), gunakan parameter **offset** (pagination) untuk menarik kelanjutan datanya (misal: offset=1000 untuk halaman kedua).

## Cara Test

Untuk menjalankan pengujian pada package ini, Anda memerlukan API Key dari Scopus (untuk fungsi Scopus, sementara Crossref dan DOI tidak memerlukan kunci). 

1. **Set Environment Variable `SCOPUS_API_KEY`**
   
   **Linux / macOS / Git Bash:**
   ```bash
   export SCOPUS_API_KEY="api_key_anda"
   ```

   **Windows (PowerShell):**
   ```powershell
   $env:SCOPUS_API_KEY="api_key_anda"
   ```

2. **Jalankan Test**
   ```bash
   go test -v
   ```

## Release Package

Untuk merilis versi terbaru dari package ini, ikuti langkah-langkah Git tag berikut:

1. **Commit dan Push ke Branch Utama**
   ```bash
   git add .
   git commit -m "chore: rilis versi terbaru"
   git push origin main
   ```

2. **Buat Tag Baru (Otomatis)**
   **Linux / macOS / Git Bash:**
   ```bash
   LATEST_TAG=$(git tag --sort=v:refname | tail -1)
   NEW_TAG=$(echo ${LATEST_TAG:-v0.0.0} | awk -F. -v OFS=. '{$NF++;print}')
   git tag $NEW_TAG
   ```
   **Windows (PowerShell):**
   ```powershell
   $latest = git tag --sort=v:refname | Select-Object -Last 1
   if (-not $latest) { $latest = "v0.0.0" }
   $parts = $latest -replace '^v','' -split '\.'
   $parts[2] = [int]$parts[2] + 1
   $newTag = "v$($parts[0]).$($parts[1]).$($parts[2])"
   git tag $newTag
   ```

3. **Push Tag ke Server**
   ```bash
   git push origin --tags
   ```