package button

type HTMLButton struct {
}

func (i *HTMLButton) Render() {

}

func (i *WindowButton) OnClick() {

}

func NewWindowButton() Button {
	return &WindowButton{}
}
