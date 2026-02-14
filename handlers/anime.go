package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/polishedfeedback/aniverse/models"
	"github.com/polishedfeedback/aniverse/services"
)

// Pattern 1
func GetTopAnime(c *gin.Context) {
	anime, err := services.GetTopAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting anime"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": anime})
}

func GetSeasonalAnime(c *gin.Context) {
	anime, err := services.GetSeasonalAnime()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting seasonal anime"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": anime})
}

func GetGenres(c *gin.Context) {
	genres, err := services.GetGenres()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting genres"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": genres})
}

// Pattern 2
func GetAnimeByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "bad request"})
		return
	}
	anime, err := services.GetAnimeByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting anime"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": anime})
}

func GetAnimeRecommendations(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "bad request"})
		return
	}
	anime, err := services.GetAnimeRecommendations(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting recommendations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": anime})
}

// Pattern 3
func SearchAnime(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "search query is required"})
		return
	}
	anime, err := services.SearchAnime(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting anime"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": anime})
}

func GetAnimeByGenre(c *gin.Context) {
	genreID, err := strconv.Atoi(c.Query("genre"))
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "invalid genre id"})
		return
	}
	anime, err := services.GetAnimeByGenre(genreID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "error getting anime"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": anime})
}
