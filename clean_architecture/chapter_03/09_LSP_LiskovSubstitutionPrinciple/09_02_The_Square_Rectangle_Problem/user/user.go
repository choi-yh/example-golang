package user

import "github.com/choi-yh/example-golang/clean_architecture/chapter_03/09_LSP_LiskovSubstitutionPrinciple/09_02_The_Square_Rectangle_Problem/rectangle"

type User struct {
	rectangle.Rectangle
}

func NewUser(rectangle rectangle.Rectangle) User {
	return User{
		Rectangle: rectangle,
	}
}
