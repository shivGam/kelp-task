package handlers

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"github.com/shivGam/kelp-task/models"
)

func GetSales(c *gin.Context){
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
	var sale models.Sale
	saleData,err:=sale.GetSaleData(companyIdInt)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Error getting sale data: "+ err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"saleData":saleData})
}