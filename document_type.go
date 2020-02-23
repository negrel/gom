package gom

// DocumentType represent a node containing a doctype.
// https://developer.mozilla.org/en-US/docs/Web/API/documentType
// https://dom.spec.whatwg.org/#documentType
type DocumentType interface {
	Node
	/* GETTERS & SETTERS (props) */
	Name() string
	PublicId() string
	SystemId() string
}

var _ DocumentType = &documentType{}

type documentType struct {
	*node
	name     string
	publicId string
	systemId string
}

func newDocumentType(name string) DocumentType {
	return &documentType{
		node: &node{
			nodeType: DocumentTypeNode,
		},
		name:     name,
		publicId: "",
		systemId: "",
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

// Name return the name of the document type
// eg "goml" for <!DOCTYPE GOML>
func (dt *documentType) Name() string {
	return dt.name
}

// PublicId return an empty string
func (dt *documentType) PublicId() string {
	return dt.publicId
}

// SystemId return an empty string
func (dt *documentType) SystemId() string {
	return dt.systemId
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

// IsEqualNode method return whether two DocumentType are equal.
func (dt *documentType) isEqualNode(other DocumentType) bool {
	if other == nil {
		goto notEqual
	}

	if dt.NodeType() != other.NodeType() {
		goto notEqual
	}

	// Document Type specific test
	if dt.name != other.Name() {
		goto notEqual
	}

	if dt.publicId != other.PublicId() {
		goto notEqual
	}

	if dt.systemId != other.SystemId() {
		goto notEqual
	}

	return true

notEqual:
	return false
}
