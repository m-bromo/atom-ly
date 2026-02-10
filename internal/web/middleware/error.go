package middleware

import (
	"github.com/gin-gonic/gin"
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
		restErr := resterrors.HandleError(err)
		c.JSON(restErr.Code, restErr)
	}
}
