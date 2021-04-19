package api

import (
	"github.com/Ajmalazeem/flurn/models"
	"github.com/Ajmalazeem/flurn/store"
)

type Loan interface {
	Create(models.LoanRequest)
	List(models.ListRequest) ([]models.LoanRequest, error)
	Get(models.GetRequest) (*models.LoanRequest, error)
	Approve(models.ApproveRequest) (*models.LoanRequest, error)
	Cancel(models.CancelRequest) (*models.LoanRequest,error)
}

type loan struct {
	loanStore store.LoanStore
}

func (l *loan) Create(req models.LoanRequest) {
	l.loanStore.Create(req)
}
func (l *loan) List(req models.ListRequest) ([]models.LoanRequest, error) {
	return l.loanStore.List(req)
}
func (l *loan) Get(req models.GetRequest) (*models.LoanRequest, error) {
	return l.loanStore.Get(req)
}
func (l *loan) Approve(req models.ApproveRequest) (*models.LoanRequest, error) {
	return l.loanStore.Approve(req)
}
func (l *loan) Cancel(req models.CancelRequest) (*models.LoanRequest, error) {
	return l.loanStore.Cancel(req)
}


func NewLoanService(loanStore store.LoanStore) Loan {
	return &loan {
		loanStore: loanStore,
	}
}
