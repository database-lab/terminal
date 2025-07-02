package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
	"time"
)

type (
	Operator interface {
		Id() int
		Type() int
		Status() int
		CreationDate() time.Time
		LocationID() int
		Suspended() bool
		FinancialClassId() int
		Deleted() bool
		LoginAttempt() *int
		SetId(id int) error
		SetType(operatorType int) error
		SetStatus(status int) error
		SetCreationDate(creationDate time.Time) error
		SetLocationID(locationID int) error
		SetSuspended(suspended bool) error
		SetFinancialClassId(financialClassId int) error
		SetDeleted(deleted bool) error
		SetLoginAttempt(loginAttempt int) error
		IsChanged() bool
		ToMap() map[string]interface{}
		Entity() *EntityOperator
	}
	operator struct {
		instance *EntityOperator
		changed  map[string]bool
	}
)

func NewOperator(instance *EntityOperator) Operator {
	if instance == nil {
		return &operator{
			instance: &EntityOperator{},
			changed:  make(map[string]bool),
		}
	}
	return &operator{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (o *operator) Entity() *EntityOperator {
	return o.instance
}

func (o *operator) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityOperator](o.instance, fieldName, DbTagName, o.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (o *operator) IsChanged() bool {
	return len(o.changed) > 0
}

func (o *operator) Id() int {
	return o.instance.ID
}

func (o *operator) Type() int {
	return o.instance.OperatorType
}

func (o *operator) Status() int {
	return o.instance.Status
}

func (o *operator) CreationDate() time.Time {
	return o.instance.CreationDate
}

func (o *operator) LocationID() int {
	return o.instance.LocationID
}

func (o *operator) Suspended() bool {
	return o.instance.Suspended
}

func (o *operator) FinancialClassId() int {
	return o.instance.FinancialClassId
}

func (o *operator) Deleted() bool {
	return o.instance.Deleted
}

func (o *operator) LoginAttempt() *int {
	return o.instance.LoginAttempt
}

func (o *operator) SetId(id int) error {
	o.instance.ID = id
	return o.setChanged("ID")
}

func (o *operator) SetType(operatorType int) error {
	if o.instance.OperatorType == operatorType {
		return nil // No change needed
	}
	o.instance.OperatorType = operatorType
	return o.setChanged("OperatorType")
}

func (o *operator) SetStatus(status int) error {
	if o.instance.Status == status {
		return nil // No change needed
	}
	o.instance.Status = status
	return o.setChanged("Status")
}

func (o *operator) SetCreationDate(creationDate time.Time) error {
	if o.instance.CreationDate.Equal(creationDate) {
		return nil // No change needed
	}
	o.instance.CreationDate = creationDate
	return o.setChanged("CreationDate")
}

func (o *operator) SetLocationID(locationID int) error {
	if o.instance.LocationID == locationID {
		return nil // No change needed
	}
	o.instance.LocationID = locationID
	return o.setChanged("LocationID")
}

func (o *operator) SetSuspended(suspended bool) error {
	if o.instance.Suspended == suspended {
		return nil // No change needed
	}
	o.instance.Suspended = suspended
	return o.setChanged("Suspended")
}

func (o *operator) SetFinancialClassId(financialClassId int) error {
	if o.instance.FinancialClassId == financialClassId {
		return nil // No change needed
	}
	o.instance.FinancialClassId = financialClassId
	return o.setChanged("FinancialClassId")
}

func (o *operator) SetDeleted(deleted bool) error {
	if o.instance.Deleted == deleted {
		return nil // No change needed
	}
	o.instance.Deleted = deleted
	return o.setChanged("Deleted")
}

func (o *operator) SetLoginAttempt(loginAttempt int) error {
	if *o.instance.LoginAttempt == loginAttempt {
		return nil // No change needed
	}
	o.instance.LoginAttempt = &loginAttempt
	return o.setChanged("LoginAttempt")
}

func (o *operator) ToMap() map[string]interface{} {
	return db.MapValuesFromChangedFields[*EntityOperator](o.instance, o.changed, DbTagName)
}
