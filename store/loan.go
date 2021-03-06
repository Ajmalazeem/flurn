package store

import (
	"github.com/Ajmalazeem/flurn/models"
	"gorm.io/gorm"
)

type LoanStore interface {
	Create(models.LoanRequest) error
	List(models.ListRequest) ([]models.LoanRequest, error)
	Get(models.GetRequest) (*models.LoanRequest, error)
	Approve(models.ApproveRequest) (*models.LoanRequest, error)
	Cancel(models.CancelRequest) (*models.LoanRequest, error)
}

type loanStore struct {
	db *gorm.DB
}

func (ls *loanStore) Create(req models.LoanRequest) error {
	return ls.db.Table("loan").Create(&req).Error
}

func (ls *loanStore) List(req models.ListRequest) ([]models.LoanRequest, error) {
	var result []models.LoanRequest

	query := ls.db.Table("loan")

	if len(req.Status) > 0 {
		query = query.Where("status in ?", req.Status)
	}

	if req.LoanAmountGreater > 0 {
		query = query.Where("loan_amount > ?", req.LoanAmountGreater)
	}

	if err := query.Debug().Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
func (ls *loanStore) Get(req models.GetRequest) (*models.LoanRequest, error) {
	var result models.LoanRequest
     
	ls.db.Table("loan").Where("id = ?", req.Id).Find(&result)


   return &result, nil
}
func (ls *loanStore) Approve(req models.ApproveRequest) (*models.LoanRequest, error) {
	var result models.LoanRequest
     
	ls.db.Table("loan").Where("id = ? ", req.Id).Updates(map[string]interface{}{"status" :req.Status})

   return &result, nil
}
func (ls *loanStore) Cancel(req models.CancelRequest) (*models.LoanRequest, error) {
	var result models.LoanRequest
     
	ls.db.Table("loan").Where("id = ? ", req.Id).Updates(map[string]interface{}{"status" :req.Status})

   return &result, nil
}


func NewLoanStore(db *gorm.DB) LoanStore {
	return &loanStore{
		db: db,
	}
}
