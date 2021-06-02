package main

import (
	service "RestAPIForGolang/service"
	"github.com/gin-gonic/gin"
)


func init()  {
	router := gin.Default()
	router.GET("/deposit/:input",service.Deposit)
	router.GET("/withdraw/:input", service.Withdraw)
	router.GET("/balance/",service.GetBalance)
	router.Run(":80")
}

type Result struct {
	Amount  int    `json:"amount"`
	Status  string `json:"status"`
	Message string `json:"message"`
}



func main() {
}
