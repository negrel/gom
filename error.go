package gom

import (
	"fmt"
	"io"
)

// GOMError is the possible error returned
// by DOM function.
// https://heycam.github.io/webidl/#idl-DOMException
type GOMError interface {
	Code() int
	Message() string
	Name() string
	String() string
	Fprint(w io.Writer)
	Print()
}

var _ GOMError = &gomError{}

type gomError struct {
	message string
	name    string
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

func newGOMError(message string, name string) GOMError {
	return &gomError{
		message: message,
		name:    name,
	}
}

// Code method return the error code
func (e *gomError) Code() int {
	return errCode[e.name]
}

// Fprint method print the error to the given
// writer
func (e *gomError) Fprint(w io.Writer) {
	fmt.Fprint(w, e.String())
}

// Message method return the error message
func (e *gomError) Message() string {
	return e.message
}

// Name method return the error name
func (e *gomError) Name() string {
	return e.name
}

// Print the GOM Error
func (e *gomError) Print() {
	fmt.Print(e.String())
}

// String method return the formatted string
// error
func (e *gomError) String() string {
	return fmt.Sprintf("[%v] - %v", e.Code(), e.Message())
}
