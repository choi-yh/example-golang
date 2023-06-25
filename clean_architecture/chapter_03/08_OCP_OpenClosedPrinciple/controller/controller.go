package controller

import (
	"github.com/choi-yh/example-golang/clean_architecture/chapter_03/08_OCP_OpenClosedPrinciple/interactor"
	"github.com/choi-yh/example-golang/clean_architecture/chapter_03/08_OCP_OpenClosedPrinciple/presenter"
)

type FinancialReportController struct {
	interactor.FinancialReportRequester
	FinancialReportPresenter
}

type FinancialReportPresenter struct {
	presenter.ScreenPresenter
	presenter.PrintPresenter
}

func (c FinancialReportController) GetFinancialReport() {

}

func (c FinancialReportController) ViewFinancialReportOnWeb() {
	c.ScreenPresenter.ScreenView()
}

func (c FinancialReportController) ViewFinancialReportOnPDF() {
	c.PrintPresenter.PrintView()
}
