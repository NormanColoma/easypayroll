package main

import(
	"encoding/json"
	"log"
	"net/http"
)

import . "./domain"

type Deductions struct {
	Total float32
	Taxes [] Tax
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
	Deductions Deductions
}

func handle(w http.ResponseWriter, r *http.Request) {
	unemploymentTax := Tax {
		Name: "Cotización Desempleo",
	}
	contingencyTax := Tax {
		Name: "Cotización Cotingencias Comunes",
	}
	unemploymentTax.CalculateTax(3000, UnemploymentTax)
	contingencyTax.CalculateTax(3000, CommonContingency)

	payroll := Payroll{
		Amount: Amount{2500, 2123.22, 2123.22, 2500},
		CompanyName: "Facebook",
		Currency: "euros",
		Deductions: Deductions{
			Total: 429.10,
			Taxes: []Tax{
				unemploymentTax,
				contingencyTax,
			},
		},
	}

	response, err := json.Marshal(payroll)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func main() {
	http.HandleFunc("/api/", handle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
