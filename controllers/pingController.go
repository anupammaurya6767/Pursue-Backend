package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// PingController Function to check if the server is running
func PingController(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Pursuse Says Hello!")
}
