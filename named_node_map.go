package gom

// NamedNodeMap interface represents a collection
// of Attr objects.
// https://developer.mozilla.org/en-US/docs/Web/API/NamedNodeMap
// https://dom.spec.whatwg.org/#interface-namednodemap
type NamedNodeMap interface {
	/* Private */
	getNamedItem(string) (Attr, bool)
	/* GETTERS & SETTERS (props) */
	Length() int
	/* METHODS */
	GetNamedItem(string) Attr
	Item(int) Attr
	SetNamedItem(Attr)
	RemoveNamedItem(string) (Attr, GOMError)
}

var _ NamedNodeMap = &namedNodeMap{}

type namedNodeMap struct {
	list map[string]*attr
}

func (n *namedNodeMap) getNamedItem(name string) (Attr, bool) {
	attr, ok := n.list[name]
	return attr, ok
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

// Item returns the Attr at the given index, or null if
// the index is higher or equal to the number of nodes.
func (n *namedNodeMap) Item() Attr {

}

// SetNamedItem Replaces, or adds, the Attr identified
// in the map by the given name.
func (n *namedNodeMap) SetNamedItem(attr Attr) {
	// Set the new attribute value
	n.list[old.Name] = attr
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
