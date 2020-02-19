package gom

// GOMRect represents a rectangle.
type GOMRect interface {
	/* GETTERS & SETTERS (props) */
	X() int
	Y() int
	Width() int
	Height() int
	Top() int
	Right() int
	Bottom() int
	Left() int
}

type gomRect struct {
	x      int
	y      int
	width  int
	height int
	top    int
	right  int
	bottom int
	left   int
}

/*****************************************************
 **************** Getters & Setters ******************
 *****************************************************/

func (r *gomRect) X() int {
	return r.x
}

func (r *gomRect) Y() int {
	return r.y
}

func (r *gomRect) Width() int {
	return r.width
}

func (r *gomRect) Height() int {
	return r.height
}

func (r *gomRect) Top() int {
	return r.top
}

func (r *gomRect) Right() int {
	return r.right
}

func (r *gomRect) Bottom() int {
	return r.bottom
}

func (r *gomRect) Left() int {
	return r.left
}
