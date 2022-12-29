package builder

// to not use factory function with 10 arguments is not productive
// provide an API for contructing object step by step
// The potential problem with the multistage building process is that a partially built and unstable product may be exposed to the client. The Builder pattern keeps the product private until it’s fully built.
type House struct {
	WindowType string
	DoorType   string
	Floor      int
}

type IHouseBuilder interface {
	withWindow(value string)
	withDoor(value string)
	withNumFloor(value int)
	getHouse() House
}

func NewHouseBuilder() IHouseBuilder {
	return &House{}
}

func (b *House) withWindow(value string) {
	b.WindowType = value
}

func (b *House) withDoor(value string) {
	b.DoorType = value
}

func (b *House) withNumFloor(value int) {
	b.Floor = value
}

func (b *House) getHouse() House {
	return House{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
