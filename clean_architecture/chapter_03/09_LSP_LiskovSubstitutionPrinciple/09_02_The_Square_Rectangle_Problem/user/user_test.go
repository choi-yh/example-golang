package user

import (
	"testing"

	"github.com/choi-yh/example-golang/clean_architecture/chapter_03/09_LSP_LiskovSubstitutionPrinciple/09_02_The_Square_Rectangle_Problem/rectangle"
	"github.com/stretchr/testify/assert"
)

func TestUser_GetArea(t *testing.T) {
	r := rectangle.NewRectangle(1, 1)

	r.SetW(5)
	r.SetH(2)

	assert.Equal(t, float32(10), r.GetArea())
}
