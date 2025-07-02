package terminal

import "time"

type EntityOperator struct {
	ID               int       `db:"id"`
	OperatorType     int       `db:"operator_type"`
	Status           int       `db:"b_status"`
	CreationDate     time.Time `db:"creation_date"`
	LocationID       int       `db:"location_id"`
	Suspended        bool      `db:"suspended"`
	FinancialClassId int       `db:"financial_class_id"`
	Deleted          bool      `db:"deleted"`
	LoginAttempt     *int      `db:"login_attempts"`
	//LastLogin        time.Time `db:"last_login"`
	//Password     string `db:"password"`
	//LastLoginAttempt *time.Time `db:"last_login_attempt"`
	//NextAllowedLogin *time.Time `db:"next_allowed_login"`
}
