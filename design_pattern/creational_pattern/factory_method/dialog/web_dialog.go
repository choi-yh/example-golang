package dialog

import "github.com/choi-yh/example-golang/design_pattern/creational_pattern/factory_method/button"

type WebDialog struct {
	Dialog
	button *button.HTMLButton
}

func (i *WebDialog) CreateButton() button.Button {
	return i.button
}

func NewWebDialog(htmlButton *button.HTMLButton) Dialog {
	return &WebDialog{
		button: htmlButton,
	}
}
