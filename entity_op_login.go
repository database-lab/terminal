package terminal

import (
	db "github.com/database-lab/dbutils"
	"time"
)

type EntityOpLogin struct {
	ID                    int       `db:"id"`
	OperatorID            int       `db:"operator_id"`
	Status                int       `db:"b_status"`
	CreatedDate           time.Time `db:"created_date"`
	ServiceID             int       `db:"service_id"`
	InternalServiceUserID string    `db:"internal_service_user_id"`
	ServiceData           db.JSONB  `db:"service_data"`
	LoginDetails          db.JSONB  `db:"login_details"`
}

func (o *EntityOpLogin) GetPassword() string {
	return o.ServiceData["secret"].(string)
}

func (o *EntityOpLogin) GetLocationID() int {
	return int(o.ServiceData["location_id"].(float64))
}

func (o *EntityOpLogin) GetMerchantID() int {
	return int(o.ServiceData["merchant_id"].(float64))
}

func (o *EntityOpLogin) GetType() int {
	return int(o.ServiceData["type"].(float64))
}
