package domain

import (
	"sort"
)

const firstSection = 12450

var irpfSections = map[float32]float32{
	12450: 0.19,
	20200: 0.24,
	35200: 0.30,
	60000: 0.37,
	60001: 0.45,
}

var irpfByAge = map[int]int{
	64: 5500,
	65: 6700,
	75: 8100,
}

var irpfByKids = map[int]int {
	1: 2400,
	2: 2700,
	3: 4000,
	4: 4500,
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

	var firstIRPFQuota float32
	var lastSectionApplied float32
	grossAux := gross - 2000
	grossCalculated := grossAux

	for _, key := range keys {
		section := float32(key)

		if key == firstSection {
			firstIRPFQuota += section * irpfSections[section]
		} else if section > grossAux {
			firstIRPFQuota += (grossAux - lastSectionApplied) * irpfSections[section]
		} else {
			firstIRPFQuota += (section - lastSectionApplied) * irpfSections[section]
		}

		grossCalculated -= section - lastSectionApplied
		lastSectionApplied = section

		if grossCalculated <= 0 {
			break
		}
	}

	var secondIRPFQuota float32 = 5500 * 0.19

	return ((firstIRPFQuota - secondIRPFQuota) / gross) * 100
}

func discountIRPFByAge(age int) int {
	switch {
	case age >= 65 && age < 75:
		return irpfByAge[65]
	case age >= 75:
		return irpfByAge[75]
	default:
		return irpfByAge[64]
	}
}

func discountIRPFByKids(kids int) int {
	if val, ok := irpfByKids[kids]; ok {
		return val
	} else {
		return 4500
	}
}

