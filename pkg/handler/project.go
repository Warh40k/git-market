package handler

import (
	"github.com/Warh40k/gitmarket/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getSearchResultsResponse struct {
	Data []domain.Project `json:"data"`
}

func (h *Handler) searchProjects(c *gin.Context) {
	searchString := c.Param("search")
	if searchString == "" {
		newErrorResponse(c, http.StatusBadRequest, "no search param", nil)
		return
	}
	data, err := h.services.SearchRepos(searchString)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "error searching projects", err)
		return
	}
	c.JSON(http.StatusOK, getSearchResultsResponse{
		Data: data,
	})
}

func (h *Handler) getProject(c *gin.Context) {

}
