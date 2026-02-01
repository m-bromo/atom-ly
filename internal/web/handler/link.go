package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/internal/service"
	"github.com/m-bromo/atom-ly/internal/web/models"
	resterrors "github.com/m-bromo/atom-ly/internal/web/rest_errors"
)

type LinkHandler struct {
	linkService service.LinkService
}

func NewLinkHandler(linkService service.LinkService) *LinkHandler {
	return &LinkHandler{
		linkService: linkService,
	}
}

func (h *LinkHandler) Shorten(c *gin.Context) {
	var payload models.ShortenPayload
	if err := c.Bind(&payload); err != nil {
		restErr := resterrors.NewBadRequestError(err.Error())
		c.Error(restErr)
		return
	}

	code, err := h.linkService.ShortenLink(c.Request.Context(), payload.Url)
	if err != nil {
		restErr := resterrors.NewInternalServerError(err.Error())
		c.Error(restErr)
		return
	}

	c.JSON(http.StatusCreated, models.ShortenResponse{
		ShortLink: code,
	})
}

func (h *LinkHandler) Redirect(c *gin.Context) {
	url, err := h.linkService.Redirect(c.Request.Context(), c.Request.URL.Path[1:])
	if err != nil {
		restErr := resterrors.NewInternalServerError(err.Error())
		c.Error(restErr)
		return
	}

	c.Redirect(http.StatusPermanentRedirect, url)
}
