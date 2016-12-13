package mtparser

import "errors"

func (s *Parser) scanBlocks() error {
	var blk Block

	mp := map[string]Node{}
	s.blk.Val = []Block{}
	bin := len(s.Blocks)

	for i := 0; i <= 100; i++ {
		if s.Scan() != '{' {
			return errors.New(s.ErrMessage('{', true))
		}

		s.Scan()
		blk.Key = s.TokenText()

		if s.Scan() != ':' {
			return errors.New(s.ErrMessage(':', true))
		}

		blk.Val = ""
		if s.Scan() != '}' {
			blk.Val = s.TokenText()

			if s.Scan() != '}' {
				return errors.New(s.ErrMessage('}', true))
			}
		}

		mp[blk.Key] = Node{
			Val: blk.Val.(string),
			Blk: bin,
			Ind: i,
		}
		s.blk.Val = append(s.blk.Val.([]Block), *&blk)

		if s.Peek() == '}' {
			break
		}
		if i == 100 {
			return errors.New("Unclosed block, expected }")
		}
	}

	s.Map[s.blk.Key] = mp
	return nil
}
