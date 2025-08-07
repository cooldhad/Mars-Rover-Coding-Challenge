package domain

import "errors"

type badRequest struct {
	error
}

func (err badRequest) IsBadRequest() bool {
	return true
}

func IsBadRequestErr(err error) bool {
	var target interface {
		IsBadRequest() bool
	}

	return errors.As(err, &target) && target.IsBadRequest()
}

func AsBadRequestErr(err error) error {
	return badRequest{error: err}
}
