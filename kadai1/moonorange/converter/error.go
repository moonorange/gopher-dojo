package converter

import (
	"errors"
	"fmt"
	"strings"

	"moonorange/dict"
)

type ErrCode string

// Conveter error error codes
var (
	DirCreateFail 	  ErrCode = "cannot create directory"
	FileCreateFail    ErrCode = "cannot create a file"
	FileEncodeFail    ErrCode = "cannot encode a file"
	InvalidFileExtCode ErrCode = "invalid file extension"
	SameExtCode ErrCode = "from and to flag value should be different from each other"
)

var (
	InvalidFileExtError error = fmt.Errorf("invalid file extension. supported format is " + strings.Join(dict.Keys(AvailFmt), " "))
	SameExtError error = errors.New("from and to flag value should be different from each other.")
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
