package domain

var ranges = map[float32]float32{
	12450: 0.19,
	20200: 0.24,
	35200: 0.30,
	60000: 0.37,
	60001: 0.45,
}
type Amount struct {
	Gross float32
	Net float32
	Paid float32
	Earnings float32
}

type Payroll struct {
	Amount Amount
	CompanyName string
	Currency string
	Deduction Deduction
}

func (payroll *Payroll)CalculatePayroll(amount Amount) {
	payroll.Amount = amount

	deduction := Deduction{}
	deduction.AddTax("Cotización Desempleo", amount.Gross, UnemploymentTax)
	deduction.AddTax("Cotización Cotingencias Comunes", amount.Gross, CommonContingency)

	payroll.Deduction = deduction
}


