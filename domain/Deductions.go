package domain

type Deduction struct {
	Total float32
	Taxes [] Tax
}

func (deduction *Deduction)AddTax(name string, gross float32, percentage float32) {
	tax := Tax {
		Name: name,
	}

	tax.Calculate(gross, percentage)
	deduction.Taxes = append(deduction.Taxes, tax)
	deduction.Total += tax.Total
}
