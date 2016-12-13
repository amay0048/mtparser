package mtparser

import (
	"bytes"
	"errors"
	"regexp"
	"text/scanner"
)

var bicSplit = regexp.MustCompile("(.{8})(.{1})(.*)")

func (s *Parser) scanHeader() (err error) {
	var c rune
	var max int
	var k string

	bin := len(s.Blocks)
	mp := map[string]Node{}
	s.blk.Val = []Header{}
	b := bytes.NewBufferString("")

	s.Mode = scanner.ScanChars

	c = s.Scan()
	switch c {
	case 'I':
		max = 21
	case 'O':
		max = 47
	default:
		max = 25
	}

	for i := 1; i <= max; i++ {
		switch c {
		case '{', '}', ':', '-', '/':
			err = errors.New(s.ErrMessage(c, false))
			return
		default:
			b.WriteRune(c)
		}

		switch max {
		case 21:
			k = inputHeader(i)
			break
		case 47:
			k = outputHeader(i)
			break
		default:
			k = basicHeader(i)
			break
		}

		if len(k) > 0 {
			s.blk.Val = append(s.blk.Val.([]Header), Header{
				Key: k,
				Val: b.String(),
			})
			mp[k] = Node{
				Val: b.String(),
				Blk: bin,
				Ind: i,
			}
			b.Reset()
		}

		if i == max {
			break
		}

		c = s.Scan()
	}

	if v, ok := mp["source"]; ok {
		if bic := bicSplit.FindStringSubmatch(v.Val); bic != nil {
			v.Det = map[string]string{
				"BIC":      bic[1],
				"terminal": bic[2],
				"branch":   bic[3],
			}
			mp["source"] = v
		}
	}
	if v, ok := mp["destination"]; ok {
		if bic := bicSplit.FindStringSubmatch(v.Val); bic != nil {
			v.Det = map[string]string{
				"BIC":      bic[1],
				"terminal": bic[2],
				"branch":   bic[3],
			}
			mp["destination"] = v
		}
	}

	s.Map[s.blk.Key] = mp
	s.Mode = scanner.ScanIdents
	return nil
}

func basicHeader(i int) string {
	switch i {
	case 1:
		return "application"
	case (2 + 1):
		return "service"
	case (12 + 2 + 1):
		return "source"
	case (4 + 12 + 2 + 1):
		return "session"
	case (6 + 4 + 12 + 2 + 1):
		return "sequence"
	}
	return ""
}

func inputHeader(i int) string {
	switch i {
	case 1:
		return "direction"
	case (3 + 1):
		return "type"
	case (12 + 3 + 1):
		return "destination"
	case (1 + 12 + 3 + 1):
		return "priority"
	case (1 + 1 + 12 + 3 + 1):
		return "monitoring"
	case (3 + 1 + 1 + 12 + 3 + 1):
		return "obsolescence"
	}
	return ""
}

func outputHeader(i int) string {
	switch i {
	case 1:
		return "direction"
	case (3 + 1):
		return "type"
	case (4 + 3 + 1):
		return "input_hhmm"
	case (6 + 4 + 3 + 1):
		return "input_ddmmyy"
	case (12 + 6 + 4 + 3 + 1):
		return "destination"
	case (4 + 12 + 6 + 4 + 3 + 1):
		return "session"
	case (6 + 4 + 12 + 6 + 4 + 3 + 1):
		return "sequence"
	case (6 + 6 + 4 + 12 + 6 + 4 + 3 + 1):
		return "out_ddmmyy"
	case (4 + 6 + 6 + 4 + 12 + 6 + 4 + 3 + 1):
		return "out_hhmm"
	case (1 + 4 + 6 + 6 + 4 + 12 + 6 + 4 + 3 + 1):
		return "priority"
	}
	return ""
}
