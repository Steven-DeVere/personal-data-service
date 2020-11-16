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
	router.GET("/projects", controllers.GetAllProjects)
	router.GET("/projects/:id", controllers.GetProject)
	router.POST("/projects", controllers.CreateProject)
	router.PUT("/projects", controllers.UpdateProject)
	router.DELETE("/projects/:id", controllers.DeleteProject)
}
