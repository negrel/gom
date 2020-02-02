package gom

// GOMError is the possible error
// returned by DOM function
// https://heycam.github.io/webidl/#idl-DOMException
type GOMError struct {
	error
	message string
	name    string
	code    int
}

// DOMError code list
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
	ErrNetwork
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
	"ErrNetwork":               ErrNetwork,
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
