package terminal

import (
	db "github.com/database-lab/dbutils"
)

type EntityOpMobile struct {
	ID              int      `db:"id"`
	OperatorID      int      `db:"operator_id"`
	Email           string   `db:"email"`
	Login           string   `db:"login"`
	Pin             *string  `db:"pin"`
	LangCode        *string  `db:"lang_code"`
	PushKey         *string  `db:"push_key"`
	ClientOS        *int     `db:"client_os"`
	FirstName       *string  `db:"first_name"`
	LastName        *string  `db:"last_name"`
	Phone           *string  `db:"phone"`
	OfflineSyncData db.JSONB `db:"offline_sync_data"`
}
