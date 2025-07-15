package services

import (
	"errors"
	"server/databases"
	"server/dtos"
	"server/models"
)

// Lấy danh sách expense theo user
func GetExpensesByUserId(userId int) ([]models.Expense, error) {
	expenses := []models.Expense{}
	if err := databases.DB.Where("user_id = ?", userId).Find(&expenses).Error; err != nil {
		return nil, errors.New("Không thể lấy danh sách expense")
	}
	return expenses, nil
}

// Thêm expense
func AddExpense(expenseReq dtos.ExpenseRequest) (*models.Expense, error) {
	expense := models.Expense{
		Date:		 expenseReq.Date,
		Description: expenseReq.Description,
		Amount:      expenseReq.Amount,
		UserId:      expenseReq.UserID,
		UpdateAt:    expenseReq.UpdateAt,
	}

	if err := databases.DB.Create(&expense).Error; err != nil {
		return nil, errors.New("Không thể tạo expense")
	}

	return &expense, nil
}

// Xoá expense
func DeleteExpense(expenseId int) error {
	expense := models.Expense{}
	if err := databases.DB.Where("id = ?", expenseId).Delete(&expense).Error; err != nil {
		return errors.New("Không thể xoá expense")
	}
	return nil
}

// Chỉnh sửa expense
func EditExpense(expenseId int, expenseReq dtos.ExpenseRequest) (*models.Expense, error) {
	expense := models.Expense{}
	if err := databases.DB.Where("id = ?", expenseId).First(&expense).Error; err != nil {
		return nil, errors.New("Không tìm thấy expense")
	}

	expense.Date = expenseReq.Date
	expense.Description = expenseReq.Description
	expense.Amount = expenseReq.Amount
	expense.UpdateAt = expenseReq.UpdateAt

	if err := databases.DB.Save(&expense).Error; err != nil {
		return nil, errors.New("Không thể cập nhật expense")
	}

	return &expense, nil
}