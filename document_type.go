package gom

// DocumentType represent a node
// containing a doctype.
// https://developer.mozilla.org/en-US/docs/Web/API/DocumentType
// https://dom.spec.whatwg.org/#documenttype
type DocumentType struct {
	name     string
	publicID string
	systemID string
}

func newDocumentType(name string) *DocumentType {
	return &DocumentType{
		name:     name,
		publicID: "",
		systemID: "",
	}
}
