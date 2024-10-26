package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomeHandler(g *gin.Context) {
	g.HTML(http.StatusOK, "index.html", nil)
}
