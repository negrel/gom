package gom

import (
	"strings"
)

/* NOTE Element missing props & methods (OFFICIAL DOM) :
 * ** Props **
 * clientLeft
 * clientTop
 * computedName
 * computedRole
 * localName
 * namespaceURI
 * part
 * prefix
 * all obsolete or non-standardized props
 *
 * ** Methods **
 * insertAdjacentElement
 * insertAdjacentHTML
 * insertAdjacentText
 * releasePointerCapture
 * removeAttributeNS
 * setPointerCapture
 * all obsolete or non-standardized methods
 */

// Element is most general base class from which all
// objects in a Document inherit.
// Element is not destined to be instancied but to be embedded.
// https://developer.mozilla.org/en-US/docs/Web/API/Element
// https://dom.spec.whatwg.org/#interface-element
type Element interface {
	/* EMBEDDED INTERFACE */
	Node
	NonDocumentTypeChildNode
	/* GETTERS & SETTERS (props) */
	Attributes() NamedNodeMap
	ClassList() []string
	ClassName() string
	ClientHeight() int
	ClientWidth() int
	Id() Attr
	InnerGOML() string
	OuterGOML() string
	ScrollHeight() int
	ScrollLeft() int
	ScrollTop() int
	ScrollWidth() int
	SetClassName(string)
	SetInnerGOML(string)
	SetOuterGOML(string)
	SetScrollTop(int)
	SetScrollLeft(int)
	TagName() string
	/* METHODS */
	GetAttribute(string) Attr
	GetAttributeNames() []string
	GetBoundingClientRect() GOMRect
	GetClientRects() []GOMRect
	GetElementsByClassName(string) GOMLCollection
	GetElementsByTagName(string) GOMLCollection
	HasAttribute(string) bool
	QuerySelector(string) Node
	QuerySelectorAll(string) NodeList
	RemoveAttribute(string)
	Scroll(x, y int)
	ScrollBy(x, y int)
	ScrollTo(x, y int)
	SetAttribute(name string, value interface{})
	ToggleAttribute(string)
}

var _ Element = &element{}
var _ Node = &element{}

type element struct {
	*node
	*nonDocumentTypeChildNode
	attributes NamedNodeMap
	classList  []string
	tagName    string
	id         Attr
}

func createElement(tagName string) Element {
	return &element{
		tagName: tagName,
	}
}

/*****************************************************
 **************** Embedded interface *****************
 *****************************************************/
// ANCHOR Embedded interface

/* Node */
/* - Props */

// NodeName return the GOML-uppercased name
func (e *element) NodeName() string {
	return strings.ToUpper(e.TagName())
}

// NodeType return the "ElementNode" type.
func (e *element) NodeType() NodeType {
	return ElementNode
}

/* - Methods */

// CloneNode return a clone of the element
func (e *element) CloneNode(deep bool) Node {
	clone := createElement(e.tagName)

	// Setting owner document
	clone.SetOwnerDocument(e.document)

	// Cloning all attributes
	eAttr := e.attributes.Values()
	for i, length := 0, e.attributes.Length(); i < length; i++ {
		eAttrClone := eAttr[i].CloneNode(false)

		clone.Attributes().SetNamedItem(eAttrClone.(Attr))
	}

	// If deep clone, cloning the children
	if deep {
		for _, child := range e.ChildNodes().Values() {
			clone.AppendChild(child.CloneNode(true))
		}
	}

	return clone
}

// IsEqualNode method return whether two Element are equal.
func (e *element) IsEqualNode(other Node) bool {
	if other == nil {
		goto notEqual
	}

	// Checking NodeType
	if e.NodeType() != other.NodeType() {
		goto notEqual
	}

	// Type switch
	switch otherEl := other.(type) {
	case Element:
		elAttrList := e.Attributes()
		otherElAttrList := otherEl.Attributes()

		// Checking attributes length
		if elAttrList.Length() != otherElAttrList.Length() {
			goto notEqual
		}

		otherElAttr := otherElAttrList.Values()
		// Check all attributes
		for i, child := range elAttrList.Values() {
			if otherAttr := otherElAttr[i]; otherAttr != nil {
				if !otherAttr.IsEqualNode(child) {
					goto notEqual
				}
			}
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

// Attributes returns a live collection of all
// attribute nodes registered to the specified node.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (e *element) Attributes() NamedNodeMap {
	return e.attributes
}

// ClassList return a live string collection of the
// class attributes element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/attributehttps://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (e *element) ClassList() (list []string) {
	return e.classList
}

// ClassName return the class attribute as a string.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *element) ClassName() string {
	return strings.Join(e.classList, " ")
}

// ClientHeight return the inner height of an element
// in pixels. It includes padding but excludes margins.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth
func (e *element) ClientHeight() int {
	return 0
}

// ClientWidth return the inner width of an element
// in pixels. It includes padding but excludes margins.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth
func (e *element) ClientWidth() int {
	return 0
}

// Id return the id property of the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (e *element) Id() Attr {
	return e.id
}

// InnerGOML return the GOML markup contained within the
// element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *element) InnerGOML() string {
	// TODO func (e *element) InnerGOML() *string
	return ""
}

// OuterGOML return the serialized GOML fragment describing
// the element including its descendants.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (e *element) OuterGOML() string {
	// TODO func (e *element) OuterGOML() *string
	return ""
}

