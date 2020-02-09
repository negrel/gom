package gom

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
// objects in a Document inherit
// https://developer.mozilla.org/en-US/docs/Web/API/Element
// https://dom.spec.whatwg.org/#interface-element
type Element interface {
	/* GETTERS & SETTERS (props) */
	Attributes() NamedNodeMap
	ClassList() []string
	ClassName() string
	ClientHeight() int
	ClientWidth() int
	Id() string
	InnerGOML() *string
	LocalName() string
	OuterGOML()
	ScrollHeight() int
	ScrollLeft() int
	ScrollTop() int
	ScrollWidth() int
	SetClassName(string)
	SetId(int)
	SetInnerGOML(string)
	SetOuterGOML(string)
	SetScrollTop(int)
	SetScrollLeft(int)
	TagName() string
	/* METHODS */
	DispatchEvent() bool
	GetElementsByClassName() GOMLCollection
	GetElementsByTagName() GOMLCollection
	HasAttribute() bool
	QuerySelector() Node
	QuerySelectorAll() NodeList
	RemoveAttribute()
	RemoveAttributeNode()
	Scroll(x, y int)
	ScrollBy(x, y int)
	ScrollTo(x, y int)
	SetAttribute(name, value string)
	ToggleAttribute(string)
}
