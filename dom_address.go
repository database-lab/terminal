package terminal

import (
	"fmt"
	db "github.com/database-lab/dbutils"
)

type (
	Address interface {
		Id() int
		City() string
		Street() string
		Zip() string
		State() string
		Region() string
		Country() string
		Description() string
		SetId(id int) error
		SetCity(city string) error
		SetStreet(street string) error
		SetZip(zip string) error
		SetState(state string) error
		SetRegion(region string) error
		SetCountry(country string) error
		SetDescription(description string) error
		IsChanged() bool
		ToMap() map[string]any
		Entity() *EntityAddress
		String() string
	}
	address struct {
		instance *EntityAddress
		changed  map[string]bool
	}
)

func NewAddress(instance *EntityAddress) Address {
	if instance == nil {
		return &address{
			instance: &EntityAddress{},
			changed:  make(map[string]bool),
		}
	}
	return &address{
		instance: instance,
		changed:  make(map[string]bool),
	}
}
func (a *address) Entity() *EntityAddress {
	return a.instance
}

func (a *address) Id() int {
	return a.instance.ID
}
func (a *address) City() string {
	return a.instance.City
}
func (a *address) Street() string {
	return a.instance.Street
}
func (a *address) Zip() string {
	return a.instance.Zip
}
func (a *address) State() string {
	return a.instance.State
}
func (a *address) Region() string {
	return a.instance.Region
}
func (a *address) Country() string {
	return a.instance.Country
}
func (a *address) Description() string {
	return a.instance.Description
}

func (a *address) setChanged(fieldName string) error {
	if err := db.SetChangedField[*EntityAddress](a.instance, fieldName, DbTagName, a.changed); err != nil {
		return fmt.Errorf("failed to set changed field `%s`: %w", fieldName, err)
	}
	return nil
}

func (a *address) IsChanged() bool {
	return len(a.changed) > 0
}

func (a *address) SetId(id int) error {
	if a.instance.ID == id {
		return nil
	}
	a.instance.ID = id
	return a.setChanged("ID")
}

func (a *address) SetCity(city string) error {
	if a.instance.City == city {
		return nil
	}
	a.instance.City = city
	return a.setChanged("City")
}

func (a *address) SetStreet(street string) error {
	if a.instance.Street == street {
		return nil
	}
	a.instance.Street = street
	return a.setChanged("Street")
}

func (a *address) SetZip(zip string) error {
	if a.instance.Zip == zip {
		return nil
	}
	a.instance.Zip = zip
	return a.setChanged("Zip")
}

func (a *address) SetState(state string) error {
	if a.instance.State == state {
		return nil
	}
	a.instance.State = state
	return a.setChanged("State")
}

func (a *address) SetRegion(region string) error {
	if a.instance.Region == region {
		return nil
	}
	a.instance.Region = region
	return a.setChanged("Region")
}

func (a *address) SetCountry(country string) error {
	if a.instance.Country == country {
		return nil
	}
	a.instance.Country = country
	return a.setChanged("Country")
}

func (a *address) SetDescription(description string) error {
	if a.instance.Description == description {
		return nil
	}
	a.instance.Description = description
	return a.setChanged("Description")
}

func (a *address) String() string {
	return fmt.Sprintf("Address(ID: %d, City: %s, Street: %s, Zip: %s, State: %s, Region: %s, Country: %s, Description: %s)",
		a.instance.ID, a.instance.City, a.instance.Street, a.instance.Zip, a.instance.State, a.instance.Region, a.instance.Country, a.instance.Description)
}

func (a *address) ToMap() map[string]interface{} {
	if !a.IsChanged() {
		return nil
	}
	return db.MapValuesFromChangedFields[*EntityAddress](a.instance, a.changed, DbTagName)
}
