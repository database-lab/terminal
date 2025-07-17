package terminal

type EntityOperatorPrivileges struct {
	ID          int `db:"id" goqu:"skipinsert,skipupdate"`
	OperatorID  int `db:"operator_id"`
	PrivilegeID int `db:"privilege_id"`
	GameID      int `db:"game_id"`
}
