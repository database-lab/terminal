package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

var (
	CommissionTypes = map[int]string{
		CommissionTypeFixedRate:       "Fixed rate",
		CommissionTypeFixedPercentage: "Percentage",
		CommissionTypeWorkHour:        "Work hour",
	}
	CommissionTypesSlice = []map[string]any{
		{"id": CommissionTypeFixedRate, "name": CommissionTypes[CommissionTypeFixedRate]},
		{"id": CommissionTypeFixedPercentage, "name": CommissionTypes[CommissionTypeFixedPercentage]},
		{"id": CommissionTypeWorkHour, "name": CommissionTypes[CommissionTypeWorkHour]},
	}
)

type (
	Commission interface {
		Id() int
		Type() int
		TypeString() string
		Description() string
		Value() float64
		SetId(id int) error
		SetType(t int) error
		SetDescription(description string) error
		SetValue(value float64) error
		IsChanged() bool
		ToMap() map[string]any
		Entity() *EntityCommission
	}
	commission struct {
		instance *EntityCommission
		changed  map[string]bool
	}
)

func NewCommission(instance *EntityCommission) Commission {
	if instance == nil {
		return &commission{
			instance: &EntityCommission{},
			changed:  make(map[string]bool),
		}
	}
	return &commission{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (c *commission) Entity() *EntityCommission {
	return c.instance
}

func (c *commission) IsChanged() bool {
	return len(c.changed) > 0
}

func (c *commission) ToMap() map[string]any {
	if !c.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityCommission](c.instance, c.changed, DbTagName)
}

func (c *commission) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityCommission](c.instance, fieldName, DbTagName, c.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`, %v", fieldName, err)
	}
	return nil
}

func (c *commission) Id() int {
	return c.instance.ID
}

func (c *commission) Type() int {
	return c.instance.Type
}

func (c *commission) TypeString() string {
	if name, ok := CommissionTypes[c.instance.Type]; ok {
		return name
	}
	return "Unknown Commission Type"
}

func (c *commission) Description() string {
	return c.instance.Description
}

func (c *commission) Value() float64 {
	return c.instance.Value
}

func (c *commission) SetId(id int) error {
	if c.instance.ID == id {
		return nil
	}
	c.instance.ID = id
	return c.setChanged("ID")
}

func (c *commission) SetType(t int) error {
	if c.instance.Type == t {
		return nil
	}
	c.instance.Type = t
	return c.setChanged("Type")
}

func (c *commission) SetDescription(description string) error {
	if c.instance.Description == description {
		return nil
	}
	c.instance.Description = description
	return c.setChanged("Description")
}

func (c *commission) SetValue(value float64) error {
	if c.instance.Value == value {
		return nil
	}
	c.instance.Value = value
	return c.setChanged("Value")
}
