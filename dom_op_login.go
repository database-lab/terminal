package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
	"time"
)

type (
	OpLogin interface {
	}
	opLogin struct {
		instance *EntityOpLogin
		changed  map[string]bool
	}
)

func NewOpLogin(instance *EntityOpLogin) OpLogin {
	if instance == nil {
		return &opLogin{
			instance: &EntityOpLogin{},
			changed:  make(map[string]bool),
		}
	}
	return &opLogin{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (o *opLogin) Entity() *EntityOpLogin {
	return o.instance
}

func (o *opLogin) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityOpLogin](o.instance, fieldName, DbTagName, o.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (o *opLogin) IsChanged() bool {
	return len(o.changed) > 0
}

func (o *opLogin) Id() int {
	return o.instance.ID
}

func (o *opLogin) OperatorId() int {
	return o.instance.OperatorID
}

func (o *opLogin) Status() int {
	return o.instance.Status
}

func (o *opLogin) CreatedDate() time.Time {
	return o.instance.CreatedDate
}

func (o *opLogin) ServiceId() int {
	return o.instance.ServiceID
}

func (o *opLogin) InternalServiceUserId() string {
	return o.instance.InternalServiceUserID
}

func (o *opLogin) ServiceData() db.JSONB {
	if o.instance.ServiceData == nil {
		return make(db.JSONB)
	}
	return o.instance.ServiceData
}

func (o *opLogin) LoginDetails() db.JSONB {
	if o.instance.LoginDetails == nil {
		return make(db.JSONB)
	}
	return o.instance.LoginDetails
}

func (o *opLogin) SetId(id int) error {
	if o.instance.ID == id {
		return nil
	}
	o.instance.ID = id
	return o.setChanged("ID")
}

func (o *opLogin) SetOperatorId(operatorID int) error {
	if o.instance.OperatorID == operatorID {
		return nil // No change needed
	}
	o.instance.OperatorID = operatorID
	return o.setChanged("OperatorID")
}

func (o *opLogin) SetStatus(status int) error {
	if o.instance.Status == status {
		return nil // No change needed
	}
	o.instance.Status = status
	return o.setChanged("Status")
}

func (o *opLogin) SetCreatedDate(createdDate time.Time) error {
	if o.instance.CreatedDate.Equal(createdDate) {
		return nil // No change needed
	}
	o.instance.CreatedDate = createdDate
	return o.setChanged("CreatedDate")
}

func (o *opLogin) SetServiceId(serviceID int) error {
	if o.instance.ServiceID == serviceID {
		return nil // No change needed
	}
	o.instance.ServiceID = serviceID
	return o.setChanged("ServiceID")
}

func (o *opLogin) SetInternalServiceUserId(internalServiceUserID string) error {
	if o.instance.InternalServiceUserID == internalServiceUserID {
		return nil // No change needed
	}
	o.instance.InternalServiceUserID = internalServiceUserID
	return o.setChanged("InternalServiceUserID")
}

func (o *opLogin) SetServiceData(serviceData db.JSONB) error {
	if serviceData == nil {
		return nil // No change needed
	}
	if db.EqualJSONB(o.instance.ServiceData, serviceData) {
		return nil // No change needed
	}
	o.instance.ServiceData = serviceData
	return o.setChanged("ServiceData")
}
