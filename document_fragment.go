package gom

// TODO DocumentFragment

// DocumentFragment object represents a
// minimal document object that has no parent
// https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment
// https://dom.spec.whatwg.org/#documentfragment
type DocumentFragment struct {
	Node
}

// NewDocumentFragment is the public constructor
// for the DocumentFragment object.
// https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment/DocumentFragment
func NewDocumentFragment() *DocumentFragment {
	// TODO func NewDocumentFragment() *DocumentFragment
	return nil
}
