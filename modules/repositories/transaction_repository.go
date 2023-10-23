package repositories

import (
	"financial-management/modules/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindAll() ([]models.Transaction, error)
	FindByID(id int) (*models.Transaction, error)
	Create(transaction *models.Transaction) (*models.Transaction, error)
	Update(transaction *models.Transaction) (*models.Transaction, error)
	Delete(id int) error
}

type TransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &TransactionRepositoryImpl{db}
}

func (tr *TransactionRepositoryImpl) FindAll() ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := tr.db.Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}

func (tr *TransactionRepositoryImpl) FindByID(id int) (*models.Transaction, error) {
	var transaction models.Transaction
	result := tr.db.First(&transaction, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &transaction, nil
}

func (tr *TransactionRepositoryImpl) Create(transaction *models.Transaction) (*models.Transaction, error) {
	result := tr.db.Create(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

func (tr *TransactionRepositoryImpl) Update(transaction *models.Transaction) (*models.Transaction, error) {
	result := tr.db.Save(transaction)
	if result.Error != nil {
		return nil, result.Error
	}
	return transaction, nil
}

func (tr *TransactionRepositoryImpl) Delete(id int) error {
	result := tr.db.Delete(&models.Transaction{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
