package types

type Block struct {
	Key string
	Val interface{}
}

type MTFile struct {
	Blocks []Block
}

func NewMTFile(m []Block) *MTFile {
	return &MTFile{Blocks: m}
}
