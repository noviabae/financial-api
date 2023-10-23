package services

import (
	"financial-management/modules/dto"
	"financial-management/modules/models"
	"financial-management/modules/repositories"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TransactionService interface {
	GetAll() ([]models.Transaction, error)
	GetByID(id int) (*models.Transaction, error)
	Create(ctx *gin.Context) (*models.Transaction, error)
	Update(ctx *gin.Context) (*models.Transaction, error)
	Delete(id int) error
}

type TransactionServiceImpl struct {
	transactionRepository repositories.TransactionRepository
}

// Create implements TransactionService.
func (ts *TransactionServiceImpl) Create(ctx *gin.Context) (*models.Transaction, error) {
	var input dto.CreateTransactionInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	transaction := models.Transaction{
		TransactionDate: input.TransactionDate,
		TransactionType: input.TransactionType,
		Amount:          input.Amount,
		CategoryID:      input.CategoryID,
		Description:     input.Description,
	}

	result, err := ts.transactionRepository.Create(&transaction)

	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete implements TransactionService.
func (ts *TransactionServiceImpl) Delete(id int) error {
	err := ts.transactionRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements TransactionService.
func (ts *TransactionServiceImpl) GetAll() ([]models.Transaction, error) {
	transaction, err := ts.transactionRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// GetByID implements TransactionService.
func (ts *TransactionServiceImpl) GetByID(id int) (*models.Transaction, error) {
	transaction, err := ts.transactionRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

// Update implements TransactionService.
func (ts *TransactionServiceImpl) Update(ctx *gin.Context) (*models.Transaction, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.UpdateTransactionInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()
	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	transaction := models.Transaction{
		TransactionID:   id,
		TransactionDate: input.TransactionDate,
		TransactionType: input.TransactionType,
		Amount:          input.Amount,
		CategoryID:      input.CategoryID,
		Description:     input.Description,
	}

	result, err := ts.transactionRepository.Update(&transaction)

	if err != nil {
		return nil, err
	}
	return result, nil

}

func NewTransactionService(transactionRepository repositories.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{transactionRepository}
}
