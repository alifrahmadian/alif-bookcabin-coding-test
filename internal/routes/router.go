package routes

import (
	"github.com/alifrahmadian/alif-bookcabin-coding-test/configs"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, handlers *configs.Handler) {
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	router.GET("/seat-map/seats-itinerary-part/:id", handlers.SeatMapHandler.GetSeatMap)
}
