package presenter

import (
	"github.com/choi-yh/example-golang/TIL/clean_architecture/chapter_03/08_OCP_OpenClosedPrinciple/view"
)

type ScreenPresenter struct {
	webView view.WebView
}

func (p ScreenPresenter) ScreenView() {
	p.webView.ViewFinancialReportOnWeb()
}
