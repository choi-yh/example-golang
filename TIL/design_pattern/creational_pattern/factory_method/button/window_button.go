package button

type WindowButton struct {
}

func (i *WindowButton) Render() {

}

func (i *HTMLButton) OnClick() {

}

func NewHTMLButton() Button {
	return &HTMLButton{}
}
