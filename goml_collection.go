package gom

// GOMLCollection is live collection of elements
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection
// https://dom.spec.whatwg.org/#htmlcollection
type GOMLCollection struct {
	list []*Element
}

// Item return the element at the given index
// of the collection.
// https://dom.spec.whatwg.org/#dom-htmlcollection-item
func (c *GOMLCollection) Item(index int) *Element {
	if index >= 0 && index < c.Length() {
		return c.list[index]
	}
	return nil
}

// Length method return the number of elements in
// the collection
func (c *GOMLCollection) Length() int {
	return len(c.list)
}
