package parser

import (
	"bufio"
	"errors"
	"strconv"
	"text/scanner"

	"github.com/lunarway/mtparser/internal/types"
)

type Parser struct {
	scanner.Scanner
	blk       types.Block
	Blocks    []types.Block
	Map       ParserMap
	ErrPrefix string
}

func newParser(r *bufio.Reader) Parser {
	var s Parser
	s.Init(r)
	s.Mode = scanner.ScanIdents
	s.Whitespace = 1<<'\t' | 1<<'\r'
	s.IsIdentRune = func(ch rune, i int) bool {
		switch ch {
		case '{', '}', ':', '-', '\n', scanner.EOF:
			return false
		}
		return true
	}
	s.ErrPrefix = "We could not parse the payment message provided."
	s.Map = map[string]map[string]Node{}
	return s
}

func (s *Parser) ErrMessage(c rune, x bool) string {
	ln := strconv.Itoa(s.Pos().Line)
	cl := strconv.Itoa(s.Pos().Column)
	xp := "Expected"
	if !x {
		xp = "Unexpected"
	}
	return s.ErrPrefix + " " + xp + " '" + string(c) + "' at line " + ln + " column " + cl
}

func (s *Parser) parse() error {
	var err error

	for s.Peek() != scanner.EOF {
		if s.Scan() != '{' {
			return errors.New(s.ErrMessage('{', true))
		}

		s.Scan()
		s.blk.Key = s.TokenText()

		if s.Scan() != ':' {
			return errors.New(s.ErrMessage(':', true))
		}

		switch s.Peek() {
		case '\n':
			if err = s.scanBody(); err != nil {
				return err
			}
			break
		case '{':
			if err = s.scanBlocks(); err != nil {
				return err
			}
			break
		default:
			if err = s.scanHeader(); err != nil {
				return err
			}
			break
		}

		if s.Scan() != '}' {
			return errors.New(s.ErrMessage('}', true))
		}

		s.Blocks = append(s.Blocks, *&s.blk)
	}

	return nil
}

func Parse(r *bufio.Reader) (*types.MTFile, error) {
	p := newParser(r)
	err := p.parse()
	if err != nil {
		return nil, err
	}
	return types.NewMTFile(p.Blocks), nil
}
