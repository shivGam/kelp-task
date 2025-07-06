package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"github.com/shivGam/kelp-task/db"
)

type Financial struct{
	CompanyId int64 `json:"companyId"`
	RevenueCurrentYear float64 `json:"revenueCurrentYear"`
	Profit float64 `json:"profit"`
	TaxRate float64 `json:"taxRate"`
	CreditRating string `json:"creditRating"`
}


func (f *Financial) calculateFinancialData(companyId int64) error {

	company, err := GetCompanyDetailsById(companyId)
	if err != nil {
		return err
	}
	f.CompanyId = company.CompanyId
	f.RevenueCurrentYear = company.RevenueBase * 0.8
	f.Profit = company.RevenueBase * 0.2
	f.TaxRate = company.RevenueBase * 0.05
	f.CreditRating = "A"
	err = insertFinancialData(f)
	if err != nil {
		return err
	}
	return nil
}

func insertFinancialData(f *Financial) error {
	time.Sleep(time.Duration(3) * time.Second)
	insertFinancialDataQuery := `
	INSERT INTO financials (company_id,revenue_current_year,profit,tax_rate,credit_rating)
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(insertFinancialDataQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(f.CompanyId, f.RevenueCurrentYear, f.Profit, f.TaxRate, f.CreditRating)
	if err != nil {
		return err
	}
	return nil
}

func (f *Financial) GetFinancialData(companyId int64) (Financial,error) {
	selectFinancialQuery := `SELECT * FROM financials WHERE company_id=?`
	row := db.DB.QueryRow(selectFinancialQuery, companyId)
	err := row.Scan(&f.CompanyId, &f.RevenueCurrentYear, &f.Profit, &f.TaxRate, &f.CreditRating)
	if err != nil {
		if err == sql.ErrNoRows {
			val ,err,  shared:= group.Do("financial"+strconv.FormatInt(companyId, 10), func() (interface{}, error) {
				newFinancial := &Financial{}
				err = newFinancial.calculateFinancialData(companyId)
				if err != nil {
					return nil, err
				}
				return *newFinancial, nil
			})
			if err != nil {
				return Financial{}, err
			}
			fmt.Printf("Financial companyid %d: Result: %v, Shared: %t\n", companyId, val, shared)
			return val.(Financial), nil
		}
		return Financial{}, err
	}
	return *f, nil
}
