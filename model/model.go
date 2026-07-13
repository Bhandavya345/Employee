package models

type Employee struct {
	EmployeeID int     `gorm:"primaryKey;column:employee_id" json:"employeeid"`
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	Experience int     `json:"experience"`
}
