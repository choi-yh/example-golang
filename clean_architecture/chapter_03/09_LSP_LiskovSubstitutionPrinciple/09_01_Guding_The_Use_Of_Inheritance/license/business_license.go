package license

import "fmt"

type BusinessLicense struct{}

func (l BusinessLicense) CalcFee() {
	fmt.Println("Calculate Business License Fee")
	return
}

func (l BusinessLicense) users() {
	return
}

func NewBusinessLicense() BusinessLicense {
	return BusinessLicense{}
}
