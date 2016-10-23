package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"lab.identitii.com/identitii/mtparser/mtparser"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readerTest(nm string) []mtparser.Block {
	file, err := os.Open(nm)
	check(err)
	defer file.Close()

	r := bufio.NewReader(file)
	t := mtparser.NewParser(r)

	return t.Parse()
}

var block1 = regexp.MustCompile(`\{1:(.+?)\}`)
var basicExp = regexp.MustCompile(`(.{1})(.{2})(.{12})(.{4})(.{6})`)
var block2 = regexp.MustCompile(`\{2:(.+?)\}`)
var appExp = regexp.MustCompile(`(.{1})(.{3})(.{4})(.{6})(.{12})(.{4})(.{6})(.{6})(.{4})(.{1})`)
var block3 = regexp.MustCompile(`\{3:(.+?)\}`) // not mt103
var block4 = regexp.MustCompile(`(?s)\{4:\n(.+?)-\}`)
var bodyExp = regexp.MustCompile(`(?s):?(.+?):(.+?)(?:\n:|-)`)
var block5 = regexp.MustCompile(`\{5:(.+)\}`)
var blockExp = regexp.MustCompile(`(?s)\{(.+?):(.*?)\}`)

type Token struct {
	Key string
	Val interface{}
}

func regexBody(s string) []Token {
	var toks []Token

	for _, m := range bodyExp.FindAllStringSubmatch(s, -1) {
		if m != nil {
			toks = append(toks, Token{
				Key: m[1],
				Val: m[2],
			})
		}
	}

	return toks
}
func regexBlock(s string) []Token {
	var toks []Token

	for _, m := range blockExp.FindAllStringSubmatch(s, -1) {
		if m != nil {
			toks = append(toks, Token{
				Key: m[1],
				Val: m[2],
			})
		}
	}

	return toks
}

func regexBasic(s string) []Token {
	var toks []Token
	var k string

	m := basicExp.FindStringSubmatch(s)
L:
	for i, g := range m {
		switch i {
		case 0:
			continue L
		case 1:
			k = "application"
			break
		case 2:
			k = "service"
			break
		case 3:
			k = "source"
			break
		case 4:
			k = "session"
			break
		case 5:
			k = "sequence"
			break
		}
		toks = append(toks, Token{
			Key: k,
			Val: g,
		})
	}
	return toks
}

func regexApp(s string) []Token {
	var toks []Token
	var k string

	m := appExp.FindStringSubmatch(s)
L:
	for i, g := range m {
		switch i {
		case 0:
			continue L
		case 1:
			k = "direction"
			break
		case 2:
			k = "type"
			break
		case 3:
			k = "input_hhmm"
			break
		case 4:
			k = "input_ddmmyy"
			break
		case 5:
			k = "destination"
			break
		case 6:
			k = "session"
			break
		case 7:
			k = "sequence"
			break
		case 8:
			k = "out_ddmmyy"
			break
		case 9:
			k = "out_hhmm"
			break
		case 10:
			k = "priority"
			break
		}
		toks = append(toks, Token{
			Key: k,
			Val: g,
		})
	}

	return toks
}

func regexTest(nm string) []Token {
	file, err := os.Open(nm)
	check(err)
	defer file.Close()

	r := bufio.NewReader(file)
	var toks []Token

	// defer debug.TrackTime(time.Now(), "RegexParse")

	msg, err := ioutil.ReadAll(r)
	check(err)
	str := string(msg)

	bl := block1.FindStringSubmatch(str)
	if bl != nil {
		toks = append(toks, Token{
			Key: "1",
			Val: regexBasic(bl[len(bl)-1]),
		})
	}
	bl = block2.FindStringSubmatch(str)
	if bl != nil {
		toks = append(toks, Token{
			Key: "2",
			Val: bl[len(bl)-1],
		})
	}
	bl = block3.FindStringSubmatch(str)
	if bl != nil {
		toks = append(toks, Token{
			Key: "3",
			Val: bl[len(bl)-1],
		})
	}
	bl = block4.FindStringSubmatch(str)
	if bl != nil {
		toks = append(toks, Token{
			Key: "4",
			Val: regexBody(bl[len(bl)-1]),
		})
	}
	bl = block5.FindStringSubmatch(str)
	if bl != nil {
		toks = append(toks, Token{
			Key: "5",
			Val: regexBlock(bl[len(bl)-1]),
		})
	}

	file.Close()
	return toks
}

func BenchmarkReader(b *testing.B) {
	s := "./test/mt-103-0001.txt"
	for i := 0; i < b.N; i++ {
		_ = readerTest(s)
	}
}

func BenchmarkRegex(b *testing.B) {
	s := "./test/mt-103-0001.txt"
	for i := 0; i < b.N; i++ {
		_ = regexTest(s)
	}
}

func main() {

	bred := testing.Benchmark(BenchmarkReader)
	breg := testing.Benchmark(BenchmarkRegex)

	fmt.Println()
	fmt.Println("Regex")
	spew.Dump(breg)
	fmt.Println()
	fmt.Println("Reader")
	spew.Dump(bred)
	fmt.Println()

	return
}
