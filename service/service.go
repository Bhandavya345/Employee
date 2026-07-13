package services

import (
	"sync"

	models "github.com/Bhandavya345/Employee/model"
	"github.com/Bhandavya345/Employee/repository"
)

var Employees []models.Employee

type ReportService struct {
	Repo repository.EmployeeRepository
}

func AddEmployee(emp *models.Employee) error {

	return repository.AddEmployee(*emp)
}

func GetEmployees() ([]models.Employee, error) {
	employees, err := repository.GetAllEmployees()
	if err != nil {
		return nil, err
	}

	return employees, nil
}

func GetEmployeeByID(id int) (*models.Employee, error) {
	return repository.GetEmployeeByID(id)
}

func UpdateSalary(id int, salary float64) error {

	return repository.UpdateSalary(id, salary)
}

func DeleteEmployee(id int) error {
	return repository.DeleteEmployee(id)
}

func AverageSalary() float64 {
	return repository.GetAverageSalary()
}

func (s *ReportService) GenerateReport() (*models.ReportResponse, error) {

	employees, err := s.Repo.GetEmployees()
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup

	countChannel := make(chan map[string]int)
	salaryChannel := make(chan map[string]float64)

	wg.Add(2)

	go func() {
		defer wg.Done()
		countChannel <- departmentCount(employees)
	}()

	go func() {
		defer wg.Done()
		salaryChannel <- averageSalary(employees)
	}()

	go func() {
		wg.Wait()
		close(countChannel)
		close(salaryChannel)
	}()

	report := &models.ReportResponse{
		DepartmentCount: <-countChannel,
		AverageSalary:   <-salaryChannel,
	}

	return report, nil
}

func departmentCount(employees []models.Employee) map[string]int {

	result := make(map[string]int)

	for _, emp := range employees {
		result[emp.Department]++
	}

	return result
}

func averageSalary(employees []models.Employee) map[string]float64 {

	total := make(map[string]float64)
	count := make(map[string]int)

	for _, emp := range employees {

		total[emp.Department] += emp.Salary
		count[emp.Department]++
	}

	average := make(map[string]float64)

	for dept := range total {
		average[dept] = total[dept] / float64(count[dept])
	}

	return average
}
