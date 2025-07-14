package terminal_test

import (
	"fmt"
	"github.com/database-lab/terminal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommissionClassItem_NewInstance_ReturnsDefaultValues(t *testing.T) {
	item := terminal.NewCommissionClassItem(nil)

	assert.NotNil(t, item)
	assert.Equal(t, 0, item.Id())
	assert.Equal(t, 0, item.CommissionId())
	assert.Equal(t, 0, item.CommissionClassId())
	assert.Equal(t, 0, item.OperationType())
	assert.Equal(t, 0, item.GameId())
	assert.False(t, item.IsChanged())
	assert.Nil(t, item.ToMap())
}

func TestCommissionClassItem_SetFields_MarksAsChanged(t *testing.T) {
	instance := &terminal.EntityCommissionClassItem{}
	item := terminal.NewCommissionClassItem(instance)

	err := item.SetId(1)
	assert.NoError(t, err)
	assert.True(t, item.IsChanged())
	assert.Equal(t, 1, item.Id())

	err = item.SetCommissionId(2)
	assert.NoError(t, err)
	assert.True(t, item.IsChanged())
	assert.Equal(t, 2, item.CommissionId())

	err = item.SetCommissionClassId(3)
	assert.NoError(t, err)
	assert.True(t, item.IsChanged())
	assert.Equal(t, 3, item.CommissionClassId())

	err = item.SetOperationType(4)
	assert.NoError(t, err)
	assert.True(t, item.IsChanged())
	assert.Equal(t, 4, item.OperationType())

	err = item.SetGameId(5)
	assert.NoError(t, err)
	assert.True(t, item.IsChanged())
	assert.Equal(t, 5, item.GameId())
}

func TestCommissionClassItem_ToMap_ReturnsChangedFields(t *testing.T) {
	instance := &terminal.EntityCommissionClassItem{}
	item := terminal.NewCommissionClassItem(instance)

	_ = item.SetId(1)
	_ = item.SetCommissionId(2)
	_ = item.SetCommissionClassId(3)

	changedMap := item.ToMap()
	fmt.Println("Changed Map:", changedMap)
	assert.NotNil(t, changedMap)
	assert.Equal(t, 2, len(changedMap))
	assert.Equal(t, 1, changedMap["id"])
	assert.Equal(t, 2, changedMap["commission_id"])
	assert.Equal(t, 3, changedMap["commission_class_id"])
}

func TestCommissionClassItem_String_ReturnsFormattedString(t *testing.T) {
	instance := &terminal.EntityCommissionClassItem{
		ID:                1,
		CommissionID:      2,
		CommissionClassID: 3,
		OperationType:     4,
		GameID:            5,
	}
	item := terminal.NewCommissionClassItem(instance)

	expected := "CommissionClassItem(ID: 1, CommissionID: 2, CommissionClassID: 3, OperationType: 4, GameID: 5)"
	assert.Equal(t, expected, item.String())
}
