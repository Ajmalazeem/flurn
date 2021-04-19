package api

type Loan interface {
	Create()
	List()
	Get()
	Approve()
	Reject()
	Cancel()
}

type loan struct{}

func (l *loan) Create()  {}
func (l *loan) List()    {}
func (l *loan) Get()     {}
func (l *loan) Approve() {}
func (l *loan) Reject()  {}
func (l *loan) Cancel()  {}

func NewLoanService() Loan {
	return &loan{}
}
