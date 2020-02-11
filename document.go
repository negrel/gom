package gom

import (
	"golang.org/x/text/encoding"
)

/* NOTE Document missing props & methods (OFFICIAL DOM) :
 * ** Props **
 * compatMode (experimental api)
 * contentType (experimental api)
 * documentURI
 * embeds
 * fonts
 * forms
 * images
 * implementation
 * lastStyleSheetSet
 * links
 * plugins
 * featurePolicy (experimental api)
 * preferredStyleSheetSet
 * scripts
 * scrollingElement
 * SelectedStyleSheetSet
 * styleSheetSets
 * timeline
 * all obsolete or non-standardized props
 *
 * ** Methods **
 * caretRangeFromPoint
 * createAttributeNS
 * createCDATASection
 * createElementNS
 * createEvent
 * createNodeIterator
 * createProcessingInstruction
 * createRange
 * createTouchList
 * createTreeWalker
 * enableStyleSheetsForSet
 * hasStorageAccess
 * requestStorageAccess
 * createExpression
 * createNSResolver
 * all obsolete or non-standardized methods
 */

// Document interface represents any page loaded
// and serves as an entry point into the page's
// content
// https://developer.mozilla.org/en-US/docs/Web/API/Document
// https://dom.spec.whatwg.org/#document
type Document interface {
	Node
	/* GETTERS & SETTERS (props) */
	Body() Node
	CharacterSet() encoding.Encoding
	DocType() DocumentType
	DocumentElement() Element
	Head() Element
	Hidden() bool
	SetBody(Node)
	SetCharacterSet(encoding.Encoding)
	/* METHODS */
	AdoptNode(Node)
	CreateAttribute() Node
	CreateComment() Node
	CreateDocumentFragment() Node
	CreateElement() Node
	CreateTextNode() Node
	GetElementsByClassName(string) Element
	GetElementsByTagName(string) Element
	ImportNode(Node, bool) Node
	GetElementById(string) Element
	QuerySelector(string) Element
	QuerySelectorAll(string) NodeList
}

var _ Document = &document{}

type document struct {
	node
	body            Node
	characterSet    encoding.Encoding
	docType         *DocumentType
	documentElement *element
	head            *element
	hidden          bool
	visibilityState string
}

// NewDocument return a new document object serving
// as an entry point into the page's content.
func NewDocument(name string) Document {
	return &document{
		node{},
		newNode(),
		nil,
		newDocumentType(name),
		nil, // TODO change for real documentElement
		nil, // TODO change for real head element
		false,
		"visible",
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

// Body return the <body> element of the current document
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (d *document) Body() Node {
	return d.body
}

// CharacterSet return the current character set used by
// the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (d *document) CharacterSet() encoding.Encoding {
	return d.characterSet
}

// DocType returns the Document Type Declaration (DTD)
// associated with current document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/doctype
func (d *document) DocType() DocumentType {
	return *d.docType
}

// DocumentElement returns the Element that is the root
// element of the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/documentElement
func (d *document) DocumentElement() Element {
	return d.documentElement
}

// Head return the <head> element of the current document
// https://developer.mozilla.org/en-US/docs/Web/API/Document/head
func (d *document) Head() Element {
	return d.head
}

// Hidden returns a Boolean value indicating if the page
// is considered hidden or not.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/hidden
func (d *document) Hidden() bool {
	return d.hidden
}

// SetBody set the body node of the document.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/body
func (d *document) SetBody(body Node) {
	d.body = body
}

// SetCharacterSet method set the document character
// set (UTF-8).
// https://developer.mozilla.org/en-US/docs/Web/API/Document/characterSet
func (d *document) SetCharacterSet(charSet encoding.Encoding) {
	d.characterSet = charSet
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

// AdoptNode transfers a node from another document
// into the document on which the method was called.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/adoptNode
func (d *document) AdoptNode(external Node) {
	// TODO func (d *document) AdoptNode(external Node)
}

// CreateAttribute method creates a new attribute node,
// and returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createAttribute
func (d *document) CreateAttribute() Node {
	// TODO func (d *document) CreateAttribute() Node
}

// CreateComment creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createComment
func (d *document) CreateComment() Node {
	// TODO func (d *document) CreateComment() Node
}

// CreateDocumentFragment creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createDocumentFragment
func (d *document) CreateDocumentFragment() Node {
	// TODO func (d *document) CreateDocumentFragment() Node
}

// CreateElement creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createElement
func (d *document) CreateElement() Node {
	// TODO func (d *document) QuerySelectorAll(selector string) NodeList
}

// CreateTextNode creates a new comment node, and
// returns it.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/createTextNode
func (d *document) CreateTextNode() Node {
	// TODO func (d *document) CreateTextNode() Node
}

// GetElementsByClassName method of Document interface
// returns an array-like object of all child elements
// which have all of the given class names.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByClassName
func (d *document) GetElementsByClassName(className string) Element {
	// TODO func (d *document) GetElementsByClassName(className string) Element
}

// GetElementsByTagName method of Document interface
// returns an array-like object of all child elements
// which have all of the given tag names.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementsByTagName
func (d *document) GetElementsByTagName(tagName string) Element {
	// TODO func (d *document) GetElementsByTagName(tagName string) Element
}

// ImportNode method creates a copy of a Node or
// DocumentFragment from another document, to be
// inserted into the current document later.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/importNode
func (d *document) ImportNode(node Node, deep bool) Node {
	// TODO func (d *document) ImportNode(node Node, deep bool) Node
}

// GetElementById returns an Element object representing
// the element whose id property matches the specified string
// https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (d *document) GetElementById(id string) Element {
	// TODO func (d *document) GetElementById(id string) Element
}

// QuerySelector returns the first Element within the document
// that matches the specified selector, or group of selectors.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelector
func (d *document) QuerySelector(selector string) Element {
	// TODO func (d *document) QuerySelector(selector string) Element
}

// QuerySelectorAll returns a static (not live) NodeList
// representing a list of the document's elements that match
// the specified group of selectors.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/querySelectorAll
func (d *document) QuerySelectorAll(selector string) NodeList {
	// TODO func (d *document) QuerySelectorAll(selector string) NodeList
}
