package main

import(
	"encoding/json"
	"log"
	"net/http"
)

type Tax struct {
	Name string
	Percentage float32
	Total float32
}

type Deductions struct {
	Total float32
	Taxes []Tax
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

	payroll := Payroll{
		Amount: Amount{2500, 2123.22, 2123.22, 2500},
		CompanyName: "Facebook",
		Currency: "euros",
		Deductions: Deductions{
			Total: 429.10,
			Taxes: []Tax{
				{ Name: "IRPF", Percentage: 7, Total: 376.88},
				{ Name: "Seguridad Social", Percentage: 4.22, Total: 52.32 },
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
