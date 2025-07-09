package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

var (
	CommissionTypeOperations = map[int]string{
		CommissionTypeOperationSales:        "Sales",
		CommissionTypeOperationValidations:  "Validations",
		CommissionTypeOperationWorkingHours: "Working hours",
		CommissionTypeOperationMonthlySales: "Monthly sales",
		CommissionTypeOperationCashBalance:  "Cash balance",
	}
	CommissionTypeOperationsSlice = []map[string]any{
		{"id": CommissionTypeOperationSales, "name": CommissionTypeOperations[CommissionTypeOperationSales]},
		{"id": CommissionTypeOperationValidations, "name": CommissionTypeOperations[CommissionTypeOperationValidations]},
		{"id": CommissionTypeOperationWorkingHours, "name": CommissionTypeOperations[CommissionTypeOperationWorkingHours]},
		{"id": CommissionTypeOperationMonthlySales, "name": CommissionTypeOperations[CommissionTypeOperationMonthlySales]},
		{"id": CommissionTypeOperationCashBalance, "name": CommissionTypeOperations[CommissionTypeOperationCashBalance]},
	}
)

type (
	CommissionClass interface {
		Id() int
		Description() string
		SetId(id int) error
		SetDescription(description string) error
		IsChanged() bool
		ToMap() map[string]any
		Entity() *EntityCommissionClass
	}
	commissionClass struct {
		instance *EntityCommissionClass
		changed  map[string]bool
	}
)

func NewCommissionClass(instance *EntityCommissionClass) CommissionClass {
	if instance == nil {
		return &commissionClass{
			instance: &EntityCommissionClass{},
			changed:  make(map[string]bool),
		}
	}
	return &commissionClass{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (c *commissionClass) Entity() *EntityCommissionClass {
	return c.instance
}

func (c *commissionClass) IsChanged() bool {
	return len(c.changed) > 0
}

func (c *commissionClass) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityCommissionClass](c.instance, fieldName, DbTagName, c.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (c *commissionClass) Id() int {
	return c.instance.ID
}

func (c *commissionClass) Description() string {
	return c.instance.Description
}

func (c *commissionClass) SetId(id int) error {
	if c.instance.ID == id {
		return nil
	}
	c.instance.ID = id
	return c.setChanged("ID")
}

func (c *commissionClass) SetDescription(description string) error {
	if c.instance.Description == description {
		return nil
	}
	c.instance.Description = description
	return c.setChanged("Description")
}

func (c *commissionClass) ToMap() map[string]any {
	if !c.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityCommissionClass](c.instance, c.changed, DbTagName)
}
