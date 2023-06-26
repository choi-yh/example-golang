package rectangle

type Rectangle struct {
	Height float32
	Width  float32
}

func NewRectangle(height, width float32) Rectangle {
	return Rectangle{
		Height: height,
		Width:  width,
	}
}

func (r *Rectangle) SetH(height float32) {
	r.Height = height
}

func (r *Rectangle) SetW(width float32) {
	r.Width = width
}

func (r *Rectangle) GetArea() float32 {
	return r.Height * r.Width
}
