package mtparser

import "bytes"

func (t *parser) readBody() []Field {
	var flds []Field
	r := t.r
	buf := bytes.NewBufferString("")
	key := ""
L:
	for {
		c, _, e := r.ReadRune()
		check(e)

		switch c {
		case 10: // \n
			if buf.Len() == 0 && len(key) == 0 {
				break
			}

			if t.Peek() == 58 { // :
				flds = append(flds, Field{
					Key: key,
					Val: buf.String(),
				})
				break
			}
		case 58: // :
			if !(buf.Len() == 0 && len(key) == 0) {
				key = buf.String()
				buf.Reset()
			}
			continue L
		case 45: // -
			break L
		}

		buf.WriteRune(c)
	}
	t.blockTerminus()
	return flds
}
