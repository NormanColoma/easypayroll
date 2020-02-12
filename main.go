package main

import(
	"encoding/json"
	"log"
	"net/http"
)
import . "./domain"

func handle(w http.ResponseWriter, r *http.Request) {
	payroll := Payroll{
		CompanyName: "Facebook",
		Currency: "euros",
	}

	payroll.CalculatePayroll(Amount{2500, 2123.22, 2123.22, 2500})

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
