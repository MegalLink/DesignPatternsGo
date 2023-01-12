package flyweight

import (
	"strings"
	"unicode"
)

type FormmattedText struct {
	plainText  string
	capitalize []bool
}

func NewFormattedText(plainTExt string) *FormmattedText {
	return &FormmattedText{
		plainText:  plainTExt,
		capitalize: make([]bool, len(plainTExt)),
	}
}

func (f *FormmattedText) Capitalize(start, end int) {
	for i := start; i <= end; i++ {
		f.capitalize[i] = true
	}
}

func (f *FormmattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(f.plainText); i++ {
		c := f.plainText[i]
		if f.capitalize[i] {
			sb.WriteRune(unicode.ToUpper(rune(c)))
		} else {
			sb.WriteRune(rune(c))
		}
	}
	return sb.String()
}

// The approach above works but capitalize method makes a for when capitalize so we are doing double for, exist a better way to do this that is flyweight
type TextRange struct {
	Start, End       int
	Capitalize, Mask bool
}

func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange
}

func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{plainText: plainText}
}

// so with this function we just append the implementation bft.Range(16, 19).Capitalize = true or Mask
func (b *BetterFormattedText) Range(start, end int) *TextRange {
	r := &TextRange{Start: start, End: end, Capitalize: false, Mask: false}
	b.formatting = append(b.formatting, r)
	return r
}

func (b *BetterFormattedText) String() string {
	sb := strings.Builder{}
	for i := 0; i < len(b.plainText); i++ {
		c := b.plainText[i]
		for _, r := range b.formatting {
			if r.Covers(i) && r.Capitalize {
				c = uint8(unicode.ToUpper(rune(c)))
			} else if r.Covers(i) && r.Mask {
				c = uint8(rune('X'))
			}
		}
		sb.WriteRune(rune(c))
	}
	return sb.String()
}
