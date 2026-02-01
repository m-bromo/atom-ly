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

		if restErr, ok := err.(*resterrors.RestErr); ok {
			restErr.Path = c.Request.URL.Path
			c.JSON(restErr.Code, restErr)
		} else {
			restErr := resterrors.NewInternalServerError("There was an unexpecter internal server error")
			restErr.Path = c.Request.URL.Path

			c.JSON(restErr.Code, restErr)
		}
	}
}
