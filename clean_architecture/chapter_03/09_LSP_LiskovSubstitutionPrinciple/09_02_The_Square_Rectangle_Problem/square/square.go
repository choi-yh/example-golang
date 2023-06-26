package square

import "github.com/choi-yh/example-golang/clean_architecture/chapter_03/09_LSP_LiskovSubstitutionPrinciple/09_02_The_Square_Rectangle_Problem/rectangle"

type Square struct {
	rectangle.Rectangle
}

func NewSquare(side float32) Square {
	return Square{
		Rectangle: rectangle.Rectangle{
			Height: side,
			Width:  side,
		},
	}
}

func (s *Square) SetSide(side float32) {
	s.Width = side
	s.Height = side
}