// ScrollHeight is a measurement of the height of an element's
// content, including content not visible on the screen due to
// overflow.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollHeight
func (e *element) ScrollHeight() (sh int) {

	return sh
}

// ScrollLeft return the number of pixels that an element's
// content is scrolled from its left edge.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (e *element) ScrollLeft() (sl int) {

	return sl
}

// ScrollTop return the number of pixels that an element's
// content is scrolled vertically.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
func (e *element) ScrollTop() (st int) {

	return st
}

// ScrollWidth is a measurement of the width of an element's
// content, including content not visible on the screen due
// to overflow.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (e *element) ScrollWidth() (sw int) {

	return sw
}

// SetClassName set the class attribute of the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *element) SetClassName(className string) {
	e.classList = strings.Split(className, " ")
}

// SetInnerGOML set the GOML markup contained within
// the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *element) SetInnerGOML(string) {
	// TODO func (e *element) SetInnerGOML(string)
}

// SetOuterGOML set the serialized GOML fragment
// describing the element including its descendants.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (e *element) SetOuterGOML(string) {
	// TODO func (e *element) SetOuterGOML(string)
}

// SetScrollTop set the number of pixels the top of
// the document is scrolled vertically.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
func (e *element) SetScrollTop(int) {}

// SetScrollLeft set the number of pixels the top of
// the document is scrolled horizontally.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (e *element) SetScrollLeft(int) {}

// TagName returns the tag name of the element on which
// it's called.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (e *element) TagName() string {
	// elements is not instanciable
	return ""
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// GetAttribute return the value of a specified attribute
// on the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (e *element) GetAttribute(attrName string) Attr {
	attr, _ := e.attributes.getNamedItem(attrName)
	return attr
}

// GetAttributeNames returns an array of attribute names
// from the current element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttributeName
func (e *element) GetAttributeNames() []string {
	attrNames := make([]string, e.attributes.Length())

	for i := 0; i < e.attributes.Length(); i++ {
		attrNames = append(attrNames, e.attributes.Item(i).Name())
	}

	return attrNames
}

// GetBoundingClientRect method returns the size of an
// element and its position relative to the viewport.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (e *element) GetBoundingClientRect() GOMRect {
	// TODO (e *element) GetBoundingClientRect() GOMRect
	return &gomRect{}
}

// GetClientRect method returns the size of an
// element and its position relative to the viewport.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getBoundingClientRect
func (e *element) GetClientRects() []GOMRect {
	// TODO (e *element) GetClientRects() []GOMRect
	return make([]GOMRect, 0)
}

// GetElementsByClassName returns a live GOMLCollection which
// contains every descendant element which has
// the specified class name or names.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (e *element) GetElementsByClassName(className string) GOMLCollection {
	// TODO func (e *element) GetElementsByClassName() GOMLCollection
	return newGOMLCollection()
}

// GetElementsByTagName method returns a live GOMLCollection
// of elements with the given tag name.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (e *element) GetElementsByTagName(tagName string) GOMLCollection {
	// TODO func (e *element) GetElementsByTagName() GOMLCollection
	return newGOMLCollection()
}

// HasAttribute method returns a Boolean value indicating
// whether the specified element has the specified attribute
// or not.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute
func (e *element) HasAttribute(name string) bool {
	_, has := e.attributes.getNamedItem(name)
	return has
}

// QuerySelector method returns the first element that is
// a descendant of the element on which it is invoked that
// matches the specified group of selectors.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (e *element) QuerySelector(selector string) Node {
	// TODO func (e *element) QuerySelector(selector string) Node
	return newNode()
}

// QuerySelectorAll returns a static (not live) NodeList
// representing a list of elements matching the specified
// group of selectors which are descendants of the element
// on which the method was called.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (e *element) QuerySelectorAll(selector string) NodeList {
	// TODO func (e *element) QuerySelectorAll(selector string) Node
	return newNodeList()
}

// RemoveAttribute removes the attribute with the specified
// name from the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/removeAttribute
func (e *element) RemoveAttribute(attrName string) {
	// TODO func (e *element) RemoveAttribute(attrName string)
}

// Scroll method of the Element interface scrolls the element
// to a particular set of coordinates inside a given element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/Scroll
func (e *element) Scroll(x, y int) {
	// TODO func (e *element) Scroll(x, y int)
}

// ScrollBy method of the Element interface scrolls an element
// by the given amount.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/ScrollBy
func (e *element) ScrollBy(x, y int) {
	// TODO func (e *element) ScrollBy(x, y int)
}

// ScrollTo method of the Element interface scrolls to a particular
// set of coordinates inside a given element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/ScrollTo
func (e *element) ScrollTo(x, y int) {
	// TODO func (e *element) ScrollTo(x, y int)
}

// SetAttribute set the given attribute with the given value
func (e *element) SetAttribute(name string, value interface{}) {
	// TODO func (e *element) SetAttribute(name string, value interface{})
}

// ToggleAttribute method of the Element interface toggles a
// Boolean attribute (removing it if it is present and
// adding it if it is not present) on the given element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/toggleAttribute
func (e *element) ToggleAttribute(name string) {
	// TODO func (e *element) ToggleAttribute(name string)
}
