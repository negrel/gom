package gom

// NonDocumentTypeChildNode interface contains
// methods that are particular to Node objects
// that can have a parent, but not suitable
// for DocumentType.
// https://developer.mozilla.org/en-US/docs/Web/API/NonDocumentTypeChildNode
// https://dom.spec.whatwg.org/#interface-nondocumenttypechildnode
type NonDocumentTypeChildNode interface {
	/* GETTERS & SETTERS (props) */
	PreviousElementSibling() Element
	NextElementSibling() Element
}

var _ NonDocumentTypeChildNode = &nonDocumentTypeChildNode{}

type nonDocumentTypeChildNode struct {
	node
}

func (ndtcn *nonDocumentTypeChildNode) PreviousElementSibling() Element {
	var node Node = ndtcn

	for {
		prvSib := node.PreviousSibling()

		// If no previous element sibling this element is
		// the first child.
		if prvSib == nil {
			return nil
		}

		// Previous node is an element, so return it
		if prvSib.NodeType() == ElementNode {
			return prvSib.(Element)
		}

		// The previous node was not an element, retrying
		// with the previous sibling as the actual node.
		node = prvSib
	}
}

func (ndtcn *nonDocumentTypeChildNode) NextElementSibling() Element {
	var node Node = ndtcn

	for {
		prvSib := node.NextSibling()

		// If no next element sibling this element is
		// the last child.
		if prvSib == nil {
			return nil
		}

		// Previous node is an element, so return it
		if prvSib.NodeType() == ElementNode {
			return prvSib.(Element)
		}

		// The previous node was not an element, retrying
		// with the previous sibling as the actual node.
		node = prvSib
	}
}
