package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type responseError struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string, err error) {
	logrus.Error(message, err.Error())
	c.AbortWithStatusJSON(statusCode, responseError{message})
}
