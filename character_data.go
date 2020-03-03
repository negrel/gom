package gom

import "strings"

// The CharacterData abstract interface represents
// a Node object that contains characters.
// https://developer.mozilla.org/en-US/docs/Web/API/CharacterData
// https://dom.spec.whatwg.org/#interface-characterdata
type CharacterData interface {
	Node
	NonDocumentTypeChildNode
	/* GETTERS & SETTERS (props) */
	Data() string
	Length() int
	SetData(string)
	/* METHODS */
	AppendData(string) string
	DeleteData(uint, uint) string
	InsertData(uint, string) string
	ReplaceData(uint, uint, string)
	SubstringData(uint, uint) string
}

var _ CharacterData = &characterData{}

type characterData struct {
	*node
	*nonDocumentTypeChildNode
	data string
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

// Data return textual data contained in this object
// https://dom.spec.whatwg.org/#dom-characterdata-data
func (cd *characterData) Data() string {
	return cd.data
}

// Length return the size of the string contained in
// CharacterData.data.
// https://dom.spec.whatwg.org/#dom-characterdata-length
func (cd *characterData) Length() int {
	return len(cd.data)
}

// SetData set the textual data conatined in this object.
// https://dom.spec.whatwg.org/#dom-characterdata-data
func (cd *characterData) SetData(data string) {
	cd.data = data
}

/*****************************************************
 ********************* Methods ***********************
 *****************************************************/

func (cd *characterData) AppendData(data string) string {
	var b strings.Builder

	b.WriteString(cd.data)
	b.WriteString(data)

	cd.data = b.String()

	return cd.data
}

func (cd *characterData) DeleteData(offset, count uint) string {
	r := []rune(cd.data)
	r = append(r[offset:], r[offset+count:]...)

	cd.data = string(r)

	return cd.data
}

func (cd *characterData) InsertData(offset uint, data string) string {
	r := []rune(cd.data)
	d := []rune(data)

	r = append(r[offset:], d...)
	r = append(r, r[offset+1:]...)

	cd.data = string(r)

	return cd.data
}

func (cd *characterData) ReplaceData(offset, count uint, data string) {
	cd.DeleteData(offset, count)
	cd.InsertData(offset, data)
}

func (cd *characterData) SubstringData(offset, count uint) string {
	r := []rune(cd.data)

	return string(r[offset : offset+count])
}
