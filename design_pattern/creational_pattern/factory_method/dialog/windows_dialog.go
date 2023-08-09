package dialog

import "github.com/choi-yh/example-golang/design_pattern/creational_pattern/factory_method/button"

type WindowsDialog struct {
	Dialog
	button *button.WindowButton
}

func (i *WindowsDialog) CreateButton() button.Button {
	return i.button
}

func (i *WindowsDialog) Render() {
	i.button.OnClick()
	i.button.Render()
}

func NewWindowDialog(windowButton *button.WindowButton) Dialog {
	return &WindowsDialog{
		button: windowButton,
	}
}
