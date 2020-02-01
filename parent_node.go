package dom

// ParentNode mixin contains methods and properties
// that are common to all types of Node objects that can have children
type ParentNode struct {
	childElementCount int
	children          interface{}
	firstElementchild interface{}
	lastElementchild  interface{}
}

func (pn *ParentNode) append()           {}
func (pn *ParentNode) prepend()          {}
func (pn *ParentNode) querySelector()    {}
func (pn *ParentNode) querySelectorAll() {}
