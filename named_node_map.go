package gom

// NamedNodeMap interface represents a collection
// of Attr objects.
// https://developer.mozilla.org/en-US/docs/Web/API/NamedNodeMap
// https://dom.spec.whatwg.org/#interface-namednodemap
type NamedNodeMap interface {
	/* GETTERS & SETTERS (props) */
	Length() int
	/* METHODS */
	GetNamedItem(string) Attr
	SetNamedItem(string) Attr
	RemoveNamedItem(string)
}
