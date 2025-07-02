package terminal

type EntityMerchant struct {
	ID               int     `db:"id" goqu:"skipinsert,skipupdate"`
	LastName         *string `db:"last_name"`
	FirstName        *string `db:"first_name"`
	Title            *string `db:"title"`
	Status           int     `db:"b_status"`
	MerchantType     int     `db:"merchant_type"`
	CompanyName      string  `db:"company_name"`
	Deleted          bool    `db:"deleted"`
	LotteryPoint     bool    `db:"lottery_point"`
	ParentMerchantID *int    `db:"parent_merchant"`
}
