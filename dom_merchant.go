package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

const (
	MerchantTypeForeign = iota
	MerchantTypeOwnMerchant
	MerchantTypeCollector
)
const (
	MerchantStatusApproved = iota + 1
	MerchantStatusActive
	MerchantStatusInactive
	MerchantStatusTerminated
	MerchantStatusPending
)

var (
	MerchantStatuses = map[int]string{
		MerchantStatusApproved:   "Approved",
		MerchantStatusActive:     "Active",
		MerchantStatusInactive:   "Inactive",
		MerchantStatusTerminated: "Terminated",
		MerchantStatusPending:    "Pending",
	}
	MerchantStatusesSlice = []map[string]any{
		{"id": MerchantStatusApproved, "name": MerchantStatuses[MerchantStatusApproved]},
		{"id": MerchantStatusActive, "name": MerchantStatuses[MerchantStatusActive]},
		{"id": MerchantStatusInactive, "name": MerchantStatuses[MerchantStatusInactive]},
		{"id": MerchantStatusTerminated, "name": MerchantStatuses[MerchantStatusTerminated]},
		{"id": MerchantStatusPending, "name": MerchantStatuses[MerchantStatusPending]},
	}
	MerchantTypes = map[int]string{
		MerchantTypeForeign:     "Franchise",
		MerchantTypeOwnMerchant: "Owner",
		MerchantTypeCollector:   "Collector only",
	}
	MerchantTypesSlice = []map[string]any{
		{"id": MerchantTypeForeign, "name": MerchantTypes[MerchantTypeForeign]},
		{"id": MerchantTypeOwnMerchant, "name": MerchantTypes[MerchantTypeOwnMerchant]},
		{"id": MerchantTypeCollector, "name": MerchantTypes[MerchantTypeCollector]},
	}
)

type (
	Merchant interface {
		Id() int
		LastName() string
		FirstName() string
		Title() string
		StatusInt() int
		StatusString() string
		TypeInt() int
		TypeString() string
		CompanyName() string
		Deleted() bool
		LotteryPoint() bool
		ParentId() int
		SetId(id int) error
		SetLastName(lastName string) error
		SetFirstName(firstName string) error
		SetTitle(title string) error
		SetStatus(status int) error
		SetType(merchantType int) error
		SetCompanyName(companyName string) error
		SetDeleted(deleted bool) error
		SetLotteryPoint(lotteryPoint bool) error
		SetParentId(parentId int) error
		IsChanged() bool
		ToMap() map[string]any
		Entity() *EntityMerchant
	}
	merchant struct {
		instance *EntityMerchant
		changed  map[string]bool
	}
)

func NewMerchant(instance *EntityMerchant) Merchant {
	if instance == nil {
		return &merchant{
			instance: &EntityMerchant{},
			changed:  make(map[string]bool),
		}
	}
	return &merchant{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (m *merchant) Entity() *EntityMerchant {
	return m.instance
}

func (m *merchant) IsChanged() bool {
	return len(m.changed) > 0
}

func (m *merchant) ToMap() map[string]any {
	if !m.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityMerchant](m.instance, m.changed, DbTagName)
}

func (m *merchant) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityMerchant](m.instance, fieldName, DbTagName, m.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}
func (m *merchant) Id() int {
	return m.instance.ID
}

func (m *merchant) LastName() string {
	if m.instance.LastName == nil {
		return ""
	}
	return *m.instance.LastName
}

func (m *merchant) FirstName() string {
	if m.instance.FirstName == nil {
		return ""
	}
	return *m.instance.FirstName
}

func (m *merchant) Title() string {
	if m.instance.Title == nil {
		return ""
	}
	return *m.instance.Title
}

func (m *merchant) StatusInt() int {
	return m.instance.Status
}

func (m *merchant) StatusString() string {
	if _, ok := MerchantStatuses[m.instance.Status]; !ok {
		return "Unknown"
	}
	return MerchantStatuses[m.instance.Status]
}

func (m *merchant) TypeInt() int {
	return m.instance.MerchantType
}

func (m *merchant) TypeString() string {
	if _, ok := MerchantTypes[m.instance.MerchantType]; !ok {
		return "Unknown"
	}
	return MerchantTypes[m.instance.MerchantType]
}

func (m *merchant) CompanyName() string {
	return m.instance.CompanyName
}

func (m *merchant) Deleted() bool {
	return m.instance.Deleted
}

func (m *merchant) LotteryPoint() bool {
	return m.instance.LotteryPoint
}

func (m *merchant) ParentId() int {
	if m.instance.ParentMerchantID == nil {
		return 0
	}
	return *m.instance.ParentMerchantID
}

func (m *merchant) SetId(id int) error {
	if m.instance.ID == id {
		return nil
	}
	m.instance.ID = id
	return m.setChanged("ID")
}

func (m *merchant) SetLastName(lastName string) error {
	if m.instance.LastName != nil && *m.instance.LastName == lastName {
		return nil
	}
	m.instance.LastName = &lastName
	return m.setChanged("LastName")
}

func (m *merchant) SetFirstName(firstName string) error {
	if m.instance.FirstName != nil && *m.instance.FirstName == firstName {
		return nil
	}
	m.instance.FirstName = &firstName
	return m.setChanged("FirstName")
}

func (m *merchant) SetTitle(title string) error {
	if m.instance.Title != nil && *m.instance.Title == title {
		return nil
	}
	m.instance.Title = &title
	return m.setChanged("Title")
}

func (m *merchant) SetStatus(status int) error {
	if m.instance.Status == status {
		return nil
	}
	m.instance.Status = status
	return m.setChanged("Status")
}

func (m *merchant) SetType(merchantType int) error {
	if m.instance.MerchantType == merchantType {
		return nil
	}
	m.instance.MerchantType = merchantType
	return m.setChanged("MerchantType")
}

func (m *merchant) SetCompanyName(companyName string) error {
	if m.instance.CompanyName == companyName {
		return nil
	}
	m.instance.CompanyName = companyName
	return m.setChanged("CompanyName")
}

func (m *merchant) SetDeleted(deleted bool) error {
	if m.instance.Deleted == deleted {
		return nil
	}
	m.instance.Deleted = deleted
	return m.setChanged("Deleted")
}

func (m *merchant) SetLotteryPoint(lotteryPoint bool) error {
	if m.instance.LotteryPoint == lotteryPoint {
		return nil
	}
	m.instance.LotteryPoint = lotteryPoint
	return m.setChanged("LotteryPoint")
}

func (m *merchant) SetParentId(parentId int) error {
	if m.instance.ParentMerchantID != nil && *m.instance.ParentMerchantID == parentId {
		return nil
	}
	if parentId == 0 {
		return nil
	} else {
		m.instance.ParentMerchantID = &parentId
	}
	return m.setChanged("ParentMerchantID")
}
