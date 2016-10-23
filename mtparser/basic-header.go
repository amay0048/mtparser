package mtparser

import "bytes"

func (t *parser) readBasicHeader() []Header {
	var heds []Header
	i := 0
	r := t.r
	buf := bytes.NewBufferString("")
L:
	for {
		i++
		c, _, e := r.ReadRune()
		check(e)
		buf.WriteRune(c)
		switch i {
		case 1:
			heds = append(heds, Header{
				Key: "application",
				Val: buf.String(),
				Ord: 1,
			})
			buf.Reset()
			continue L
		case (2 + 1):
			heds = append(heds, Header{
				Key: "service",
				Val: buf.String(),
				Ord: 2,
			})
			buf.Reset()
			continue L
		case (12 + 2 + 1):
			heds = append(heds, Header{
				Key: "source",
				Val: buf.String(),
				Ord: 3,
			})
			buf.Reset()
			continue L
		case (4 + 12 + 2 + 1):
			heds = append(heds, Header{
				Key: "session",
				Val: buf.String(),
				Ord: 4,
			})
			buf.Reset()
			continue L
		case (6 + 4 + 12 + 2 + 1):
			heds = append(heds, Header{
				Key: "sequence",
				Val: buf.String(),
				Ord: 5,
			})
			buf.Reset()
			break L
		}
	}
	t.blockTerminus()
	return heds
}
