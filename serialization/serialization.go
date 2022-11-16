package serialization

import (
	"bufio"

	"github.com/lunarway/mtparser/internal/parser"
	"github.com/lunarway/mtparser/internal/serializer"
	"github.com/lunarway/mtparser/internal/types"
)

func Serialize(file *types.MTFile) []byte {
	return serializer.Serialize(file)
}

func Parse(r *bufio.Reader) (*types.MTFile, error) {
	return parser.Parse(r)
}
