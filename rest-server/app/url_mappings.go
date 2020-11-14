package app

import (
	"github.com/devere-here/personal-data-service/rest-server/controllers"
)

func mapUrls() {
	router.GET("ping", controllers.Ping)

	// article endpoints
	router.GET("/articles", controllers.GetAllArticles)
	router.GET("/articles/:id", controllers.GetArticle)
	router.POST("/articles", controllers.CreateArticle)
	router.PUT("/articles", controllers.UpdateArticle)
	router.DELETE("/articles/:id", controllers.DeleteArticle)

	// project endpoints
	router.GET("/projects", controllers.GetAllSideProjects)
	router.GET("/projects/:id", controllers.GetSideProject)
	router.POST("/projects", controllers.CreateSideProject)
	router.PUT("/projects", controllers.UpdateSideProject)
	router.DELETE("/projects/:id", controllers.DeleteSideProject)
}
