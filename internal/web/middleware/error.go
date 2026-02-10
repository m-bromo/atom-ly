package middleware

import (
	"github.com/gin-gonic/gin"
	resterrors "github.com/m-bromo/atom-ly/internal/web/rest_errors"
)

type ErrorMiddleware interface {
	HandleErrors(c *gin.Context)
}

type errorMiddleware struct {
	errorHandler *resterrors.ErrorHandler
}

func NewErrorMiddleware(errorHandler *resterrors.ErrorHandler) ErrorMiddleware {
	return &errorMiddleware{
		errorHandler: errorHandler,
	}
}

func (m *errorMiddleware) HandleErrors(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		restErr := m.errorHandler.HandleError(err)
		restErr.Path = c.Request.URL.RawPath
		c.JSON(restErr.Code, restErr)
	}
}
