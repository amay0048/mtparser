package serializer

import (
	"bytes"

	"github.com/lunarway/mtparser/internal/parser"
	"github.com/lunarway/mtparser/internal/types"
)

type serializer struct {
	Blk types.Block // Most recent block.
	m   []types.Block
	buf *bytes.Buffer
}

func newSerializer(m []types.Block) *serializer {
	sl := new(serializer)
	sl.m = m
	sl.buf = bytes.NewBufferString("")
	return sl
}

func (s *serializer) serialize() []byte {
	for _, b := range s.m {
		// hopefully this is synchronous
		// TODO: come back and implement with atomic counters??
		s.Blk = b
		s.buf.WriteString("{")
		s.buf.WriteString(s.Blk.Key)
		s.buf.WriteString(":")
		switch b.Key {
		case "1":
			s.serializeHeader()
			break
		case "2":
			s.serializeHeader()
			break
		case "4": // wont work for mt 104
			s.serializeBody()
			break
		default:
			s.serializeBlock()
		}
		s.buf.WriteString("}")
	}
	return s.buf.Bytes()
}

func (s *serializer) serializeHeader() {
	for _, h := range s.Blk.Val.([]parser.Header) {
		s.buf.WriteString(h.Val)
	}
}

func (s *serializer) serializeBody() {
	s.buf.WriteString("\n")
	for _, f := range s.Blk.Val.([]parser.Field) {
		s.buf.WriteString(":")
		s.buf.WriteString(f.Key)
		s.buf.WriteString(":")
		s.buf.WriteString(f.Val)
		s.buf.WriteString("\n")
	}
	s.buf.WriteString("-")
}

func (s *serializer) serializeBlock() {
	for _, b := range s.Blk.Val.([]types.Block) {
		s.buf.WriteString("{")
		s.buf.WriteString(b.Key)
		s.buf.WriteString(":")
		s.buf.WriteString(b.Val.(string))
		s.buf.WriteString("}")
	}
}

func Serialize(file *types.MTFile) []byte {
	serializer := newSerializer(file.Blocks)
	return serializer.serialize()
}
