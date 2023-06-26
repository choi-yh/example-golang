package license

type License interface {
	CalcFee()
}

type BusinessLicense struct{}
