package mtparser

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
	"text/scanner"
)

func New(r *bufio.Reader) Parser {
	var s Parser
	s.Init(r)
	// s.Filename = "mt-message"
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

// Parse an MT103 string. Return a result
func ParseFromString(mt103 string) (mtParser Parser, err error) {

	mtReader := strings.NewReader(mt103)
	mtBufio := bufio.NewReader(mtReader)
	mtParser = New(mtBufio)
	err = mtParser.Parse()

	return
}

func (s *Parser) Parse() error {
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
