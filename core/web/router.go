package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jyotirmoydotdev/openfy/auth"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/signup", auth.RegisterUser)
	router.POST("/login", auth.LoginUser)
	router.POST("/admin/signup", func(ctx *gin.Context) {
		hashAdmin := auth.HashAdmin()
		if hashAdmin {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "Admin signup is not allowed", // TODO: rephrase error message
			})
			return
		}
		auth.RegisterAdmin(ctx)
	})
	router.POST("/admin/login", auth.LoginAdmin)

	router.GET("/products", GetAllProducts)
	router.GET("/products/:id", GetProduct)

	authGroup := router.Group("/admin", auth.AuthenticateMiddleware())
	{
		authGroup.POST("/products/new", Create)
		authGroup.PUT("/products/:id", Update)
		authGroup.DELETE("/products/:id", Delete)
	}

	return router
}