package construction_patterns

import "fmt"

type Person struct {
	// Address
	StreetName, Postcode, City string

	// Occupation
	CompanyName, JobTitle string
	AnnualIncome          float32
}

func (p *Person) String() string {
	return fmt.Sprintf("%s/%s/%s", p.StreetName, p.Postcode, p.City)
}

type PersonBuilder struct {
	Person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

func (p *PersonBuilder) AddAddress(
	streetName, postcode, city string,
) {
	p.Person.StreetName = streetName
	p.Person.Postcode = postcode
	p.Person.City = city
}

func (p *PersonBuilder) AddOccupationInfo(
	company, job string, income float32,
) {
	p.Person.CompanyName = company
	p.Person.JobTitle = job
	p.Person.AnnualIncome = income
}
