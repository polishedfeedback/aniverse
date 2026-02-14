package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/polishedfeedback/aniverse/config"
	"github.com/polishedfeedback/aniverse/models"
)

// Helper functions
// Fetch fetches the data and returns back the response
func Fetch(url string, target any) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// Response functions
// GetTopAnime uses helper fetch to fetch the data from the provided URL and returns response
func GetTopAnime() ([]models.Anime, error) {
	url := fmt.Sprintf("%s/top/anime", config.JikanBaseURL)
	var result models.AnimeListResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// SearchAnime takes in a query string and returns back the anime
func SearchAnime(query string) ([]models.Anime, error) {
	url := fmt.Sprintf("%s/anime?q=%s", config.JikanBaseURL, url.QueryEscape(query))

	var result models.AnimeListResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetSeasonalAnime returns the anime for that current season
func GetSeasonalAnime() ([]models.Anime, error) {
	url := fmt.Sprintf("%s/seasons/now", config.JikanBaseURL)

	var result models.AnimeListResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetGenre returns the anime genre
func GetGenres() ([]models.Genre, error) {
	url := fmt.Sprintf("%s/genres/anime", config.JikanBaseURL)
	var result models.GenreListResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetAnimeRecommendations
func GetAnimeRecommendations(id int) ([]models.Recommendation, error) {
	url := fmt.Sprintf("%s/anime/%d/recommendations", config.JikanBaseURL, id)
	var result models.RecommendationResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetAnimeByGenre takes in a genre and returns the anime for that genre
func GetAnimeByGenre(genreID int) ([]models.Anime, error) {
	url := fmt.Sprintf("%s/anime?genres=%d", config.JikanBaseURL, genreID)
	var result models.AnimeListResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}
	return result.Data, nil
}

// GetAnimeByID takes an ID and returns a pointer to *models.Anime
func GetAnimeByID(id int) (*models.Anime, error) {
	url := fmt.Sprintf("%s/anime/%d", config.JikanBaseURL, id)
	var result models.SingleAnimeResponse
	err := Fetch(url, &result)
	if err != nil {
		return nil, err
	}
	return &result.Data, nil
}
