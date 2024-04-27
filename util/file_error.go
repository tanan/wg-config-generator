package util

import (
	"fmt"
)

type FileError struct {
	kind  FileErrorType
	Path  string
	inner error
}

type FileErrorType string

const (
	FileNotFound         = FileErrorType("FileNotFound")
	FileReadFailure      = FileErrorType("FileReadFailure")
	FileUnmarshalFailure = FileErrorType("FileUnmarshalFailure")
	FileWriteFailure     = FileErrorType("FileWriteFailure")
	PermissionDenied     = FileErrorType("PermissionDenied")
)

func NewFileError(kind FileErrorType, path string, err error) FileError {
	return FileError{
		kind:  kind,
		Path:  path,
		inner: err,
	}
}

func (err FileError) Is(target error) bool {
	t, ok := target.(FileError)
	return ok && err.kind == t.kind
}

func (fe FileError) Error() string {
	return fmt.Sprintf("File error. path: %s, err: %v", fe.Path, fe.inner)
}

func (fe FileError) Unwrap() error {
	return fe.inner
}
