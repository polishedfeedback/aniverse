package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/polishedfeedback/aniverse/config"
	"github.com/polishedfeedback/aniverse/models"
)

// func GetAnimeByID(id int) (*models.Anime, error)

// func GetAnimeRecommendations(id int) ([]models.Anime, error)
// func GetGenres() ([]models.Genre, error)
// func GetAnimeByGenre(genreID int) ([]models.Anime, error)

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
