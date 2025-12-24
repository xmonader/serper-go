package serper

// Common country codes (gl)
const (
	GLUnitedStates   = "us"
	GLUnitedKingdom  = "uk"
	GLCanada         = "ca"
	GLGermany        = "de"
	GLFrance         = "fr"
	GLJapan          = "jp"
	GLAustralia      = "au"
	GLBrazil         = "br"
	GLIndia          = "in"
	GLChina          = "cn"
)

// Common language codes (hl)
const (
	HLEnglish    = "en"
	HLSpanish    = "es"
	HLFrench     = "fr"
	HLGerman     = "de"
	HLItalian    = "it"
	HLJapanese   = "ja"
	HLKorean     = "ko"
	HLChinese    = "zh-cn"
	HLPortuguese = "pt"
	HLRussian    = "ru"
)

// Request represents the parameters for any Serper search.
type Request struct {
	Q           string `json:"q"`
	Gl          string `json:"gl,omitempty"`
	Hl          string `json:"hl,omitempty"`
	Num         int    `json:"num,omitempty"`
	Autocorrect bool   `json:"autocorrect,omitempty"`
	Page        int    `json:"page,omitempty"`
	Type        string `json:"type,omitempty"`
	Location    string `json:"location,omitempty"`
	Tbs         string `json:"tbs,omitempty"`
	Safe        string `json:"safe,omitempty"`
}

// Parameters mirrors the search parameters returned in responses.
type Parameters struct {
	Q           string `json:"q"`
	Gl          string `json:"gl,omitempty"`
	Hl          string `json:"hl,omitempty"`
	Num         int    `json:"num,omitempty"`
	Autocorrect bool   `json:"autocorrect,omitempty"`
	Page        int    `json:"page,omitempty"`
	Type        string `json:"type,omitempty"`
	Location    string `json:"location,omitempty"`
}

// BaseResponse contains fields common to all search responses.
type BaseResponse struct {
	SearchParameters Parameters `json:"searchParameters"`
	Credits          int        `json:"credits"`
}

type SearchResponse struct {
	BaseResponse
	Organic         []OrganicResult  `json:"organic"`
	KnowledgeGraph  *KnowledgeGraph  `json:"knowledgeGraph,omitempty"`
	PeopleAlsoAsk   []PeopleAlsoAsk  `json:"peopleAlsoAsk,omitempty"`
	RelatedSearches []RelatedSearch  `json:"relatedSearches,omitempty"`
}

type OrganicResult struct {
	Title      string            `json:"title"`
	Link       string            `json:"link"`
	Snippet    string            `json:"snippet"`
	Position   int               `json:"position"`
	Date       string            `json:"date,omitempty"`
	Sitelinks  []Sitelink        `json:"sitelinks,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
}

type Sitelink struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type RelatedSearch struct {
	Query string `json:"query"`
}

type KnowledgeGraph struct {
	Title             string                 `json:"title"`
	Type              string                 `json:"type"`
	Website           string                 `json:"website,omitempty"`
	ImageUrl          string                 `json:"imageUrl,omitempty"`
	Description       string                 `json:"description,omitempty"`
	DescriptionSource string                 `json:"descriptionSource,omitempty"`
	DescriptionLink   string                 `json:"descriptionLink,omitempty"`
	Attributes        map[string]interface{} `json:"attributes,omitempty"`
}

type PeopleAlsoAsk struct {
	Question string `json:"question"`
	Snippet  string `json:"snippet"`
	Title    string `json:"title"`
	Link     string `json:"link"`
}

type ImageResponse struct {
	BaseResponse
	Images []ImageResult `json:"images"`
}

type ImageResult struct {
	Title           string `json:"title"`
	ImageUrl        string `json:"imageUrl"`
	ImageWidth      int    `json:"imageWidth,omitempty"`
	ImageHeight     int    `json:"imageHeight,omitempty"`
	ThumbnailUrl    string `json:"thumbnailUrl,omitempty"`
	ThumbnailWidth  int    `json:"thumbnailWidth,omitempty"`
	ThumbnailHeight int    `json:"thumbnailHeight,omitempty"`
	Source          string `json:"source"`
	Domain          string `json:"domain"`
	Link            string `json:"link"`
	GoogleUrl       string `json:"googleUrl"`
	Position        int    `json:"position"`
}

type NewsResponse struct {
	BaseResponse
	News []NewsResult `json:"news"`
}

type NewsResult struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Snippet  string `json:"snippet"`
	Date     string `json:"date"`
	Source   string `json:"source"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Position int    `json:"position"`
}

type VideoResponse struct {
	BaseResponse
	Videos []VideoResult `json:"videos"`
}

type VideoResult struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Snippet  string `json:"snippet"`
	ImageUrl string `json:"imageUrl,omitempty"`
	Duration string `json:"duration,omitempty"`
	Source   string `json:"source"`
	Channel  string `json:"channel,omitempty"`
	Date     string `json:"date,omitempty"`
	Position int    `json:"position"`
}

type PlacesResponse struct {
	BaseResponse
	Places []PlaceResult `json:"places"`
}

type PlaceResult struct {
	Position     int     `json:"position"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Rating       float64 `json:"rating,omitempty"`
	RatingCount  int     `json:"ratingCount,omitempty"`
	Category     string  `json:"category,omitempty"`
	PhoneNumber  string  `json:"phoneNumber,omitempty"`
	Website      string  `json:"website,omitempty"`
	Cid          string  `json:"cid,omitempty"`
	ThumbnailUrl string  `json:"thumbnailUrl,omitempty"`
}
