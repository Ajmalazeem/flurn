package store

import (
	"github.com/Ajmalazeem/flurn/models"
	"gorm.io/gorm"
)

type LoanStore interface {
	Create(models.LoanRequest) error
	List() ([]models.LoanRequest, error)
	Get()
}

type loanStore struct {
	db *gorm.DB
}

func (ls *loanStore) Create(req models.LoanRequest) error {
	return ls.db.Table("loan").Create(&req).Error
}

func (ls *loanStore) List() ([]models.LoanRequest, error) {
	var result []models.LoanRequest

	if err := ls.db.Table("loan").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
func (ls *loanStore) Get() {}

func NewLoanStore(db *gorm.DB) LoanStore {
	return &loanStore{
		db: db,
	}
}
