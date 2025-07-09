package terminal

const (
	DbTagName = "db"
)

const (
	BStatusCreated    = 1
	BStatusActive     = 2
	BStatusInactive   = 3
	BStatusTerminated = 4
	BStatusPending    = 5
)

const (
	CommissionTypeFixedRate       = 1 // Fixed amount per transaction
	CommissionTypeFixedPercentage = 2 // Fixed percentage of transaction amount
	CommissionTypeWorkHour        = 3 // Fixed amount per hour worked
)

const (
	CommissionTypeOperationSales        = 1
	CommissionTypeOperationValidations  = 2 // Validation of transactions
	CommissionTypeOperationWorkingHours = 3 // Working hours related operations
	CommissionTypeOperationMonthlySales = 4 // Monthly sales operations
	CommissionTypeOperationCashBalance  = 5 // Cash balance operations
)
