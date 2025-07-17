package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

type (
	OperatorPrivileges interface {
		Id() int
		OperatorId() int
		PrivilegeId() int
		GameId() int
		SetId(id int) error
		SetOperatorId(operatorID int) error
		SetPrivilegeId(privilegeID int) error
		SetGameId(gameID int) error
		IsChanged() bool
		ToMap() map[string]interface{}
		Entity() *EntityOperatorPrivileges
		String() string
	}
	operatorPrivileges struct {
		instance *EntityOperatorPrivileges
		changed  map[string]bool
	}
)

func NewOperatorPrivileges(instance *EntityOperatorPrivileges) OperatorPrivileges {
	if instance == nil {
		return &operatorPrivileges{
			instance: &EntityOperatorPrivileges{},
			changed:  make(map[string]bool),
		}
	}
	return &operatorPrivileges{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (o *operatorPrivileges) Id() int {
	return o.instance.ID
}

func (o *operatorPrivileges) OperatorId() int {
	return o.instance.OperatorID
}

func (o *operatorPrivileges) PrivilegeId() int {
	return o.instance.PrivilegeID
}

func (o *operatorPrivileges) GameId() int {
	return o.instance.GameID
}

func (o *operatorPrivileges) Entity() *EntityOperatorPrivileges {
	return o.instance
}

func (o *operatorPrivileges) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityOperatorPrivileges](o.instance, fieldName, DbTagName, o.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (o *operatorPrivileges) IsChanged() bool {
	return len(o.changed) > 0
}

func (o *operatorPrivileges) String() string {
	return fmt.Sprintf("OperatorPrivileges{ID: %d, OperatorID: %d, PrivilegeID: %d, GameID: %d}",
		o.instance.ID, o.instance.OperatorID, o.instance.PrivilegeID, o.instance.GameID)
}

func (o *operatorPrivileges) SetId(id int) error {
	if o.instance.ID == id {
		return nil // No change needed
	}
	o.instance.ID = id
	return o.setChanged("ID")
}

func (o *operatorPrivileges) SetOperatorId(operatorID int) error {
	if o.instance.OperatorID == operatorID {
		return nil // No change needed
	}
	o.instance.OperatorID = operatorID
	return o.setChanged("OperatorID")
}

func (o *operatorPrivileges) SetPrivilegeId(privilegeID int) error {
	if o.instance.PrivilegeID == privilegeID {
		return nil // No change needed
	}
	o.instance.PrivilegeID = privilegeID
	return o.setChanged("PrivilegeID")
}

func (o *operatorPrivileges) SetGameId(gameID int) error {
	if o.instance.GameID == gameID {
		return nil // No change needed
	}
	o.instance.GameID = gameID
	return o.setChanged("GameID")
}

func (o *operatorPrivileges) ToMap() map[string]interface{} {
	if !o.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityOperatorPrivileges](o.instance, o.changed, DbTagName)
}
