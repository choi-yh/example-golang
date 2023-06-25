package interactor

type FinancialReportRequester struct{}

type FinancialDataGateway struct{}

type (
	FinancialReportRequest  struct{}
	FinancialReportResponse struct {
		FinancialReport
	}
)

type (
	FinancialData struct {
	}
	FinancialReport struct {
	}
)

type FinancialReportGenerator struct {
	FinancialReportRequester
	FinancialDataGateway
}
