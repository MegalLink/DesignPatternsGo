package facade

type Buffer struct {
	width, height int
	buffer        []rune
}

func NewBuffer(width, height int) *Buffer {
	return &Buffer{width: width, height: height, buffer: make([]rune, width*height)}
}

func (b *Buffer) At(index int) rune {
	return b.buffer[index]
}

type Viewport struct {
	buffer *Buffer
	offset int
}

func NewViewport(buffer *Buffer) *Viewport {
	return &Viewport{buffer: buffer}
}

func (v *Viewport) GetCharacterAt(index int) rune {
	return v.buffer.At(v.offset + index)
}

// facade for viewports and buffers
type Console struct {
	buffer    []*Buffer
	viewports []*Viewport
	offset    int
}

func NewConsole() *Console {
	b := NewBuffer(200, 150)
	v := NewViewport(b)
	return &Console{[]*Buffer{b}, []*Viewport{v}, 0}
}

// so with this Console with can GetCharacter like in viewport
func (c *Console) GetCharacterAt(index int) rune {
	return c.viewports[0].GetCharacterAt(index)
}
