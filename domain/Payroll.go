package domain

import (
	"sort"
)

var irpfSections = map[float32]float32{
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
	deduction.AddTax("IRPF", amount.Gross, calculateIRPF(amount.Gross))

	payroll.Deduction = deduction
}

func calculateIRPF(gross float32) float32 {
	var keys []float64

	for k := range irpfSections {
		keys = append(keys, float64(k))
	}
	sort.Float64s(keys)

	var firstQuota float32
	var lastSectionApplied float32
	grossAux := gross - 2000
	grossCalculated := grossAux

	for _, key := range keys {
		section := float32(key)

		if key == 12450 {
			firstQuota += section * irpfSections[section]
		} else if section > grossAux {
			firstQuota += (grossAux - lastSectionApplied) * irpfSections[section]
		} else {
			firstQuota += (section - lastSectionApplied) * irpfSections[section]
		}

		grossCalculated -= section - lastSectionApplied
		lastSectionApplied = section

		if grossCalculated <= 0 {
			break
		}
	}

	var secondQuota float32 = 5500 * 0.19

	return ((firstQuota - secondQuota) / gross) * 100
}


