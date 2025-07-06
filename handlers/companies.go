package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/shivGam/kelp-task/db"
	"github.com/shivGam/kelp-task/models"
)

func GetCompanies(c *gin.Context){
	var company models.Company
	rows,err:=db.DB.Query("SELECT * FROM companies")
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Error fetching companies"})
		return
	}
	defer rows.Close()
	var companies []models.Company
	for rows.Next(){
		err=rows.Scan(&company.CompanyId,&company.FoundingYear,&company.EmployeeCount,&company.CountryCode,&company.RevenueBase,&company.OperatingCostBase)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{"error":"Error fetching companies"})
			return
		}
		companies=append(companies,company)
	}
	c.JSON(http.StatusOK,gin.H{"companies":companies})
}
