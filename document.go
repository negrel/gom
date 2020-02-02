package gom

import (
	"golang.org/x/text/encoding"
)

/* NOTE Document missing props & methods :
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
 * all obsolete or non-standardized methods
 */

// Document interface represents any page loaded
// and serves as an entry point into the page's
// content
// https://developer.mozilla.org/en-US/docs/Web/API/Document
// https://dom.spec.whatwg.org/#document
type Document struct {
	Node
	body            *Node
	characterSet    encoding.Encoding
	docType         DocumentType
	documentElement Element
	head            Element
	hidden          bool
	visibilityState string
}

// AdoptNode transfers a node from another
// document into the document on which the
// method was called.
// https://developer.mozilla.org/en-US/docs/Web/API/Document/adoptNode
func (d *Document) AdoptNode(other Node) {

}
