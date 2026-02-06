package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/atom-ly/internal/service"
	resterrors "github.com/m-bromo/atom-ly/internal/web/rest_errors"
)

type ErrorMiddleware interface {
	HandleErrors(c *gin.Context)
}

type errorMiddleware struct {
}

func NewErrorMiddleware() ErrorMiddleware {
	return &errorMiddleware{}
}

func (m *errorMiddleware) HandleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err

		switch {
		case errors.Is(err, service.ErrUrlNotFound):
			restErr := resterrors.NewNotFoundError("url not found")
			restErr.Path = c.Request.URL.Path
			c.JSON(restErr.Code, restErr)

		default:
			restErr := resterrors.NewInternalServerError("There was an unexpecter internal server error")
			restErr.Path = c.Request.URL.Path
			c.JSON(restErr.Code, restErr)
		}

	}
}
