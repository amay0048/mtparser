package acceptance_test

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/lunarway/mtparser/internal/parser"
	"github.com/lunarway/mtparser/internal/serializer"
	"github.com/stretchr/testify/require"
)

func TestSerialization(t *testing.T) {
	for i := 1; i < 9; i++ {
		filename := fmt.Sprintf("mt-103-%04d.txt", i)

		t.Run(filename, func(t *testing.T) {
			filepath := fmt.Sprintf("./files/%s", filename)
			input, err := os.ReadFile(filepath)
			require.NoError(t, err)

			mtfile, err := parser.Parse(bufio.NewReader(bytes.NewReader(input)))
			require.NoError(t, err)

			serialized := serializer.Serialize(mtfile)
			require.Equal(t, input, serialized)
		})
	}
}
