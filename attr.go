package gom

import "strings"

// Attr interface represents one of a DOM element's
// attributes as an object.
// https://developer.mozilla.org/en-US/docs/Web/API/Attr
// https://dom.spec.whatwg.org/#attr
type Attr interface {
	/* EMBEDDED INTERFACE */
	Node
	/* GETTERS & SETTERS (props) */
	Name() string
	OwnerElement() Element
	Value() string
	SetValue(string)
}

var _ Attr = &attr{}
var _ Node = &attr{}

type attr struct {
	*node
	name         string
	ownerElement Element
	value        string
}

func createAttribute(name string) Attr {
	return &attr{
		node:         &node{},
		name:         strings.ToLower(name),
		ownerElement: nil,
		value:        "",
	}
}

/*****************************************************
 **************** Embedded interface *****************
 *****************************************************/
// ANCHOR Embedded interface

/* Node */
/* - Props */

func (a *attr) NodeName() string {
	return a.Name()
}

// NodeType return the "AttributeNode" type.
func (a *attr) NodeType() NodeType {
	return AttributeNode
}

/* - Methods */

// CloneNode return a clone of the Attr
func (a *attr) CloneNode(_ bool) Node {
	clone := createAttribute(a.name)

	clone.SetValue(a.value)

	return clone
}

// IsEqualNode return wether or not two Attr are equal
func (a *attr) IsEqualNode(other Node) bool {
	if other == nil {
		goto notEqual
	}

	// Checking NodeType
	if a.NodeType() != other.NodeType() {
		goto notEqual
	}

	// Type switch
	switch otherAttr := other.(type) {
	case Attr:
		if a.Name() != otherAttr.Name() {
			goto notEqual
		}

		if a.Value() != otherAttr.Value() {
			goto notEqual
		}

	default:
		goto notEqual
	}

	return true

notEqual:
	return false
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

// Name return the attribute name of an element
// https://developer.mozilla.org/en-US/docs/Web/API/Attr/localName
func (a *attr) Name() string {
	return a.name
}

// OwnerElement return the element holding the attribute
func (a *attr) OwnerElement() Element {
	return a.ownerElement
}

// Value return the attribute value of an element
func (a *attr) Value() string {
	return a.value
}

// SetValue set the attribute value of an element
func (a *attr) SetValue(value string) {
	a.value = value
}
