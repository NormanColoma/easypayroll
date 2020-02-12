package domain

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

	unemploymentTax := Tax {
		Name: "Cotización Desempleo",
	}
	contingencyTax := Tax {
		Name: "Cotización Cotingencias Comunes",
	}
	unemploymentTax.CalculateTax(amount.Gross, UnemploymentTax)
	contingencyTax.CalculateTax(amount.Gross, CommonContingency)

	deduction := Deduction{}
	deduction.AddTax(unemploymentTax)
	deduction.AddTax(contingencyTax)

	payroll.Deduction = deduction
}
