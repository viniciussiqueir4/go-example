package controllers

import (
	"kyc/dtos"
	"kyc/repositories/postgres"
	"kyc/services"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var dto dtos.ProcessDTO

	err := c.ShouldBindJSON(&dto)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	repo := &postgres.ProcessRepository{}
	service := services.NewProcessService(repo)

	result, err := service.CreateProcess(dto)

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, result)
}

func GetProductById(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}
