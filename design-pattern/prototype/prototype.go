package prototype

type Cloner interface {
	Clone() Cloner
}

type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius}
}

func (c Circle) Clone() Cloner {
	return NewCircle(c.radius)
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

type Rectangle struct {
	height int
	width  int
}

func NewRectangle(height, width int) *Rectangle {
	return &Rectangle{
		height: height,
		width:  width,
	}
}

func (r Rectangle) Clone() Cloner {
	return NewRectangle(r.height, r.width)
}

func (r Rectangle) Area() int {
	return r.height * r.width
}
