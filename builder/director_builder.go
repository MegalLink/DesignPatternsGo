package builder

// directos let us set properties in our builder, otherwise builder can be modified and set properties from main.
type Director struct {
	builder IHouseBuilder
}

func NewDirector(b IHouseBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IHouseBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() House {
	d.builder.withDoor("wooden door")
	d.builder.withWindow("glass window")
	d.builder.withNumFloor(10)
	return d.builder.getHouse()
}
