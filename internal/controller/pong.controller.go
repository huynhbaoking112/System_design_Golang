package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func GetPongController() *PongController {
	return &PongController{}
}

func (pc *PongController) Pong(c *gin.Context) {
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"uid":     uid,
		"users": []string{
			"cr7",
			"m10",
			"king",
		},
	})
}
