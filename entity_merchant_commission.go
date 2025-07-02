package terminal

import "time"

type EntityMerchantCommission struct {
	ID                int       `db:"id" goqu:"skipinsert,skipupdate"`
	CommissionClassID int       `db:"commission_class_id"`
	MerchantID        int       `db:"merchant_id"`
	ValidFrom         time.Time `db:"valid_from"`
}
