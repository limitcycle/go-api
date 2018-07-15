package models

import "errors"

var (
	INTERNAL_SERVER_ERROR = errors.New("Internal Server Error")
	NOT_FOUND_ERROR       = errors.New("Your request Item is not found")
	CONFLIT_ERROR         = errors.New("Your Item already exist")
)
