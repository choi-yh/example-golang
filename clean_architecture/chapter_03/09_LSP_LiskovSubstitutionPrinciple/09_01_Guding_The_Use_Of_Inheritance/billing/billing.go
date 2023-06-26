package billing

import "github.com/choi-yh/example-golang/clean_architecture/chapter_03/09_LSP_LiskovSubstitutionPrinciple/09_01_Guding_The_Use_Of_Inheritance/license"

type Billing struct {
	license.License
}

func (b Billing) CalcFee() {
	b.License.CalcFee()
}

func NewBilling(license license.License) Billing {
	return Billing{
		License: license,
	}
}
