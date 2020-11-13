package app

import (
	"github.com/devere-here/personal-data-service/rest-server/controllers"
)

func mapUrls() {
	router.GET("ping", controllers.Ping)

	router.GET("/articles", controllers.GetAllArticles)
	router.GET("/articles/:id", controllers.GetArticle)
	router.POST("/articles", controllers.CreateArticle)
	router.PUT("/articles", controllers.UpdateArticle)
	router.DELETE("/articles/:id", controllers.DeleteArticle)
}