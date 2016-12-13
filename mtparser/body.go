package mtparser

import "errors"

func (s *Parser) scanBody() error {
	var fld Field
	var p rune
	mp := map[string]Node{}
	bin := len(s.Blocks)

	s.blk.Val = []Field{}

	if s.Scan() != '\n' {
		return errors.New(s.ErrMessage('\n', true))
	}

	for i2 := 0; s.Peek() != '}'; i2++ {

		if s.Scan() != ':' {
			return errors.New(s.ErrMessage(':', true))
		}

		s.Scan()
		fld.Key = s.TokenText()

		if s.Scan() != ':' {
			return errors.New(s.ErrMessage(':', true))
		}

		fld.Val = ""

		for i := 0; i <= 100; i++ {
			t := s.Scan()
			p = s.Peek()

			if p == '-' && t == '\n' {
				break
			}
			if p == ':' {
				break
			}

			fld.Val += s.TokenText()

			if i == 100 {
				return errors.New("Unclosed body, expected -")
			}
		}

		mp[fld.Key] = Node{
			Val: fld.Val,
			Blk: bin,
			Ind: i2,
		}
		s.blk.Val = append(s.blk.Val.([]Field), *&fld)

		if p == '-' {
			s.Scan()
			break
		}
	}

	s.Map[s.blk.Key] = mp
	return nil
}
