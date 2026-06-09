# refs
Reference Lookup Golang Package

## Fitur & Penggunaan

Package ini menyediakan fungsi untuk mengambil referensi dan sitasi dari berbagai sumber penyedia pustaka akademik secara langsung.

### 1. Ekstraksi via doi.org (Content Negotiation)
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

### 2. Ekstraksi Terstruktur via Crossref API
Metode ini berguna jika Anda ingin *programmability* lebih mendalam di Go. Data ditarik dari API resmi Crossref dan langsung di-unmarshal ke dalam *Struct Golang* (`*CrossrefWork`).

```go
import (
    "fmt"
    "github.com/ifcoid/refs"
)

crossrefData, err := refs.GetCrossrefWork("10.1109/OJCOMS.2026.3666740")
if err == nil && crossrefData != nil {
    fmt.Println("Judul Artikel:", crossrefData.Title[0])
    fmt.Println("Penerbit:", crossrefData.Publisher)
    fmt.Println("Penulis Utama:", crossrefData.Author[0].Given, crossrefData.Author[0].Family)
}
```

## Cara Test

Untuk menjalankan pengujian pada package ini, Anda memerlukan API Key dari Scopus. Ikuti langkah-langkah berikut:

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
   Pastikan semua kode terbaru sudah di-commit dan di-push ke repository.
   ```bash
   git add .
   git commit -m "chore: rilis versi terbaru"
   git push origin main
   ```

2. **Buat Tag Baru (Otomatis)**
   Anda bisa membuat tag baru secara otomatis (hanya menaikkan angka terakhir/patch, misal dari `v0.1.0` ke `v0.1.1`) tanpa perlu mengingat versi terakhir:

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
   Kirim tag yang baru dibuat ke GitHub:
   ```bash
   git push origin --tags
   ```

4. **Gunakan Versi Terbaru**
   Setelah tag berhasil di-push, project lain dapat mengunduh dan menggunakan versi tersebut dengan perintah:
   ```bash
   go get github.com/ifcoid/refs@v0.1.1
   ```