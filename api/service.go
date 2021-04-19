package api

import (
	"github.com/Ajmalazeem/flurn/models"
	"github.com/Ajmalazeem/flurn/store"
)

type Loan interface {
	Create(models.LoanRequest)
	List() ([]models.LoanRequest, error)
	Get() *models.LoanRequest
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
func (l *loan) List() ([]models.LoanRequest, error) {
	return l.loanStore.List()
}
func (l *loan) Get() *models.LoanRequest {
	return nil
}
func (l *loan) Approve() {}
func (l *loan) Reject()  {}
func (l *loan) Cancel()  {}

func NewLoanService(loanStore store.LoanStore) Loan {
	return &loan{
		loanStore: loanStore,
	}
}
