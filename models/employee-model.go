package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"github.com/shivGam/kelp-task/db"
	"golang.org/x/sync/singleflight"
)


type Employee struct {
	CompanyId     int64   `json:"companyId"`
	Engineers     int64   `json:"engineers"`
	Managers      int64   `json:"managers"`
	AttritionRate float64 `json:"attritionRate"`
	AvgTenure     float64 `json:"avgTenure"`
}

var group = singleflight.Group{}

func (e *Employee) calculateEmployeeData(companyId int64) error {

	company, err := GetCompanyDetailsById(companyId)
	if err != nil {
		return err
	}
	e.CompanyId = company.CompanyId
	e.Engineers = int64(float64(company.EmployeeCount) + 0.1*float64(company.EmployeeCount))
	e.Managers = int64(float64(company.EmployeeCount) + 0.05*float64(company.EmployeeCount))
	e.AttritionRate = float64(company.EmployeeCount) / float64(e.Engineers)
	e.AvgTenure = float64(company.EmployeeCount) / float64(e.Engineers)
	err = insertEmployeeData(e)
	if err != nil {
		return err
	}
	return nil
}

func insertEmployeeData(e *Employee) error {
	time.Sleep(time.Duration(3) * time.Second)
	insertEmployeeDataQuery := `
	INSERT INTO employees (company_id,engineers,managers,attrition_rate,avg_tenure)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(insertEmployeeDataQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.CompanyId, e.Engineers, e.Managers, e.AttritionRate, e.AvgTenure)
	if err != nil {
		return err
	}
	return nil
}

func (e *Employee) GetEmployeeData(companyId int64) (Employee,error) {
	selectEmployeeQuery := `SELECT * FROM employees WHERE company_id=?`
	row := db.DB.QueryRow(selectEmployeeQuery, companyId)
	err := row.Scan(&e.CompanyId, &e.Engineers, &e.Managers, &e.AttritionRate, &e.AvgTenure)
	if err != nil {
		if err == sql.ErrNoRows {
			val ,err,  shared:= group.Do("employee"+strconv.FormatInt(companyId, 10), func() (interface{}, error) {
				newEmployee := &Employee{}
				err = newEmployee.calculateEmployeeData(companyId)
				if err != nil {
					return nil, err
				}
				return *newEmployee, nil
			})
			if err != nil {
				return Employee{}, err
			}
			fmt.Printf("Employee companyid %d: Result: %v, Shared: %t\n", companyId, val, shared)
			return val.(Employee), nil
		}
		return Employee{}, err
	}
	return *e, nil
}
