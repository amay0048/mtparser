package parser

type ParserMap map[string]map[string]Node

type Node struct {
	Val string            `json:"value" bson:"value"`
	Blk int               `json:"-" bson:"-"`
	Ind int               `json:"-" bson:"-"`
	Det map[string]string `json:"detail" bson:"detail"`
}

type Field struct {
	Key string
	Val string
}

type Header struct {
	Key string
	Val string
}
