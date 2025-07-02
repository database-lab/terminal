package terminal

import "time"

type EntityLocation struct {
	ID               int        `db:"id" goqu:"skipinsert,skipupdate"`
	OldID            *int       `db:"old_id"`
	MerchantID       int        `db:"merchant_id"`
	NewMerchantID    *int       `db:"new_merchant_id"`
	PrevMerchantID   *int       `db:"prev_merchant_id"`
	Name             string     `db:"name"`
	Status           int        `db:"b_status"`
	StatusReasonNote *string    `db:"status_reason_note"`
	StatusDate       *time.Time `db:"status_date"`
	OpeningHoursID   *int       `db:"opening_hours_id"`
	CollectorID      *int       `db:"collector_id"`
	AddressID        *int       `db:"address_id"`
	Deleted          bool       `db:"deleted"`
}
