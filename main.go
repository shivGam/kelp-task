package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shivGam/kelp-task/db"
	"github.com/shivGam/kelp-task/handlers"	
)

func main(){
	db.InitDB(3)
	server:=gin.Default()
	server.GET("/financials",handlers.GetFinancials)
	server.GET("/sales",handlers.GetSales)
	server.GET("/employees",handlers.GetEmployees)
	server.GET("/companies",handlers.GetCompanies)
	server.Run(":8080")
}
