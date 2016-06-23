package plalib

import (
// "errors"
)

type PlaErrorCode int

const (
	TARGET_RUN_ERROR PlaErrorCode = 1 + iota
	TARGET_NOT_FOUND
	PLAFILE_NOT_READABLE
)

type PlaError struct {
	Error error
	Code  PlaErrorCode
}
