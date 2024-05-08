package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controllers "github.com/kundu-ramit/mercor_assignment/controller"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	controller := controllers.NewQueryController()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "healthy",
		})
	})

	router.GET("/query/nlp", controller.ProcessNLPQuery)

	return router
}
