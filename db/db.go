package db

import (
	"database/sql"
	"fmt"
	"math"
	"math/rand"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB(count int) {
	var err error
	DB, err = sql.Open("sqlite", "companies.db")
	if err != nil {
		panic("Error opening database:" + err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)
	createTable()
	err = insertRandomData(count)
	if err != nil {
		panic("Error inserting random data:" + err.Error())
	}
}

func createTable() {

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS companies(
	company_id INTEGER PRIMARY KEY AUTOINCREMENT,
	founding_year INTEGER NOT NULL,
	employee_count INTEGER NOT NULL,
	country_code TEXT NOT NULL,
	revenue_base REAL NOT NULL,
	operating_cost_base REAL NOT NULL
	)`

	result, err := DB.Exec(createTableQuery)
	if err != nil {
		panic("Error creating table:" + err.Error())
	}
	fmt.Printf("%v", result)

	createFinancialsTableQuery := `
	CREATE TABLE IF NOT EXISTS financials(
	company_id INTEGER PRIMARY KEY REFERENCES companies(company_id),
	revenue_current_year REAL NOT NULL,
	profit REAL NOT NULL,
	tax_rate REAL NOT NULL,
	credit_rating TEXT NOT NULL
	)`

	result, err = DB.Exec(createFinancialsTableQuery)
	if err != nil {
		panic("Error creating financials table:" + err.Error())
	}
	fmt.Printf("%v", result)

	createSalesTableQuery := `
	CREATE TABLE IF NOT EXISTS sales(
	company_id INTEGER PRIMARY KEY REFERENCES companies(company_id),
	total_units_sold INTEGER NOT NULL,
	avg_sale_price REAL NOT NULL,
	top_region TEXT NOT NULL
	)`

	result, err = DB.Exec(createSalesTableQuery)
	if err != nil {
		panic("Error creating sales table:" + err.Error())
	}
	fmt.Printf("%v", result)

	createEmployeesTableQuery := `
	CREATE TABLE IF NOT EXISTS employees(
	company_id INTEGER PRIMARY KEY REFERENCES companies(company_id),
	engineers INTEGER NOT NULL,
	managers INTEGER NOT NULL,
	attrition_rate REAL NOT NULL,
	avg_tenure REAL NOT NULL
	)`

	result, err = DB.Exec(createEmployeesTableQuery)
	if err != nil {
		panic("Error creating employees table:" + err.Error())
	}
	fmt.Printf("%v", result)
}

func insertRandomData(count int) error {
	insertCompanyDataQuery := `
	INSERT INTO companies(founding_year,employee_count,country_code,revenue_base,operating_cost_base)
	VALUES(?,?,?,?,?)
	`
	for range count {
		foundingYear := int64(2000 + rand.Intn(25))
		employeeCount := int64(100 + rand.Intn(500))
		countryCode := []string{"US", "IN", "DE", "SG"}[rand.Intn(4)]
		revenueBase := math.Round((1000000 * (1 + rand.Float64() * 10))*100)/100
		operatingCostBase := math.Round((1000000 * (0.5 + rand.Float64() * 5))*100)/100
		stmt, err := DB.Prepare(insertCompanyDataQuery)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(foundingYear, employeeCount, countryCode, revenueBase, operatingCostBase)
		stmt.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
