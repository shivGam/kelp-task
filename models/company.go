package models

import (
	"github.com/shivGam/kelp-task/db"
	"database/sql"
	"errors"
)

type Company struct {
	CompanyId int64 `json:"companyId"`
	FoundingYear int64 `json:"foundingYear"`
	EmployeeCount int64 `json:"employeeCount"`
	CountryCode string `json:"countryCode"`
	RevenueBase float64 `json:"revenueBase"`
	OperatingCostBase float64 `json:"operatingCostBase"`
}


func GetCompanyDetailsById(companyId int64) (Company,error){
    selectQuery:= `SELECT * FROM companies WHERE company_id=?`
    row:= db.DB.QueryRow(selectQuery,companyId)
    var company Company
    err:=row.Scan(&company.CompanyId,&company.FoundingYear,&company.EmployeeCount,&company.CountryCode,&company.RevenueBase,&company.OperatingCostBase)
    if err!=nil{
        if err==sql.ErrNoRows{
            return Company{},errors.New("company not found")
        }
        return Company{},err
    }
    return company,nil
}


