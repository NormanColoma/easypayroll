package domain

type Deduction struct {
	Total float32
	Taxes [] Tax
}

func (deduction *Deduction)AddTax(tax Tax) {
	deduction.Taxes = append(deduction.Taxes, tax)
	deduction.Total += tax.Total
}
