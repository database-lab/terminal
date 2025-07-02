package terminal

type EntityCommissionClass struct {
	ID          int    `db:"id" goqu:"skipinsert,skipupdate"`
	Description string `db:"description"`
}
