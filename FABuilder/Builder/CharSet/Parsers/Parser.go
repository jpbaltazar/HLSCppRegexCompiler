package Parsers

import (
	"thesisGoRemake/FABuilder/Builder/CharSet"
)

type Parser interface {
	ParseCCLiteral(string) (uint64, error)
	ParseCharInterval(string) (CharSet.CharInterval, error)
	ParseCharIntervals([]string) ([]CharSet.CharInterval, error)
	CreateNewLineCharSet() *CharSet.CharSet //charset must contain "\n"
	CreateNullCharSet() *CharSet.CharSet
	CreateNewLineNullCharSet() *CharSet.CharSet //charset must contain "\n" and "\0"
	InvertCharSet(*CharSet.CharSet)

	IntervalToString(interval CharSet.CharInterval) string
}
