package gom

// DocumentFragment object represents a
// minimal document object that has no parent
// https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment
// https://dom.spec.whatwg.org/#documentfragment
type DocumentFragment struct {
	Node
}

func newDocumentFragment