package gom

import (
	"golang.org/x/text/encoding"
)

/* TODO Document missing props & methods :
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
 * all obsolete props
 *
 * ** Methods **
 * all obsolete methods
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
