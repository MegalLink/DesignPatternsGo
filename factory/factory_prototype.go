package factory

type Position int64

const (
	Developer = iota
	Manager
)

// used to set predefined objects
func NewEmployeePrototype(role Position) *Employee {
	switch role {
	case Developer:
		return &Employee{"", "developer", 60000}
	case Manager:
		return &Employee{"", "manager", 80000}
	default:
		panic("unsupported role")
	}
}
