package gom

// Attr interface represents one of a DOM element's
// attributes as an object.
// https://developer.mozilla.org/en-US/docs/Web/API/Attr
// https://dom.spec.whatwg.org/#attr
type Attr interface {
	Name() string
	OwnerElement() Element
	Value() string
}

type attr struct {
	name         string
	ownerElement Element
	value        string
}

// Name return the attribute name of an element
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/localName
func (a *attr) Name() string {
	return a.name
}

// OwnerElement return the element holding the attribute
func (a *attr) OwnerElement() Element {
	return a.ownerElement
}

// Value return the attribute value of an element
func (a *attr) Value() string {
	return a.value
}
