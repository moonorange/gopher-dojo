package converter

import (
	"fmt"
)

type ErrCode string

// Conveter error error codes
var (
	DirCreateFail 	  ErrCode = "cannot create directory"
	FileCreateFail    ErrCode = "cannot create a file"
	FileEncodeFail    ErrCode = "cannot encode a file"
)

// Error type of converter
type ConvError struct {
	Err error
	Code ErrCode
	FilePath string
}

// Wrapped error for ConvError
func (e *ConvError) Error() string {
	return fmt.Sprintln(e.Code, e.FilePath)
}

// Unwrap wrapped error to see actual library error
func (e *ConvError) Unwrap() error { return e.Err }
