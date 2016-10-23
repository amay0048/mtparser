package mtparser

import (
	"bufio"
	"errors"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Block struct {
	Key string
	Val interface{}
}

type Field struct {
	Key string
	Val interface{}
	Ord int
}

type Header struct {
	Key string
	Val interface{}
	Ord int
}

type mtParser interface {
	// Reads and returns an array of all blocks
	Parse() []Block
	// Reads and returns the next block
	ReadBlock() Block
	// reads and returns the body fields as per mt spec
	readBody() Block
	// reads and returns the basic headers as per mt spec
	readBasicHeader() Block
	// reads and returns the application headers as per mt spec
	readAppHeader() Block
	// check the next rune is the block terminus char }
	blockTerminus()
	// Return the next rune then UnreadRune
	Peek() rune
	// TODO: may implement depending on requirements
	// Line() int
	// Col() int
}

type parser struct {
	Blk    Block // Most recent block.
	r      *bufio.Reader
	line   int
	column int
}

func (t *parser) blockTerminus() {
	c, _, e := t.r.ReadRune()
	check(e)
	if c != 125 {
		e = errors.New("Expected terminus of block")
		check(e)
	}
}

func (t *parser) Peek() rune {
	r := t.r
	nxt, _, e := r.ReadRune()
	check(e)
	r.UnreadRune()
	return nxt
}

func (t *parser) Parse() []Block {
	var blks []Block
	for {
		t.ReadBlock()
		if t.Blk.Val == nil {
			break
		}
		blks = append(blks, t.Blk)
	}
	return blks
}

func NewParser(r *bufio.Reader) *parser {
	tz := new(parser)
	tz.r = r
	return tz
}
