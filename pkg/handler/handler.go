package handler

import (
	"github.com/Warh40k/gitmarket/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	project := router.Group("/projects")
	{
		project.GET("/search", h.searchProjects)
		project.GET("/:id", h.getProject)

		release := router.Group("/releases")
		{
			release.GET("/", h.getAllReleases)
			release.GET("/:id", h.getRelease)

			files := router.Group("/files")
			{
				files.GET("/", h.getReleaseFiles)
				files.GET("/", h.getFile)
			}
		}
	}

	return router
}
