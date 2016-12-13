package mtparser

import "text/scanner"

type ParserMap map[string]map[string]Node

type Parser struct {
	scanner.Scanner
	blk       Block
	Blocks    []Block
	Map       ParserMap
	ErrPrefix string
}

type Node struct {
	Val string            `json:"value" bson:"value"`
	Blk int               `json:"-" bson:"-"`
	Ind int               `json:"-" bson:"-"`
	Det map[string]string `json:"detail" bson:"detail"`
}

type Block struct {
	Key string
	Val interface{}
}

type Field struct {
	Key string
	Val string
}

type Header struct {
	Key string
	Val string
}
