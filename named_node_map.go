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
	SetNamedItem(string, string) Attr
	RemoveNamedItem(string) (Attr, GOMError)
}

var _ NamedNodeMap = &namedNodeMap{}

type namedNodeMap struct {
	list map[string]*attr
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

func (n *namedNodeMap) Length() int {
	return len(n.list)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

// GetNamedItem return the attribute corresponding to
// the given name.
func (n *namedNodeMap) GetNamedItem(name string) Attr {
	return n.list[name]
}

// SetNamedItem replace or adds the specified attribute
// with the given value.
func (n *namedNodeMap) SetNamedItem(name, value string) Attr {
	// Get the old attribute
	old := *n.list[name]

	// Set the new attribute value
	n.list[name].SetValue(value)

	// Return the old
	return &old
}

// RemoveNamedItem remove the specified attribute.
func (n *namedNodeMap) RemoveNamedItem(name string) (Attr, GOMError) {
	ptrAttr := n.list[name]
	attr := *ptrAttr

	// Check if attribute exist
	if ptrAttr == nil {
		return nil, newGOMError("The attr to be removed is not part of this element", "ErrNotFound")
	}

	delete(n.list, name)

	return &attr, nil
}
