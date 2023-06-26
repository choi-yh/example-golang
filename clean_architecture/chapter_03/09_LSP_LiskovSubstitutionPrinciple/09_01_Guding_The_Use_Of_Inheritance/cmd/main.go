package main

import (
	"github.com/choi-yh/example-golang/clean_architecture/chapter_03/09_LSP_LiskovSubstitutionPrinciple/09_01_Guding_The_Use_Of_Inheritance/billing"
	"github.com/choi-yh/example-golang/clean_architecture/chapter_03/09_LSP_LiskovSubstitutionPrinciple/09_01_Guding_The_Use_Of_Inheritance/license"
)

func main() {
	businessLicense := license.NewBusinessLicense()
	businessBilling := billing.NewBilling(businessLicense)
	businessBilling.CalcFee()

	personalLicense := license.NewPersonalLicense()
	personalBilling := billing.NewBilling(personalLicense)
	personalBilling.CalcFee()
}
