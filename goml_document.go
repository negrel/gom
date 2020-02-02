package gom

// GOMLDocument is a kind of document like
// HTMLDocument or XMLDocument. This document
// is used for non-web document.
type GOMLDocument struct {
	Document
	title      string
	readyState bool
}
