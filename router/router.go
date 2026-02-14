package router

import (
	"github.com/gin-gonic/gin"
	"github.com/polishedfeedback/aniverse/handlers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	anime := r.Group("/api/anime")
	anime.GET("/top", handlers.GetTopAnime)
	anime.GET("/seasonal", handlers.GetSeasonalAnime)
	anime.GET("/genres", handlers.GetGenres)
	anime.GET("/search", handlers.SearchAnime)
	anime.GET("/genre", handlers.GetAnimeByGenre)
	anime.GET("/:id", handlers.GetAnimeByID)
	anime.GET("/:id/recommendations", handlers.GetAnimeRecommendations)

	return r
}
