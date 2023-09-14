package license

import "fmt"

type PersonalLicense struct{}

func (l PersonalLicense) CalcFee() {
	fmt.Println("Calculate Personal License Fee")
	return
}

func NewPersonalLicense() PersonalLicense {
	return PersonalLicense{}
}
