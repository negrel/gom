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
	/* GETTERS & SETTERS (props) */
	Attributes() NamedNodeMap
	ClassList() []string
	ClassName() string
	ClientHeight() int
	ClientWidth() int
	Id() *string
	InnerGOML() string
	LocalName() string
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

type element struct {
	attributes NamedNodeMap
	classList  []*string
	tagName    string
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

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
	// TODO func (e *element) ClientHeight() int
}

// ClientWidth return the inner width of an element
// in pixels. It includes padding but excludes margins.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/clientWidth
func (e *element) ClientWidth() int {
	// TODO func (e *element) ClientWidth() int
}

// Id return the id property of the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (e *element) Id() *string {
	return &e.Id
}

// InnerGOML return the GOML markup contained within the
// element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/innerHTML
func (e *element) InnerGOML() string {
	// TODO func (e *element) InnerGOML() *string
}

// LocalName return the local part of the qualified name of
// an element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/localName
func (e *element) LocalName() string {
	// TODO func (e *element) LocalName() string
}

// OuterGOML return the serialized GOML fragment describing
// the element including its descendants.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/outerHTML
func (e *element) OuterGOML() string {
	// TODO func (e *element) OuterGOML() *string
}

// ScrollHeight is a measurement of the height of an element's
// content, including content not visible on the screen due to
// overflow.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollHeight
func (e *element) ScrollHeight() int {
	// TODO func (e *element) ScrollHeight() int
}

// ScrollLeft return the number of pixels that an element's
// content is scrolled from its left edge.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (e *element) ScrollLeft() int {
	// TODO func (e *element) ScrollLeft() int
}

// ScrollTop return the number of pixels that an element's
// content is scrolled vertically.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollTop
func (e *element) ScrollTop() int {
	// TODO func (e *element) ScrollTop() int
}

// ScrollWidth is a measurement of the width of an element's
// content, including content not visible on the screen due
// to overflow.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollWidth
func (e *element) ScrollWidth() int {
	// TODO func (e *element) ScrollWidth() int
}

// SetClassName set the class attribute of the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (e *element) SetClassName(string) {
	// TODO func (e *element) SetClassName(string)
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
func (e *element) SetScrollTop(int) {
	// TODO func (e *element) SetScrollTop(int)
}

// SetScrollLeft set the number of pixels the top of
// the document is scrolled horizontally.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/scrollLeft
func (e *element) SetScrollLeft(int) {
	// TODO func (e *element) SetScrollLeft(int)
}

// TagName returns the tag name of the element on which
// it's called.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/tagName
func (e *element) TagName() string {
	// elements is not instanciable
	return nil
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

// GetAttribute return the value of a specified attribute
// on the element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttribute
func (e *element) GetAttribute(attrName string) Attr {
	// TODO func (e *element) GetAttribute(string) Attr
}

// GetAttributeNames returns an array of attribute names
// from the current element.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getAttributeName
func (e *element) GetAttributeNames() []string {
	// TODO func (e *element) GetAttributeName() []string
}

// GetElementsByClassName returns a live GOMLCollection which
// contains every descendant element which has
// the specified class name or names.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (e *element) GetElementsByClassName(className string) GOMLCollection {
	// TODO func (e *element) GetElementsByClassName() GOMLCollection
}

// GetElementsByTagName method returns a live GOMLCollection
// of elements with the given tag name.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (e *element) GetElementsByTagName(tagName string) GOMLCollection {
	// TODO func (e *element) GetElementsByTagName() GOMLCollection
}

// HasAttribute method returns a Boolean value indicating
// whether the specified element has the specified attribute
// or not.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute
func (e *element) HasAttribute(name string) bool {
	// TODO func (e *element) HasAttribute() bool
}

// QuerySelector method returns the first element that is
// a descendant of the element on which it is invoked that
// matches the specified group of selectors.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelector
func (e *element) QuerySelector(selector string) Node {
	// TODO func (e *element) QuerySelector(selector string) Node
}

// QuerySelectorAll returns a static (not live) NodeList
// representing a list of elements matching the specified
// group of selectors which are descendants of the element
// on which the method was called.
// https://developer.mozilla.org/en-US/docs/Web/API/Element/querySelectorAll
func (e *element) QuerySelectorAll(selector string) NodeList {
	// TODO func (e *element) QuerySelectorAll(selector string) Node
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
