package handlers

import(
	"github.com/gin-gonic/gin"
	"github.com/shivGam/kelp-task/models"
	"strconv"
	"net/http"
)

func GetEmployees(c *gin.Context){
	companyId:=c.Query("companyId")
	if companyId==""{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid company ID"})
		return
	}
	companyIdInt,err:=strconv.ParseInt(companyId,10,64)
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Invalid company ID"})
		return
	}
	var employee models.Employee
	employeeData,err:=employee.GetEmployeeData(companyIdInt)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Error getting employee data: "+ err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"employeeData":employeeData})
}