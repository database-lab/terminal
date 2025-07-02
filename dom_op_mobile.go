package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

type (
	OpMobile interface {
		Id() int
		OperatorId() int
		Email() string
		Login() string
		Pin() string
		LangCode() string
		PushKey() string
		ClientOS() int
		FirstName() string
		LastName() string
		Phone() string
		OfflineSyncData() db.JSONB
		SetId(id int) error
		SetOperatorId(operatorID int) error
		SetEmail(email string) error
		SetLogin(login string) error
		SetPin(pin string) error
		SetLangCode(langCode string) error
		SetPushKey(pushKey string) error
		SetClientOS(clientOS int) error
		SetFirstName(firstName string) error
		SetLastName(lastName string) error
		SetPhone(phone string) error
		SetOfflineSyncData(data db.JSONB) error
		IsChanged() bool
		Entity() *EntityOpMobile
		ToMap() map[string]interface{}
	}
	opMobile struct {
		instance *EntityOpMobile
		changed  map[string]bool
	}
)

func NewOpMobile(instance *EntityOpMobile) OpMobile {
	if instance == nil {
		return &opMobile{
			instance: &EntityOpMobile{},
			changed:  make(map[string]bool),
		}
	}
	return &opMobile{
		instance: instance,
		changed:  make(map[string]bool),
	}
}

func (o *opMobile) Entity() *EntityOpMobile {
	return o.instance
}

func (o *opMobile) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityOpMobile](o.instance, fieldName, DbTagName, o.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (o *opMobile) IsChanged() bool {
	return len(o.changed) > 0
}

func (o *opMobile) Id() int {
	return o.instance.ID
}

func (o *opMobile) OperatorId() int {
	return o.instance.OperatorID
}

func (o *opMobile) Email() string {
	return o.instance.Email
}

func (o *opMobile) Login() string {
	return o.instance.Login
}

func (o *opMobile) Pin() string {
	if o.instance.Pin == nil {
		return ""
	}
	return *o.instance.Pin
}

func (o *opMobile) LangCode() string {
	if o.instance.LangCode == nil {
		return ""
	}
	return *o.instance.LangCode
}

func (o *opMobile) PushKey() string {
	if o.instance.PushKey == nil {
		return ""
	}
	return *o.instance.PushKey
}

func (o *opMobile) ClientOS() int {
	if o.instance.ClientOS == nil {
		return 0
	}
	return *o.instance.ClientOS
}

func (o *opMobile) FirstName() string {
	if o.instance.FirstName == nil {
		return ""
	}
	return *o.instance.FirstName
}

func (o *opMobile) LastName() string {
	if o.instance.LastName == nil {
		return ""
	}
	return *o.instance.LastName
}

func (o *opMobile) Phone() string {
	if o.instance.Phone == nil {
		return ""
	}
	return *o.instance.Phone
}

func (o *opMobile) OfflineSyncData() db.JSONB {
	if o.instance.OfflineSyncData == nil {
		return make(db.JSONB)
	}
	return o.instance.OfflineSyncData
}

func (o *opMobile) SetId(id int) error {
	o.instance.ID = id
	return o.setChanged("ID")
}

func (o *opMobile) SetOperatorId(operatorID int) error {
	if o.instance.OperatorID == operatorID {
		return nil // No change needed
	}
	o.instance.OperatorID = operatorID
	return o.setChanged("OperatorID")
}

func (o *opMobile) SetEmail(email string) error {
	if o.instance.Email == email {
		return nil // No change needed
	}
	o.instance.Email = email
	return o.setChanged("Email")
}

func (o *opMobile) SetLogin(login string) error {
	if o.instance.Login == login {
		return nil // No change needed
	}
	o.instance.Login = login
	return o.setChanged("Login")
}

func (o *opMobile) SetPin(pin string) error {
	if o.instance.Pin != nil && *o.instance.Pin == pin {
		return nil // No change needed
	}
	o.instance.Pin = &pin
	return o.setChanged("Pin")
}

func (o *opMobile) SetLangCode(langCode string) error {
	if o.instance.LangCode != nil && *o.instance.LangCode == langCode {
		return nil // No change needed
	}
	o.instance.LangCode = &langCode
	return o.setChanged("LangCode")
}

func (o *opMobile) SetPushKey(pushKey string) error {
	if o.instance.PushKey != nil && *o.instance.PushKey == pushKey {
		return nil // No change needed
	}
	o.instance.PushKey = &pushKey
	return o.setChanged("PushKey")
}

func (o *opMobile) SetClientOS(clientOS int) error {
	if o.instance.ClientOS != nil && *o.instance.ClientOS == clientOS {
		return nil // No change needed
	}
	o.instance.ClientOS = &clientOS
	return o.setChanged("ClientOS")
}

func (o *opMobile) SetFirstName(firstName string) error {
	if o.instance.FirstName != nil && *o.instance.FirstName == firstName {
		return nil // No change needed
	}
	o.instance.FirstName = &firstName
	return o.setChanged("FirstName")
}

func (o *opMobile) SetLastName(lastName string) error {
	if o.instance.LastName != nil && *o.instance.LastName == lastName {
		return nil // No change needed
	}
	o.instance.LastName = &lastName
	return o.setChanged("LastName")
}

func (o *opMobile) SetPhone(phone string) error {
	if o.instance.Phone != nil && *o.instance.Phone == phone {
		return nil // No change needed
	}
	o.instance.Phone = &phone
	return o.setChanged("Phone")
}

func (o *opMobile) SetOfflineSyncData(data db.JSONB) error {
	if o.instance.OfflineSyncData != nil {
		return nil // No change needed
	}
	if db.EqualJSONB(o.instance.OfflineSyncData, data) {
		return nil // No change needed
	}
	o.instance.OfflineSyncData = data
	return o.setChanged("OfflineSyncData")
}

func (o *opMobile) ToMap() map[string]interface{} {
	return db.MapValuesFromChangedFields[*EntityOpMobile](o.instance, o.changed, DbTagName)
}
