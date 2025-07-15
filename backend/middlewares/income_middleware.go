package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"server/databases"
	"server/models"
)

func CheckOwnIncome() gin.HandlerFunc {
	return func(c *gin.Context) {
		incomeIdStr := c.Param("id")
		incomeId, err := strconv.Atoi(incomeIdStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"msg":  "fail to convert income id to int",
			})
			return
		}

		user := c.MustGet("user").(*models.User)

		fmt.Println("UserId:", user.ID, "- ExpenseId:", incomeId)

		income := models.Income{}
		err = databases.DB.Where("id = ? AND user_id = ?", incomeId, user.ID).First(&income).Error
		if err != nil || income.ID != incomeId {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized to access this income",
			})
			return
		}

		c.Next()
	}
}
