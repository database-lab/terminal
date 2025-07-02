package terminal

type EntityAddress struct {
	ID          int    `db:"id" goqu:"skipinsert,skipupdate"`
	City        string `db:"city"`
	Street      string `db:"street"`
	Zip         string `db:"zip"`
	State       string `db:"state"`
	Region      string `db:"region"`
	Country     string `db:"country"`
	Description string `db:"description"`
}
