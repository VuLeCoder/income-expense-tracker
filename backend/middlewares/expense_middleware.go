package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"server/databases"
	"server/models"
)

func CheckOwnExpense() gin.HandlerFunc {
	return func(c *gin.Context) {
		expenseIdStr := c.Param("id")
		expenseId, err := strconv.Atoi(expenseIdStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "fail to convert expense id to int",
			})
			return
		}

		user := c.MustGet("user").(*models.User)

		fmt.Println("UserId:", user.ID, "- ExpenseId:", expenseId)

		expense := models.Expense{}
		err = databases.DB.Where("id = ? AND user_id = ?", expenseId, user.ID).First(&expense).Error
		if err != nil || expense.ID != expenseId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized to access this expense",
			})
			return
		}

		c.Next()
	}
}
