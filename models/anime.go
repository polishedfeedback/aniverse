package models

type Anime struct {
	MalID         int      `json:"mal_id"`
	URL           string   `json:"url"`
	Title         string   `json:"title"`
	TitleEnglish  string   `json:"title_english"`
	TitleJapanese string   `json:"title_japanese"`
	Synopsis      string   `json:"synopsis"`
	Episodes      int      `json:"episodes"`
	Year          int      `json:"year"`
	Score         float32  `json:"score"`
	Rank          int      `json:"rank"`
	Status        string   `json:"status"`
	Type          string   `json:"type"`
	Genres        []Genre  `json:"genres"`
	Studios       []Studio `json:"studios"`
	Images        Image    `json:"images"`
	Trailer       Trailer  `json:"trailer"`
}

type Genre struct {
	MalID int    `json:"mal_id"`
	Name  string `json:"name"`
}

type Studio struct {
	MalID int    `json:"mal_id"`
	Name  string `json:"name"`
}

type ImageURLs struct {
	ImageURL      string `json:"image_url"`
	LargeImageURL string `json:"large_image_url"`
}

type Image struct {
	JPG ImageURLs `json:"jpg"`
}

type Trailer struct {
	YoutubeID string `json:"youtube_id"`
	URL       string `json:"url"`
}

type Pagination struct {
	LastVisiblePage int   `json:"last_visible_page"`
	HasNextPage     bool  `json:"has_next_page"`
	CurrentPage     int   `json:"current_page"`
	Items           Items `json:"items"`
}

type Items struct {
	Count   int `json:"count"`
	Total   int `json:"total"`
	PerPage int `json:"per_page"`
}

type RecommendationEntry struct {
	MalID  int    `json:"mal_id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Images Image  `json:"images"`
}

type Recommendation struct {
	Entry RecommendationEntry `json:"entry"`
	Votes int                 `json:"votes"`
}

// Responses
type SingleAnimeResponse struct {
	Data Anime `json:"data"`
}

type AnimeListResponse struct {
	Pagination Pagination `json:"pagination"`
	Data       []Anime    `json:"data"`
}

type RecommendationResponse struct {
	Data []Recommendation `json:"data"`
}

type GenreListResponse struct {
	Data []Genre `json:"data"`
}
