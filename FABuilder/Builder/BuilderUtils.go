package Builder

import (
	"strconv"
	"strings"
)

func (d *DFABuilderListener) SeparateTypeModifierFromQuantifier(quantifier string) (string, string) {
	if len(quantifier) == 1 { //e.g. "+"
		return quantifier, ""
	} else if len(quantifier) == 2 { //e.g. "++"
		return quantifier[:0], quantifier[1:]
	} else {
		//all these cases:
		//{n}
		//{n}+
		//{n}*
		//{n,}
		//...
		//{n,m}
		//...

		found := strings.Index(quantifier, "}")
		if found == -1 {
			return "", ""
		}

		return quantifier[:found], quantifier[found:]
	}
}

func (d *DFABuilderListener) ParseNumberedQuantifier(quantifier string) (int64, int64) {
	quantifier = quantifier[1:] //strip "{"

	comma := strings.Index(quantifier, ",")
	end := strings.Index(quantifier, "}")

	if comma == -1 { //exact quantifier
		//we already know the value will be allowed
		val, err := strconv.ParseInt(quantifier[0:end], 10, 64)
		if err != nil {
			print(err)
		}

		return val, val
	}

	minimum, err := strconv.ParseInt(quantifier[0:comma], 10, 64)
	if err != nil {
		print(err)
	}

	if comma == end-1 { //"...,}" //at least
		return minimum, -1
	}

	//between
	maximum, err := strconv.ParseInt(quantifier[comma+1:end], 10, 64)
	if err != nil {
		print(err)
	}

	return minimum, maximum
}
