package presenter

import "github.com/choi-yh/example-golang/clean_architecture/chapter_03/08_OCP_OpenClosedPrinciple/view"

type PrintPresenter struct {
	pdfView view.PDFView
}

func (p PrintPresenter) PrintView() {
	p.pdfView.ViewFinancialReportOnPDF()
}
