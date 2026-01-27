package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/internal/web/handler"
)

func SetupRoutes(c *gin.Engine, h *handler.LinkHandler) {
	c.POST("/shorten", h.Shorten)
	c.GET("/:code", h.Rediretct)
}
