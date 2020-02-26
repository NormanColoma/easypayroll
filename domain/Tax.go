package domain

const UnemploymentTax = 1.55
const CommonContingency = 4.7

type Tax struct {
	Name string
	Percentage float32
	Total float32
}

func (tax *Tax) Calculate(gross float32, taxPercentage float32) {
	tax.Percentage = taxPercentage
	tax.Total = gross * (tax.Percentage / 100)
}
