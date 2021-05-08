package initialize

import (
	"github.com/akazwz/go-gin-restful-api/middleware"
	"github.com/akazwz/go-gin-restful-api/model/response"
	"github.com/akazwz/go-gin-restful-api/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func Routers() *gin.Engine {
	var router = gin.Default()
	router.Static("/public", "./public")
	//cors
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:    []string{"*"},
	}))

	router.NoRoute(response.NotFound)
	//go-swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Teapot
	router.GET("teapot", func(c *gin.Context) {
		c.JSON(http.StatusTeapot, gin.H{
			"message": "I'm a teapot",
			"story": "This code was defined in 1998 " +
				"as one of the traditional IETF April Fools' jokes," +
				" in RFC 2324, Hyper Text Coffee Pot Control Protocol," +
				" and is not expected to be implemented by actual HTTP servers." +
				" However, known implementations do exist.",
		})
	})

	publicRouterV1 := router.Group("v1")
	{
		routers.InitBaseRouter(publicRouterV1)
	}

	privateGroupV1 := router.Group("v1")
	privateGroupV1.Use(middleware.JWTAuth())
	{
		routers.InitBookRouter(privateGroupV1)
		routers.InitFileRouter(privateGroupV1)
		routers.InitUserRouter(privateGroupV1)
	}

	return router
}
