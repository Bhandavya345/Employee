package controllers

import (
	"encoding/json"
	"net/http"

	"strconv"
	"strings"

	models "github.com/Bhandavya345/Employee/model"
	services "github.com/Bhandavya345/Employee/service"
)

type ReportController struct {
	Service services.ReportService
}

func AddEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request", http.StatusMethodNotAllowed)
		return
	}

	var emp models.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = services.AddEmployee(&emp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	employees, err := services.GetEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(employees)
}

func EmployeeHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/employees/")
	parts := strings.Split(path, "/")

	employeeid, err := strconv.Atoi(parts[0])
	if err != nil {
		http.Error(w, "Invalid Employee ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:

		employee, err := services.GetEmployeeByID(employeeid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(employee)

	case http.MethodPut:

		if len(parts) == 2 && parts[1] == "salary" {

			var data struct {
				Salary float64 `json:"salary"`
			}

			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			err := services.UpdateSalary(employeeid, data.Salary)
			if err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Salary Updated"))
			return
		}

		http.Error(w, "Invalid URL", http.StatusBadRequest)

	case http.MethodDelete:

		err := services.DeleteEmployee(employeeid)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Employee Deleted"))

	default:

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func AverageSalary(w http.ResponseWriter, r *http.Request) {

	average := services.AverageSalary()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{
		"average_salary": average,
	})

}

func (c *ReportController) GenerateReport(w http.ResponseWriter, r *http.Request) {

	report, err := c.Service.GenerateReport()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}
