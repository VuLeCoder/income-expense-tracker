package routers

import (
	"net/http"
	"server/controllers"
	"server/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {

	noAuth := router.Group("/api")
	{
		noAuth.POST("/register", controllers.Register)
		noAuth.POST("/login", controllers.Login)
	}

	auth := router.Group("/api")
	{
		auth.GET("/profiles", middlewares.CheckLogin(), controllers.GetUserProfile)

		income := auth.Group("/income")
		{
			income.GET("", middlewares.CheckLogin(), controllers.GetIncomes)
			income.POST("", middlewares.CheckLogin(), controllers.AddIncome)
			income.DELETE("/:id", middlewares.CheckLogin(), middlewares.CheckOwnIncome(), controllers.DeleteIncome)
			income.PUT("/:id", middlewares.CheckLogin(), middlewares.CheckOwnIncome(), controllers.EditIncome)
		}

		expense := auth.Group("/expense")
		{			
			expense.GET("", middlewares.CheckLogin(), controllers.GetExpenses)
			expense.POST("", middlewares.CheckLogin(), controllers.AddExpense)
			expense.DELETE("/:id", middlewares.CheckLogin(), middlewares.CheckOwnExpense(), controllers.DeleteExpense)
			expense.PUT("/:id", middlewares.CheckLogin(), middlewares.CheckOwnExpense(), controllers.EditExpense)
		}
	}

	router.GET("/vuz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
}