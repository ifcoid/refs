package refs

type SearchResponse struct {
	SearchResults struct {
		Entry []struct {
			Title string `json:"dc:title"`
			DOI   string `json:"prism:doi"`
		} `json:"entry"`
		TotalResults string `json:"opensearch:totalResults"`
	} `json:"search-results"`
}

type FullTextResponse struct {
	FullTextRetrievalResponse struct {
		CoreData struct {
			URL              string `json:"prism:url"`
			Identifier       string `json:"dc:identifier"`
			EID              string `json:"eid"`
			DOI              string `json:"prism:doi"`
			PII              string `json:"pii"`
			Title            string `json:"dc:title"`
			PublicationName  string `json:"prism:publicationName"`
			AggregationType  string `json:"prism:aggregationType"`
			PubType          string `json:"pubType"`
			ISSN             string `json:"prism:issn"`
			Volume           string `json:"prism:volume"`
			StartingPage     string `json:"prism:startingPage"`
			PageRange        string `json:"prism:pageRange"`
			ArticleNumber    string `json:"articleNumber"`
			Format           string `json:"dc:format"`
			CoverDate        string `json:"prism:coverDate"`
			CoverDisplayDate string `json:"prism:coverDisplayDate"`
			Copyright        string `json:"prism:copyright"`
			Publisher        string `json:"prism:publisher"`
			Description      string `json:"dc:description"`
			Creators         []struct {
				Name string `json:"$"`
			} `json:"dc:creator"`
			Subjects []struct {
				Name string `json:"$"`
			} `json:"dcterms:subject"`
			OpenAccess         string `json:"openaccess"`
			OpenAccessArticle  bool   `json:"openaccessArticle"`
			OpenArchiveArticle bool   `json:"openArchiveArticle"`
			Links              []struct {
				Href string `json:"@href"`
				Rel  string `json:"@rel"`
			} `json:"link"`
		} `json:"coredata"`
		ScopusID     string      `json:"scopus-id"`
		ScopusEID    string      `json:"scopus-eid"`
		OriginalText interface{} `json:"originalText"`
	} `json:"full-text-retrieval-response"`
}

type ScopusAbstractResponse struct {
	AbstractsRetrievalResponse struct {
		AuthKeywords struct {
			AuthorKeyword []struct {
				Value string `json:"$"`
			} `json:"author-keyword"`
		} `json:"authkeywords"`
		IdxTerms struct {
			MainTerm []struct {
				Value string `json:"$"`
			} `json:"mainterm"`
		} `json:"idxterms"`
		SubjectAreas struct {
			SubjectArea []struct {
				Value  string `json:"$"`
				Abbrev string `json:"@abbrev"`
			} `json:"subject-area"`
		} `json:"subject-areas"`
	} `json:"abstracts-retrieval-response"`
}

// CrossrefResponse merepresentasikan balikan utama dari API Crossref
type CrossrefResponse struct {
	Status  string       `json:"status"`
	Message CrossrefWork `json:"message"`
}

// CrossrefWork merepresentasikan detail dari metadata suatu publikasi
type CrossrefWork struct {
	Publisher      string           `json:"publisher"`
	Title          []string         `json:"title"`
	URL            string           `json:"URL"`
	DOI            string           `json:"DOI"`
	Type           string           `json:"type"`
	Subject        []string         `json:"subject"`
	Author         []CrossrefAuthor `json:"author"`
	Issued         CrossrefDate     `json:"issued"`
	Created        CrossrefDate     `json:"created"`
	ContainerTitle []string         `json:"container-title"`
	Volume         string           `json:"volume"`
	Issue          string           `json:"issue"`
	Page           string           `json:"page"`
	Abstract       string           `json:"abstract"`
	Language       string           `json:"language"`
	ISSN           []string         `json:"ISSN"`
	ISBN           []string         `json:"ISBN"`
}

// CrossrefAuthor memuat detail nama penulis dan afiliasinya
type CrossrefAuthor struct {
	Given       string `json:"given"`
	Family      string `json:"family"`
	Sequence    string `json:"sequence"`
	Affiliation []struct {
		Name string `json:"name"`
	} `json:"affiliation"`
}

// CrossrefDate merepresentasikan tanggal keluaran / buatan di format Crossref
type CrossrefDate struct {
	DateParts [][]int `json:"date-parts"`
	DateTime  string  `json:"date-time"`
}

// CSLJSON merepresentasikan format metadata standar CSL-JSON
type CSLJSON struct {
	Type           string      `json:"type"`
	Title          string      `json:"title"`
	Author         []CSLAuthor `json:"author"`
	Issued         CSLDate     `json:"issued"`
	Publisher      string      `json:"publisher"`
	URL            string      `json:"URL"`
	DOI            string      `json:"DOI"`
	Volume         string      `json:"volume"`
	Issue          string      `json:"issue"`
	Page           string      `json:"page"`
	ContainerTitle string      `json:"container-title"`
	Abstract       string      `json:"abstract"`
}

// CSLAuthor merepresentasikan detail penulis pada CSL-JSON
type CSLAuthor struct {
	Family string `json:"family"`
	Given  string `json:"given"`
}

// CSLDate merepresentasikan format tanggal pada CSL-JSON
type CSLDate struct {
	DateParts [][]int `json:"date-parts"`
}

// CrossrefSearchResponse menampung hasil query
type CrossrefSearchResponse struct {
	Status  string `json:"status"`
	Message struct {
		TotalResults int            `json:"total-results"`
		ItemsPerPage int            `json:"items-per-page"`
		Items        []CrossrefWork `json:"items"`
	} `json:"message"`
}

// CrossrefJournalResponse menampung response jurnal
type CrossrefJournalResponse struct {
	Status  string         `json:"status"`
	Message CrossrefJournal `json:"message"`
}

// CrossrefJournal menampung profil jurnal
type CrossrefJournal struct {
	Title     string   `json:"title"`
	Publisher string   `json:"publisher"`
	ISSN      []string `json:"ISSN"`
}

// CrossrefFunderResponse menampung response funder
type CrossrefFunderResponse struct {
	Status  string        `json:"status"`
	Message CrossrefFunder `json:"message"`
}

// CrossrefFunder menampung profil penyandang dana
type CrossrefFunder struct {
	ID       string   `json:"id"`
	Location string   `json:"location"`
	Name     string   `json:"name"`
	AltNames []string `json:"alt-names"`
	URI      string   `json:"uri"`
}

// CrossrefMemberResponse menampung response member/publisher
type CrossrefMemberResponse struct {
	Status  string         `json:"status"`
	Message CrossrefMember `json:"message"`
}

// CrossrefMember menampung profil institusi/publisher
type CrossrefMember struct {
	ID            int      `json:"id"`
	PrimaryName   string   `json:"primary-name"`
	Names         []string `json:"names"`
	Prefixes      []string `json:"prefixes"`
	Location      string   `json:"location"`
}

// CrossrefTypeListResponse menampung daftar tipe
type CrossrefTypeListResponse struct {
	Status  string `json:"status"`
	Message struct {
		Items []CrossrefType `json:"items"`
	} `json:"message"`
}

// CrossrefType memuat label tipe publikasi
type CrossrefType struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

// CrossrefLicenseListResponse menampung daftar lisensi
type CrossrefLicenseListResponse struct {
	Status  string `json:"status"`
	Message struct {
		Items []CrossrefLicense `json:"items"`
	} `json:"message"`
}

// CrossrefLicense memuat informasi lisensi
type CrossrefLicense struct {
	URL       string `json:"URL"`
	WorkCount int    `json:"work-count"`
}
