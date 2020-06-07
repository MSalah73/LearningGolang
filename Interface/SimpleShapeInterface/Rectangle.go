package SimpleShapeInterface

type Rectangle struct {
	width  float64
	height float64
}

func NewRectangle(w float64, h float64) *Rectangle {
	return &Rectangle{width: w, height: h}
}

// Notice: The (r * Rectangle) is called a pointer receiver - Area
// of type Rectangle witch impelments the interface Shape
// https://www.yellowduck.be/posts/pointer-vs-value-receivers/
func (r *Rectangle) Area() float64 {
	return r.height * r.width
}
