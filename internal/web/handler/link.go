package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/internal/service"
	"github.com/m-bromo/atom-ly/internal/web/models"
)

type LinkHandler struct {
	linkService service.LinkService
}

func (h *LinkHandler) Shorten(c *gin.Context) {
	var payload models.ShortenPayload
	if err := c.Bind(&payload); err != nil {
		c.Error(err)
		return
	}

	code, err := h.linkService.ShortenLink(c.Request.Context(), payload.Url)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, models.ShortenResponse{
		ShortLink: code,
	})
}

func (h *LinkHandler) Rediretct(c *gin.Context)
