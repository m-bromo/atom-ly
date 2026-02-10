package resterrors

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/m-bromo/atom-ly/internal/hasher"
	repository "github.com/m-bromo/atom-ly/internal/repository/link"
)

func HandleError(err error) *RestErr {
	var validationErr validator.ValidationErrors

	switch {
	case errors.As(err, &validationErr):
		restErr := handleValidationErrors(err)
		return restErr

	case errors.Is(err, repository.ErrLinkNotFound):
		restErr := NewNotFoundError("url not found")
		return restErr

	case errors.Is(err, hasher.ErrInvalidCode):
		restErr := NewBadRequestError("the inserted code is invalid")
		return restErr

	default:
		restErr := NewInternalServerError("There was an unexpecter internal server error")
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
