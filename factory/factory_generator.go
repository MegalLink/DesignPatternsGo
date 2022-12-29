package factory

type Employee struct {
	Name        string
	Position    string
	AnualIncome int
}

// factory generator functional approach, gets params to set full object later
func NewEmployeeFactory(position string, annualIncome int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, annualIncome}
	}
}

// Factory generator other approach it does the same but needs more functionality , better user functional approach
// but int this approach we can modify Position and Anual Income , with functional Factory we can't do that
type EmployeeFactory struct {
	Position    string
	AnualIncome int
}

func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.AnualIncome}
}

func NewEmployeeFactory2(position string, annualIncome int) *EmployeeFactory {
	return &EmployeeFactory{position, annualIncome}
}
