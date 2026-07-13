package models

type ReportResponse struct {
	DepartmentCount map[string]int     `json:"department_count"`
	AverageSalary   map[string]float64 `json:"average_salary"`
}
