package Flags

import "strings"

type Flags struct {
	//perl
	PERL_i, PERL_s, PERL_m, PERL_x bool

	//pcre
	PCRE_A, PCRE_E, PCRE_G bool

	//SNORT
	SNORT_R, SNORT_U, SNORT_I, SNORT_P                                     bool
	SNORT_H                                                                bool
	SNORT_D, SNORT_M, SNORT_C, SNORT_K, SNORT_S, SNORT_Y, SNORT_B, SNORT_O bool

	//OTHERS
	StartOfString bool //separate flag as it can also be set like this: "/^.../"
}

func NewFlags(str string) Flags {
	f := Flags{}

	//PERL
	if strings.Contains(str, "i") {
		f.PERL_i = true
	}
	if strings.Contains(str, "s") {
		f.PERL_s = true
	}
	if strings.Contains(str, "m") {
		f.PERL_m = true
	}
	if strings.Contains(str, "x") {
		f.PERL_x = true
	}

	//PCRE
	if strings.Contains(str, "A") {
		f.PCRE_A = true
	}
	if strings.Contains(str, "E") {
		f.PCRE_E = true
	}
	if strings.Contains(str, "G") {
		f.PCRE_G = true
	}

	//SNORT
	if strings.Contains(str, "R") {
		f.SNORT_R = true
	}
	if strings.Contains(str, "U") {
		f.SNORT_U = true
	}
	if strings.Contains(str, "I") {
		f.SNORT_I = true
	}
	if strings.Contains(str, "P") {
		f.SNORT_P = true
	}

	if strings.Contains(str, "H") {
		f.SNORT_H = true
	}

	if strings.Contains(str, "D") {
		f.SNORT_D = true
	}
	if strings.Contains(str, "M") {
		f.SNORT_M = true
	}
	if strings.Contains(str, "C") {
		f.SNORT_C = true
	}
	if strings.Contains(str, "K") {
		f.SNORT_K = true
	}
	if strings.Contains(str, "S") {
		f.SNORT_S = true
	}
	if strings.Contains(str, "Y") {
		f.SNORT_Y = true
	}
	if strings.Contains(str, "B") {
		f.SNORT_B = true
	}
	if strings.Contains(str, "O") {
		f.SNORT_O = true
	}

	return f
}
