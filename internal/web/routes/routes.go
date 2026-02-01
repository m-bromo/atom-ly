package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/internal/web/handler"
	"github.com/m-bromo/atom-ly/internal/web/middleware"
)

func SetupRoutes(c *gin.Engine, h *handler.LinkHandler, m middleware.ErrorMiddleware) {
	c.POST("/shorten", h.Shorten, m.HandleErrors)
	c.GET("/:code", h.Rediretct, m.HandleErrors)
}
