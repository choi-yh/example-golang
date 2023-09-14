package dialog

import (
	button2 "github.com/choi-yh/example-golang/TIL/design_pattern/creational_pattern/factory_method/button"
)

type WebDialog struct {
	Dialog
	button *button2.HTMLButton
}

func (i *WebDialog) CreateButton() button2.Button {
	return i.button
}

func NewWebDialog(htmlButton *button2.HTMLButton) Dialog {
	return &WebDialog{
		button: htmlButton,
	}
}
