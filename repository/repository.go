package repository

import (
	"errors"

	database "github.com/Bhandavya345/Employee/DB"
	models "github.com/Bhandavya345/Employee/model"
	"gorm.io/gorm"
)

type EmployeeRepository struct{}

func (r *EmployeeRepository) GetEmployees() ([]models.Employee, error) {

	var employees []models.Employee

	err := database.DB.Find(&employees).Error
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func AddEmployee(emp models.Employee) error {

	return database.DB.Create(&emp).Error
}

func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := database.DB.Find(&employees).Error
	return employees, err
}

func GetEmployeeByID(id int) (*models.Employee, error) {
	var employee models.Employee

	err := database.DB.Where("employee_id = ?", id).First(&employee).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("employee not found")
		}
		return nil, err
	}

	return &employee, nil
}

func UpdateSalary(id int, salary float64) error {
	result := database.DB.Model(&models.Employee{}).
		Where("employee_id = ?", id).
		Update("salary", salary)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}

	return nil
}

func DeleteEmployee(id int) error {
	result := database.DB.Where("employee_id = ?", id).
		Delete(&models.Employee{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("employee not found")
	}

	return nil
}

func GetAverageSalary() float64 {
	var employees []models.Employee
	database.DB.Find(&employees)

	if len(employees) == 0 {
		return 0
	}

	total := 0.0

	for _, emp := range employees {
		total += emp.Salary
	}

	return total / float64(len(employees))
}
