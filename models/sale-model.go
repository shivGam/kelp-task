package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"github.com/shivGam/kelp-task/db"
)

type Sale struct{
	CompanyId int64 `json:"companyId"`
	TotalUnitsSold int64 `json:"totalUnitsSold"`
	AvgSalePrice float64 `json:"avgSalePrice"`
	TopRegion string `json:"topRegion"`
}


func (s *Sale) calculateSaleData(companyId int64) error {

	company, err := GetCompanyDetailsById(companyId)
	if err != nil {
		return err
	}
	s.CompanyId = company.CompanyId
	s.TotalUnitsSold = int64(float64(company.EmployeeCount) + 0.1*float64(company.EmployeeCount))
	s.AvgSalePrice = float64(company.RevenueBase) / float64(s.TotalUnitsSold)
	s.TopRegion = company.CountryCode
	err = insertSaleData(s)
	if err != nil {
		return err
	}
	return nil
}

func insertSaleData(s *Sale) error {
	time.Sleep(time.Duration(3) * time.Second)
	insertSaleDataQuery := `
	INSERT INTO sales (company_id,total_units_sold,avg_sale_price,top_region)
	VALUES (?,?,?,?)
	`
	stmt, err := db.DB.Prepare(insertSaleDataQuery)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.CompanyId, s.TotalUnitsSold, s.AvgSalePrice, s.TopRegion)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sale) GetSaleData(companyId int64) (Sale,error) {
	selectSaleQuery := `SELECT * FROM sales WHERE company_id=?`
	row := db.DB.QueryRow(selectSaleQuery, companyId)
	err := row.Scan(&s.CompanyId, &s.TotalUnitsSold, &s.AvgSalePrice, &s.TopRegion)
	if err != nil {
		if err == sql.ErrNoRows {
			val ,err,  shared:= group.Do("sale"+strconv.FormatInt(companyId, 10), func() (interface{}, error) {
				newSale := &Sale{}
				err = newSale.calculateSaleData(companyId)
				if err != nil {
					return nil, err
				}
				return *newSale, nil
			})
			if err != nil {
				return Sale{}, err
			}
			fmt.Printf("Sale companyid %d: Result: %v, Shared: %t\n", companyId, val, shared)
			return val.(Sale), nil
		}
		return Sale{}, err
	}
	return *s, nil
}
