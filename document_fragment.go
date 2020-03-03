package gom

// TODO DocumentFragment

// DocumentFragment object represents a
// minimal document object that has no parent
// https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment
// https://dom.spec.whatwg.org/#documentfragment
type DocumentFragment interface {
	/* EMBEDDED INTERFACE */
	Node
}

type documentFragment struct {
	node
}

// NewDocumentFragment is the public constructor
// for the DocumentFragment object.
// https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment/DocumentFragment
func createDocumentFragment() DocumentFragment {
	// TODO func createDocumentFragment() DocumentFragment
	return &documentFragment{}
}
