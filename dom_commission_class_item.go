package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

type (
	CommissionClassItem interface {
		Entity() *EntityCommissionClassItem
		SetCommission(commission Commission)
		Commission() Commission
		SetCommissionClass(commissionClass CommissionClass)
		CommissionClass() CommissionClass
		IsChanged() bool
		ToMap() map[string]any
		Id() int
		CommissionId() int
		CommissionClassId() int
		OperationType() int
		GameId() int
		SetId(id int) error
		SetCommissionId(commissionId int) error
		SetCommissionClassId(commissionClassId int) error
		SetOperationType(operationType int) error
		SetGameId(gameId int) error
		String() string
	}
	commissionClassItem struct {
		instance        *EntityCommissionClassItem
		commission      Commission
		commissionClass CommissionClass
		changed         map[string]bool
	}
)

func NewCommissionClassItem(instance *EntityCommissionClassItem) CommissionClassItem {
	if instance == nil {
		return &commissionClassItem{
			instance: &EntityCommissionClassItem{},
			changed:  make(map[string]bool),
		}
	}
	return &commissionClassItem{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (c *commissionClassItem) Entity() *EntityCommissionClassItem {
	return c.instance
}

func (c *commissionClassItem) SetCommission(commission Commission) {
	c.commission = commission
}

func (c *commissionClassItem) Commission() Commission {
	return c.commission
}

func (c *commissionClassItem) SetCommissionClass(commissionClass CommissionClass) {
	c.commissionClass = commissionClass
}

func (c *commissionClassItem) CommissionClass() CommissionClass {
	return c.commissionClass
}

func (c *commissionClassItem) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityCommissionClassItem](c.instance, fieldName, DbTagName, c.changed); err != nil {
		return fmt.Errorf("failed to set changed field %s: %w", fieldName, err)
	}
	return nil
}

func (c *commissionClassItem) IsChanged() bool {
	return len(c.changed) > 0
}

func (c *commissionClassItem) ToMap() map[string]any {
	if !c.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityCommissionClassItem](c.instance, c.changed, DbTagName)
}

func (c *commissionClassItem) Id() int {
	return c.instance.ID
}

func (c *commissionClassItem) CommissionId() int {
	return c.instance.CommissionID
}

func (c *commissionClassItem) CommissionClassId() int {
	return c.instance.CommissionClassID
}

func (c *commissionClassItem) OperationType() int {
	return c.instance.OperationType
}

func (c *commissionClassItem) GameId() int {
	return c.instance.GameID
}

func (c *commissionClassItem) SetId(id int) error {
	c.instance.ID = id
	return c.setChanged("ID")
}

func (c *commissionClassItem) SetCommissionId(commissionId int) error {
	c.instance.CommissionID = commissionId
	return c.setChanged("CommissionID")
}

func (c *commissionClassItem) SetCommissionClassId(commissionClassId int) error {
	c.instance.CommissionClassID = commissionClassId
	return c.setChanged("CommissionClassID")
}

func (c *commissionClassItem) SetOperationType(operationType int) error {
	c.instance.OperationType = operationType
	return c.setChanged("OperationType")
}

func (c *commissionClassItem) SetGameId(gameId int) error {
	c.instance.GameID = gameId
	return c.setChanged("GameID")
}

func (c *commissionClassItem) String() string {
	return fmt.Sprintf("CommissionClassItem(ID: %d, CommissionID: %d, CommissionClassID: %d, OperationType: %d, GameID: %d)",
		c.instance.ID, c.instance.CommissionID, c.instance.CommissionClassID, c.instance.OperationType, c.instance.GameID)
}
