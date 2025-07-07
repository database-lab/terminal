package terminal

type EntityCommission struct {
	ID          int     `db:"id" goqu:"skipinsert,skipupdate"`
	Type        int     `db:"type"`
	Description string  `db:"description"`
	Value       float64 `db:"value"`
}
