package prototype

import (
	"encoding/json"

	"github.com/Jeffail/gabs/v2"
)

// prototype is used when i want to get something that is already designed and add extra functionality or change just something from this for prototype is neccesary DeepCopy
type Person struct {
	Name    string
	Address *Address
	Friends []string
}

type Address struct {
	StreetAddress, City, Country string
}

// with this method deep copy we need to handle al cases from pointers , child pointer, slices, etc. so better no use this method
func (p *Person) DeepCopy() *Person {
	newPerson := *p // important to copy like this
	newPerson.Address = p.Address.DeepCopy()
	copy(newPerson.Friends, p.Friends)
	return &newPerson
}

func (a *Address) DeepCopy() *Address {
	return &Address{

		StreetAddress: a.StreetAddress,
		City:          a.City,
		Country:       a.Country,
	}
}

// Better use this way to copy , you can use marshal to transform to bytes directly but i like gabs library xD
func (p *Person) DeepCopySerialization() *Person {
	var newPerson Person
	bytes := gabs.Wrap(p).Bytes()
	json.Unmarshal(bytes, &newPerson)

	return &newPerson
}

// prototype Implementation

var defaultPerson = Person{
	Address: &Address{
		StreetAddress: "456",
		City:          "Quito",
		Country:       "Ecuador"},
	Friends: []string{"One"},
}

func newPerson(proto *Person, name string) *Person {
	result := proto.DeepCopySerialization()
	result.Name = name
	return result
}

func NewPrototypedPerson(name string) *Person {
	return newPerson(&defaultPerson, name)
}
