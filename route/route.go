package route

import (
	"be-awards/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // atau ganti "*" dengan domain tertentu
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(CORSMiddleware())

	voteRoutes := router.Group("/vote")
	{
		voteRoutes.POST("/create", controller.CreateVote)
		voteRoutes.POST("/getVote", controller.GetJumlahVote)
		voteRoutes.POST("/process", controller.VoteProcess)
	}

	nominasiRoutes := router.Group("/nominasi")
	{
		nominasiRoutes.GET("", controller.GetAllNominasi)
	}

	return router
}
