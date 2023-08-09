package dialog

import "github.com/choi-yh/example-golang/design_pattern/creational_pattern/factory_method/button"

type Dialog interface {
	CreateButton() button.Button
	Render()
}
