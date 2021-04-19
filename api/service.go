package api

import (
	"github.com/Ajmalazeem/flurn/models"
	"github.com/Ajmalazeem/flurn/store"
)

type Loan interface {
	Create(models.LoanRequest)
	List(models.ListRequest) ([]models.LoanRequest, error)
	Get(models.GetRequest) (*models.LoanRequest, error)
	Approve()
	Reject()
	Cancel()
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
func (l *loan) Approve() {}
func (l *loan) Reject()  {}
func (l *loan) Cancel()  {}

func NewLoanService(loanStore store.LoanStore) Loan {
	return &loan {
		loanStore: loanStore,
	}
}
