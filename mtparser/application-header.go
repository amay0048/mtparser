package mtparser

import "bytes"

func (t *parser) readAppHeader() []Header {
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
				Key: "direction",
				Val: buf.String(),
				Ord: 1,
			})
			buf.Reset()
			continue L
		case (3 + 1):
			heds = append(heds, Header{
				Key: "type",
				Val: buf.String(),
				Ord: 2,
			})
			buf.Reset()
			continue L
		case (4 + 3 + 1):
			heds = append(heds, Header{
				Key: "input_hhmm",
				Val: buf.String(),
				Ord: 3,
			})
			buf.Reset()
			continue L
		case (6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "input_ddmmyy",
				Val: buf.String(),
				Ord: 4,
			})
			buf.Reset()
			continue L
		case (12 + 6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "destination",
				Val: buf.String(),
				Ord: 5,
			})
			buf.Reset()
			continue L
		case (4 + 12 + 6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "session",
				Val: buf.String(),
				Ord: 6,
			})
			buf.Reset()
			continue L
		case (6 + 4 + 12 + 6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "sequence",
				Val: buf.String(),
				Ord: 7,
			})
			buf.Reset()
			continue L
		case (6 + 6 + 4 + 12 + 6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "out_ddmmyy",
				Val: buf.String(),
				Ord: 8,
			})
			buf.Reset()
			continue L
		case (4 + 6 + 6 + 4 + 12 + 6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "out_hhmm",
				Val: buf.String(),
				Ord: 9,
			})
			buf.Reset()
			continue L
		case (1 + 4 + 6 + 6 + 4 + 12 + 6 + 4 + 3 + 1):
			heds = append(heds, Header{
				Key: "priority",
				Val: buf.String(),
				Ord: 10,
			})
			buf.Reset()
			break L
		}
	}
	t.blockTerminus()
	return heds
}
