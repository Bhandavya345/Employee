package routes

import (
	"net/http"

	controllers "github.com/Bhandavya345/Employee/controller"
)

func RegisterRoutes() {

	ctrl := controllers.ReportController{}

	http.HandleFunc("/employees", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			controllers.AddEmployee(w, r)
		} else if r.Method == http.MethodGet {
			controllers.GetEmployees(w, r)
		}
	})

	http.HandleFunc("/employees/", controllers.EmployeeHandler)

	http.HandleFunc("/employees/average-salary", controllers.AverageSalary)

	http.HandleFunc("/reports", ctrl.GenerateReport)
}
