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

func GetExpenses(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	expenses, err := services.GetExpensesByUserId(user.ID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Get expenses successfully",
		"data": gin.H{
			"expenses": expenses,
		},
	})
}

func AddExpense(c *gin.Context) {
	var expenseReq dtos.ExpenseRequest
	if err := c.ShouldBindJSON(&expenseReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.MustGet("user").(*models.User)
	expenseReq.UserID = user.ID
	fmt.Println("User ID:", expenseReq.UserID)

	expense, err := services.AddExpense(expenseReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Expense added successfully",
		"data": gin.H{
			"expense": expense,
		},
	})
}

func DeleteExpense(c *gin.Context) {
	expenseIdStr := c.Param("id")
	expenseId, err := strconv.Atoi(expenseIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	err = services.DeleteExpense(expenseId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Expense deleted successfully",
		"data": gin.H{
			"expenseId": expenseId,
		},
	})
}

func EditExpense(c *gin.Context) {
	expenseIdStr := c.Param("id")
	expenseId, err := strconv.Atoi(expenseIdStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid expense ID"})
		return
	}

	var expenseReq dtos.ExpenseRequest
	if err := c.ShouldBindJSON(&expenseReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	expense, err := services.EditExpense(expenseId, expenseReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Expense edited successfully",
		"data": gin.H{
			"expense": expense,
		},
	})
}
