package gom

// GOMLDocument is a kind of document like
// HTMLDocument or XMLDocument. This document
// is used for non-web document.
type GOMLDocument struct {
	document
	title      string
	readyState bool
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
