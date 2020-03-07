package gom

import (
	"github.com/negrel/gom/exception"
	e "github.com/negrel/gom/exception"
)

/* NOTE Element missing props & methods (OFFICIAL DOM) :
 * ** Props **
 * assignedSlot
 * all obsolete or non-standardized props
 *
 * ** Methods **
 * all obsolete or non-standardized methods
 */

// Text interface represents the textual content
// of Element or Attr.
// https://developer.mozilla.org/en-US/docs/Web/API/Text
// https://dom.spec.whatwg.org/#interface-text
type Text interface {
	/* EMBEDDED INTERFACE */
	CharacterData
	/* GETTERS & SETTERS (props) */
	WholeText() string
	/* METHODS */
	SplitText(uint) (Text, exception.Exception)
}

var _ Text = &text{}

type text struct {
	*characterData
}

// createTextNode return a new Text node.
func createTextNode(content string) Text {
	return &text{
		&characterData{
			data: content,
		},
	}
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/
// ANCHOR Getters & Setters

func (t *text) WholeText() string {
	return t.Data()
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/
// ANCHOR Methods

// SplitText method breaks the Text node into two nodes
// at the specified offset, keeping both nodes in the
// tree as siblings.
// https://developer.mozilla.org/en-US/docs/Web/API/Text/splitText
// https://dom.spec.whatwg.org/#dom-text-splittext
func (t *text) SplitText(offset uint) (Text, e.Exception) {
	var count uint = uint(t.Length()) - offset

	// If offset is greater than length
	if count < 0 {
		return nil, e.RangeError("The offset %v is larger than the Text node's length.", string(offset))
	}

	// Substring is data of the new node
	newTextData := t.SubstringData(offset, count)
	// Creating new node
	newText := createTextNode(newTextData)
	newText.SetOwnerDocument(t.OwnerDocument())

	if parent := t.ParentNode(); parent != nil {
		parent.InsertBefore(newText, t.NextSibling())

	}

	// Deleting data of the current Text node
	t.DeleteData(offset, count)

	return newText, nil
}
