package mtserializer_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/amay0048/mtparser/mtparser"
	"github.com/amay0048/mtparser/mtserializer"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readerTest(nm string) mtparser.Parser {
	file, err := os.Open(nm)
	check(err)
	defer file.Close()

	r := bufio.NewReader(file)
	psr := mtparser.New(r)
	if err = psr.Parse(); err != nil {
		fmt.Println(err)
	}

	return psr
}

func serializerTest(m []mtparser.Block) string {
	sl := mtserializer.New(m)
	return string(sl.Serialize())
}

func benchmarkReader(b *testing.B) {
	s := "./test/mt-103-0001.txt"
	for i := 0; i < b.N; i++ {
		_ = readerTest(s)
	}
}

func benchmarkSerialiser(b *testing.B) {
	f := "./test/mt-103-0001.txt"
	psr := readerTest(f)
	for i := 0; i < b.N; i++ {
		_ = serializerTest(psr.Blocks)
	}
}

func TestTickle(t *testing.T) {
	var psr mtparser.Parser
	var msg string
	var sts []string

	mtparser.TextRegexCompilation()
	fmt.Println("Regex compilation success")

	psr = readerTest("./test/mt-103-0001.txt")
	msg = serializerTest(psr.Blocks)
	// fmt.Println(msg)

	sts = psr.BodyValueStructured("23E")
	fmt.Println(sts)
	sts = psr.BodyValueStructured("32A")
	fmt.Println(sts)
	sts = psr.BodyValueStructured("33B")
	fmt.Println(sts)
	sts = psr.BodyValueStructured("71F")
	fmt.Println(sts)

	psr = readerTest("./test/mt-103-0002.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0002.txt", msg)

	psr = readerTest("./test/mt-103-0003.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0003.txt", msg)

	psr = readerTest("./test/mt-103-0004.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0004.txt", msg)

	psr = readerTest("./test/mt-103-0005.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0005.txt", msg)

	psr = readerTest("./test/mt-103-0006.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0006.txt", msg)

	psr = readerTest("./test/mt-103-0007.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0007.txt", msg)

	psr = readerTest("./test/mt-103-0008.txt")
	msg = serializerTest(psr.Blocks)
	fmt.Println("./test/mt-103-0008.txt", msg)

	reader := testing.Benchmark(benchmarkReader)
	serial := testing.Benchmark(benchmarkSerialiser)

	fmt.Println()
	fmt.Println("Reader")
	fmt.Println(reader)
	fmt.Println()
	fmt.Println("Serializer")
	fmt.Println(serial)
	fmt.Println()

	return
}
