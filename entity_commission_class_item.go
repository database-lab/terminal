package terminal

type EntityCommissionClassItem struct {
	ID                int `db:"id" goqu:"skipinsert,skipupdate"`
	CommissionID      int `db:"commission_id"`
	CommissionClassID int `db:"commission_class_id"`
	OperationType     int `db:"operation_type"`
	GameID            int `db:"game_id"`
}
