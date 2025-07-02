package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
	"time"
)

type (
	Location interface {
		Id() int
		OldId() int
		MerchantId() int
		NewMerchantId() int
		PrevMerchantId() int
		Name() string
		Status() int
		StatusReasonNote() string
		StatusDate() time.Time
		OpeningHoursId() int
		CollectorId() int
		AddressId() int
		Deleted() bool
		SetId(id int) error
		SetOldId(oldId int) error
		SetMerchantId(merchantId int) error
		SetNewMerchantId(newMerchantId int) error
		SetPrevMerchantId(prevMerchantId int) error
		SetName(name string) error
		SetStatus(status int) error
		SetStatusReasonNote(statusReasonNote string) error
		SetStatusDate(statusDate time.Time) error
		SetOpeningHoursId(openingHoursId int) error
		SetCollectorId(collectorId int) error
		SetAddressId(addressId int) error
		SetDeleted(deleted bool) error
		IsChanged() bool
		ToMap() map[string]any
		Entity() *EntityLocation
	}
	location struct {
		instance *EntityLocation
		changed  map[string]bool
	}
)

func NewLocation(instance *EntityLocation) Location {
	if instance == nil {
		return &location{
			instance: &EntityLocation{},
			changed:  make(map[string]bool),
		}
	}
	return &location{
		instance: instance,
		changed:  make(map[string]bool),
	}
}
func (l *location) Entity() *EntityLocation {
	return l.instance
}

func (l *location) IsChanged() bool {
	return len(l.changed) > 0
}

func (l *location) ToMap() map[string]any {
	if !l.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityLocation](l.instance, l.changed, DbTagName)
}

func (l *location) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityLocation](l.instance, fieldName, DbTagName, l.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (l *location) Id() int {
	return l.instance.ID
}

func (l *location) OldId() int {
	if l.instance.OldID == nil {
		return 0
	}
	return *l.instance.OldID
}

func (l *location) MerchantId() int {
	return l.instance.MerchantID
}

func (l *location) NewMerchantId() int {
	if l.instance.NewMerchantID == nil {
		return 0
	}
	return *l.instance.NewMerchantID
}

func (l *location) PrevMerchantId() int {
	if l.instance.PrevMerchantID == nil {
		return 0
	}
	return *l.instance.PrevMerchantID
}

func (l *location) Name() string {
	return l.instance.Name
}

func (l *location) Status() int {
	return l.instance.Status
}

func (l *location) StatusReasonNote() string {
	if l.instance.StatusReasonNote == nil {
		return ""
	}
	return *l.instance.StatusReasonNote
}

func (l *location) StatusDate() time.Time {
	if l.instance.StatusDate == nil {
		return time.Time{}
	}
	return *l.instance.StatusDate
}

func (l *location) OpeningHoursId() int {
	if l.instance.OpeningHoursID == nil {
		return 0
	}
	return *l.instance.OpeningHoursID
}

func (l *location) CollectorId() int {
	if l.instance.CollectorID == nil {
		return 0
	}
	return *l.instance.CollectorID
}

func (l *location) AddressId() int {
	if l.instance.AddressID == nil {
		return 0
	}
	return *l.instance.AddressID
}

func (l *location) Deleted() bool {
	return l.instance.Deleted
}

func (l *location) SetId(id int) error {
	l.instance.ID = id
	return l.setChanged("ID")
}

func (l *location) SetOldId(oldId int) error {
	if oldId == 0 {
		return nil
	} else {
		l.instance.OldID = &oldId
	}
	return l.setChanged("OldID")
}

func (l *location) SetMerchantId(merchantId int) error {
	l.instance.MerchantID = merchantId
	return l.setChanged("MerchantID")
}

func (l *location) SetNewMerchantId(newMerchantId int) error {
	if newMerchantId == 0 {
		return nil
	} else {
		l.instance.NewMerchantID = &newMerchantId
	}
	return l.setChanged("NewMerchantID")
}

func (l *location) SetPrevMerchantId(prevMerchantId int) error {
	if prevMerchantId == 0 {
		return nil
	} else {
		l.instance.PrevMerchantID = &prevMerchantId
	}
	return l.setChanged("PrevMerchantID")
}

func (l *location) SetName(name string) error {
	l.instance.Name = name
	return l.setChanged("Name")
}

func (l *location) SetStatus(status int) error {
	l.instance.Status = status
	return l.setChanged("Status")
}

func (l *location) SetStatusReasonNote(statusReasonNote string) error {
	if statusReasonNote == "" {
		return nil
	} else {
		l.instance.StatusReasonNote = &statusReasonNote
	}
	return l.setChanged("StatusReasonNote")
}

func (l *location) SetStatusDate(statusDate time.Time) error {
	if statusDate.IsZero() {
		return nil
	} else {
		l.instance.StatusDate = &statusDate
	}
	return l.setChanged("StatusDate")
}

func (l *location) SetOpeningHoursId(openingHoursId int) error {
	if openingHoursId == 0 {
		return nil
	} else {
		l.instance.OpeningHoursID = &openingHoursId
	}
	return l.setChanged("OpeningHoursID")
}

func (l *location) SetCollectorId(collectorId int) error {
	if collectorId == 0 {
		return nil
	} else {
		l.instance.CollectorID = &collectorId
	}
	return l.setChanged("CollectorID")
}

func (l *location) SetAddressId(addressId int) error {
	if addressId == 0 {
		return nil
	} else {
		l.instance.AddressID = &addressId
	}
	return l.setChanged("AddressID")
}

func (l *location) SetDeleted(deleted bool) error {
	l.instance.Deleted = deleted
	return l.setChanged("Deleted")
}
