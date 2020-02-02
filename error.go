package gom

import (
	"fmt"
	"io"
)

// GOMError is the possible error
// returned by DOM function
// https://heycam.github.io/webidl/#idl-DOMException
type GOMError struct {
	error
	message string
	name    string
	code    int
}

// GOMError code list
const (
	_ = iota
	ErrHierarchyRequest
	ErrWrongDocument
	ErrInvalidCharacter
	ErrNoModificationAllowed
	ErrNotFound
	ErrNotSupported
	ErrInUseAttribute
	ErrInvalidState
	ErrSyntax
	ErrInvalidModification
	ErrSecurity
	ErrAbort
	ErrQuotaExceeded
	ErrTimeout
	ErrInvalidNodeType
	ErrDataClone
)

var errCode = map[string]int{
	"ErrHierarchyRequest":      ErrHierarchyRequest,
	"ErrWrongDocument":         ErrWrongDocument,
	"ErrInvalidCharacter":      ErrInvalidCharacter,
	"ErrNoModificationAllowed": ErrNoModificationAllowed,
	"ErrNotFound":              ErrNotFound,
	"ErrNotSupported":          ErrNotSupported,
	"ErrInUseAttribute":        ErrInUseAttribute,
	"ErrInvalidState":          ErrInvalidState,
	"ErrSyntax":                ErrSyntax,
	"ErrInvalidModification":   ErrInvalidModification,
	"ErrSecurity":              ErrSecurity,
	"ErrAbort":                 ErrAbort,
	"ErrQuotaExceeded":         ErrQuotaExceeded,
	"ErrTimeout":               ErrTimeout,
	"ErrInvalidNodeType":       ErrInvalidNodeType,
	"ErrDataClone":             ErrDataClone,
}

func newGOMError(message string, name string) *GOMError {
	return &GOMError{
		message: message,
		name:    name,
	}
}

// Code method return the error code
func (e *GOMError) Code() int {
	return errCode[e.name]
}

// Message method return the error message
func (e *GOMError) Message() string {
	return e.message
}

// Name method return the error name
func (e *GOMError) Name() string {
	return e.name
}

// String method return the formatted string
// error
func (e *GOMError) String() string {
	return fmt.Sprintf("[%v] - %v", e.Code(), e.Message())
}

// Print the GOM Error
func (e *GOMError) Print() {
	fmt.Print(e.String())
}

// Fprint method print the error to the given
// writer
func (e *GOMError) Fprint(w io.Writer) {
	fmt.Fprint(w, e.String())
}
