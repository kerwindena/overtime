package overtime

type Overtime struct {
	Date           *Date
	Overtime       int
	Reason         string
	TimeOfApproval Time
	Employee       *Employee
}
