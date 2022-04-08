package Parsers

import (
	"thesisGoRemake/FABuilder/Builder/CharSet"
)

type UTF8 struct {
}

func (u UTF8) ParseCCLiteral(s string) (uint64, error) {
	panic("implement me")
}

func (u UTF8) ParseCharInterval(s string) (CharSet.CharInterval, error) {
	panic("implement me")
}
