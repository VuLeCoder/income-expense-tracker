package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"server/dtos"
	"server/models"
	"server/services"
)

func GetIncomes(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	incomes, err := services.GetIncomesByUserId(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get incomes successfully",
		"data": gin.H{
			"incomes": incomes,
		},
	})
}

func AddIncome(c *gin.Context) {
	var incomeReq dtos.IncomeRequest
	if err := c.ShouldBindJSON(&incomeReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*models.User)
	incomeReq.UserID = user.ID
	fmt.Println("User ID:", incomeReq.UserID)

	income, err := services.AddIncome(incomeReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Income added successfully",
		"data": gin.H{
			"income": income,
		},
	})
}

func DeleteIncome(c *gin.Context) {
	incomeIdStr := c.Param("id")
	incomeId, err := strconv.Atoi(incomeIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid income ID"})
		return
	}

	err = services.DeleteIncome(incomeId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Income deleted successfully",
		"data": gin.H{
			"incomeId": incomeId,
		},
	})
}

func EditIncome(c *gin.Context) {
	incomeIdStr := c.Param("id")
	incomeId, err := strconv.Atoi(incomeIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid income ID"})
		return
	}

	var incomeReq dtos.IncomeRequest
	if err := c.ShouldBindJSON(&incomeReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	income, err := services.EditIncome(incomeId, incomeReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Income edited successfully",
		"data": gin.H{
			"income": income,
		},
	})
}
