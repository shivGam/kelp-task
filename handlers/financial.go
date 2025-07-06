package handlers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/shivGam/kelp-task/models"
)

func GetFinancials(c *gin.Context){
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
	var financial models.Financial
	financialData,err:=financial.GetFinancialData(companyIdInt)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Error getting financial data: "+ err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"financialData":financialData})
	
}