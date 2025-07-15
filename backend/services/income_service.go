package services

import (
	"errors"
	"server/databases"
	"server/dtos"
	"server/models"
)

// Lấy danh sách income theo user
func GetIncomesByUserId(userId int) ([]models.Income, error) {
	incomes := []models.Income{}
	if err := databases.DB.Where("user_id = ?", userId).Find(&incomes).Error; err != nil {
		return nil, errors.New("Không thể lấy danh sách income")
	}
	return incomes, nil
}

// Thêm income
func AddIncome(incomeReq dtos.IncomeRequest) (*models.Income, error) {
	income := models.Income{
		Date:        incomeReq.Date,
		Description: incomeReq.Description,
		Amount:      incomeReq.Amount,
		UserId:      incomeReq.UserID,
		UpdateAt:    incomeReq.UpdateAt,
	}

	if err := databases.DB.Create(&income).Error; err != nil {
		return nil, errors.New("Không thể tạo income")
	}

	return &income, nil
}

// Xoá income
func DeleteIncome(incomeId int) error {
	income := models.Income{}
	if err := databases.DB.Where("id = ?", incomeId).Delete(&income).Error; err != nil {
		return errors.New("Không thể xoá income")
	}
	return nil
}

// Chỉnh sửa income
func EditIncome(incomeId int, incomeReq dtos.IncomeRequest) (*models.Income, error) {
	income := models.Income{}
	if err := databases.DB.Where("id = ?", incomeId).First(&income).Error; err != nil {
		return nil, errors.New("Không tìm thấy income")
	}

	income.Date = incomeReq.Date
	income.Description = incomeReq.Description
	income.Amount = incomeReq.Amount
	income.UpdateAt = incomeReq.UpdateAt

	if err := databases.DB.Save(&income).Error; err != nil {
		return nil, errors.New("Không thể cập nhật income")
	}

	return &income, nil
}
