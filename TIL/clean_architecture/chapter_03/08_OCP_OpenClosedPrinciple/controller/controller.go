package controller

import (
	"github.com/choi-yh/example-golang/TIL/clean_architecture/chapter_03/08_OCP_OpenClosedPrinciple/interactor"
	presenter2 "github.com/choi-yh/example-golang/TIL/clean_architecture/chapter_03/08_OCP_OpenClosedPrinciple/presenter"
)

type FinancialReportController struct {
	interactor.FinancialReportRequester
	FinancialReportPresenter
}

type FinancialReportPresenter struct {
	presenter2.ScreenPresenter
	presenter2.PrintPresenter
}

func (c FinancialReportController) GetFinancialReport() {

}

func (c FinancialReportController) ViewFinancialReportOnWeb() {
	c.ScreenPresenter.ScreenView()
}

func (c FinancialReportController) ViewFinancialReportOnPDF() {
	c.PrintPresenter.PrintView()
}
