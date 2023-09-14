package dialog

import (
	button2 "github.com/choi-yh/example-golang/TIL/design_pattern/creational_pattern/factory_method/button"
)

type WindowsDialog struct {
	Dialog
	button *button2.WindowButton
}

func (i *WindowsDialog) CreateButton() button2.Button {
	return i.button
}

func (i *WindowsDialog) Render() {
	i.button.OnClick()
	i.button.Render()
}

func NewWindowDialog(windowButton *button2.WindowButton) Dialog {
	return &WindowsDialog{
		button: windowButton,
	}
}
