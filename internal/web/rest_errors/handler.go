package resterrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/m-bromo/atom-ly/internal/hasher"
	repository "github.com/m-bromo/atom-ly/internal/repository/link"
	"github.com/m-bromo/atom-ly/pkg/logger"
)

type ErrorHandler struct {
	log *logger.Logger
}

func NewErrorHandler(log *logger.Logger) *ErrorHandler {
	return &ErrorHandler{
		log: log,
	}
}

func (h *ErrorHandler) HandleError(err error) *RestErr {
	var validationErr validator.ValidationErrors

	switch {
	case errors.As(err, &validationErr):
		restErr := handleValidationErrors(err)
		h.log.Log.Warn(
			"validation failed for request",
			"error", err.Error(),
		)
		return restErr

	case errors.Is(err, repository.ErrLinkNotFound):
		restErr := NewNotFoundError("url not found")
		h.log.Log.Warn(
			"requested resource not found",
			"error", err.Error(),
		)
		return restErr

	case errors.Is(err, hasher.ErrInvalidCode):
		restErr := NewBadRequestError("the inserted code is invalid")
		h.log.Log.Warn(
			"invalid code provided by client",
			"error", err.Error(),
		)
		return restErr

	default:
		restErr := NewInternalServerError("there was an unexpected internal server error")
		h.log.Log.Error(
			"unexpected internal server error",
			"error", err.Error(),
		)
		return restErr
	}
}

func handleValidationErrors(err error) *RestErr {
	var causes []Causes

	if err != nil {
		if v, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range v {
				cause := Causes{
					Field:   fieldErr.ActualTag(),
					Message: fieldErr.Error(),
				}

				causes = append(causes, cause)
			}

			return NewBadRequestValidationError("fields are invalid", causes)
		}
	}

	return NewBadRequestError("some unexpected error has occurred")
}
