package errors

import (
	"net/http"

	e "miniTwitter/pkg/errors"
)

var (
	ErrAccountNotExists       = e.NewError(http.StatusBadRequest, "account not exists")
	ErrAccountAlreadyExists   = e.NewError(http.StatusBadRequest, "account with this username already exists")
)
