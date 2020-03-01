package gom

// GOMLCollection is live collection of elements
// https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection
// https://dom.spec.whatwg.org/#htmlcollection
type GOMLCollection interface {
	Item(int) Element
}

type gomlCollection struct {
	list []Element
}

func newGOMLCollection() *gomlCollection {
	return &gomlCollection{
		list: []Element{},
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

// Length method return the number of elements in
// the collection
func (c *gomlCollection) Length() int {
	return len(c.list)
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

// Item return the element at the given index
// of the collection.
// https://dom.spec.whatwg.org/#dom-htmlcollection-item
func (c *gomlCollection) Item(index int) Element {
	if index >= 0 && index < c.Length() {
		return c.list[index]
	}
	return nil
}
