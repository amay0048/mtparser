package mtparser

import (
	"bytes"
	"io"
)

func (t *parser) ReadBlock() Block {
	var blk Block
	r := t.r
	buf := bytes.NewBufferString("")
	key := ""

L:
	for {
		c, _, e := r.ReadRune()

		if e == io.EOF {
			break L
		} else {
			check(e)
		}

		switch c {
		case 123: // {
			if len(key) == 0 {
				continue
			}

			val := []Block{}
			for { // Keep processing next until }
				val = append(val, t.ReadBlock())
				if t.Peek() != 123 {
					break
				}
			}
			blk.Key = key
			blk.Val = val
			t.blockTerminus()

			break L
		case 125: // }
			blk.Key = key
			blk.Val = buf.String()

			break L
		case 58: // :
			if len(key) == 0 {
				key = buf.String()
				buf.Reset()
			}

			switch key {
			case "1":
				blk.Key = key
				blk.Val = t.readBasicHeader()
				break L
			case "2":
				blk.Key = key
				blk.Val = t.readAppHeader()
				break L
			case "4":
				blk.Key = key
				blk.Val = t.readBody()
				break L
			}

			break
		default:
			buf.WriteRune(c)
		}
	}

	t.Blk = blk
	return t.Blk
}
