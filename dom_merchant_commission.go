package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
	"time"
)

type (
	MerchantCommission interface {
		Id() int
		MerchantId() int
		CommissionClassId() int
		ValidFrom() time.Time
		SetId(id int) error
		SetMerchantId(merchantID int) error
		SetCommissionClassId(commissionClassID int) error
		SetValidFrom(validFrom time.Time) error
		IsChanged() bool
		Entity() *EntityMerchantCommission
	}
	merchantCommission struct {
		instance *EntityMerchantCommission
		changed  map[string]bool
	}
)

func NewMerchantCommission(instance *EntityMerchantCommission) MerchantCommission {
	if instance == nil {
		return &merchantCommission{
			instance: &EntityMerchantCommission{},
			changed:  make(map[string]bool),
		}
	}
	return &merchantCommission{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (m *merchantCommission) Entity() *EntityMerchantCommission {
	return m.instance
}

func (m *merchantCommission) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityMerchantCommission](m.instance, fieldName, DbTagName, m.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (m *merchantCommission) IsChanged() bool {
	return len(m.changed) > 0
}

func (m *merchantCommission) Id() int {
	return m.instance.ID
}

func (m *merchantCommission) MerchantId() int {
	return m.instance.MerchantID
}

func (m *merchantCommission) CommissionClassId() int {
	return m.instance.CommissionClassID
}

func (m *merchantCommission) ValidFrom() time.Time {
	return m.instance.ValidFrom
}

func (m *merchantCommission) SetId(id int) error {
	m.instance.ID = id
	return m.setChanged("ID")
}

func (m *merchantCommission) SetMerchantId(merchantID int) error {
	m.instance.MerchantID = merchantID
	return m.setChanged("MerchantID")
}

func (m *merchantCommission) SetCommissionClassId(commissionClassID int) error {
	m.instance.CommissionClassID = commissionClassID
	return m.setChanged("CommissionClassID")
}

func (m *merchantCommission) SetValidFrom(validFrom time.Time) error {
	m.instance.ValidFrom = validFrom
	return m.setChanged("ValidFrom")
}
