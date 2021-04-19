package models

type LoanRequest struct {
	CustomerName string  `json:"customer_name,omitempty" gorm:"column:customer_name"`
	PhoneNumber  string  `json:"phone_number,omitempty" gorm:"column:phone_number"`
	EmailID      string  `json:"email_id,omitempty" gorm:"column:email_id"`
	LoanAmount   float64 `json:"loan_amount,omitempty" gorm:"column:loan_amount"`
	Status       string  `json:"status,omitempty" gorm:"column:status"`
	CreditScore  float64 `json:"credit_score,omitempty" gorm:"column:credit_score"`
}

type ListRequest struct {
	Status            []string
	LoanAmountGreater float64
}
