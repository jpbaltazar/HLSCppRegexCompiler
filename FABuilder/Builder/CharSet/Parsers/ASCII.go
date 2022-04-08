package Parsers

import (
	"fmt"
	"strconv"
	"strings"
	"thesisGoRemake/FABuilder/Builder/CharSet"
)

type ASCII struct {
}

func (A ASCII) ParseCCLiteral(s string) (uint64, error) {
	if len(s) == 1 { //'a'
		return uint64(s[0]), nil
	} else { //'\xFF'
		val, err := strconv.ParseUint(s[2:4], 16, 64)
		if err != nil {
			print(fmt.Sprintf("Error! Failed to parse %s", s))
			return 0, err
		}

		return val, nil
	}
}

func (A ASCII) ParseCharInterval(interval string) (CharSet.CharInterval, error) {
	set := CharSet.CharInterval{}
	if interval == "-" { //exception
		singleChar := uint64("-"[0])

		set.Start = singleChar
		set.Stop = singleChar

	} else {
		parts := strings.SplitN(interval, "-", 2)

		if interval[0] == "-"[0] { //'--a'
			set.Start = uint64("-"[0])

			stop, err := A.ParseCCLiteral(parts[1])
			if err != nil {
				return set, err
			}
			set.Stop = stop
		} else if interval[len(interval)-1] == "-"[0] { //'a--'
			start, err := A.ParseCCLiteral(parts[0])
			if err != nil {
				return set, err
			}
			set.Start = start

			set.Stop = uint64("-"[0])
		} else { //'a-c'
			start, err := A.ParseCCLiteral(parts[0])
			if err != nil {
				return set, err
			}

			stop, err := A.ParseCCLiteral(parts[1])
			if err != nil {
				return set, err
			}

			set.Start = start
			set.Stop = stop
		}

	}

	return set, nil
}

func (A ASCII) ParseCharIntervals(strIntervals []string) ([]CharSet.CharInterval, error) {
	ret := make([]CharSet.CharInterval, 0)

	for _, strinterval := range strIntervals {
		interval, err := A.ParseCharInterval(strinterval)
		if err != nil {
			return nil, err
		}

		ret = append(ret, interval)
	}

	return ret, nil
}

func (A ASCII) CreateNullCharSet() *CharSet.CharSet {
	c := CharSet.NewCharSet(CharSet.CharInterval{
		Start: 0,
		Stop:  0,
	})

	return c
}

func (A ASCII) CreateNewLineCharSet() *CharSet.CharSet {
	c := CharSet.NewCharSet(CharSet.CharInterval{
		Start: uint64("\n"[0]),
		Stop:  uint64("\n"[0]),
	})

	return c
}

func (A ASCII) CreateNewLineNullCharSet() *CharSet.CharSet {
	c := A.CreateNullCharSet()

	c.Add(CharSet.CharInterval{
		Start: uint64("\n"[0]),
		Stop:  uint64("\n"[0]),
	})

	return c
}

func (A ASCII) InvertCharSet(set *CharSet.CharSet) {
	set.Minimize()

	newCharIntervals := make([]CharSet.CharInterval, 0)

	start := uint64(0)
	stop := uint64(0)
	for _, interval := range set.Intervals {
		if interval.Stop > start {
			stop = interval.Start - 1
		}

		newCharIntervals = append(newCharIntervals, CharSet.CharInterval{
			Start: start,
			Stop:  stop})
	}

	if start < 256 {
		stop = 255
		newCharIntervals = append(newCharIntervals, CharSet.CharInterval{
			Start: start,
			Stop:  stop,
		})
	}

	set.Intervals = newCharIntervals
}

func (A ASCII) IntervalToString(interval CharSet.CharInterval) string {
	str := ""
	if interval.Start > 32 && interval.Start < 127 {
		str += fmt.Sprintf("%c", interval.Start)
	} else {
		str += fmt.Sprintf("%d", interval.Stop)
	}

	if interval.Start != interval.Stop {
		str += "-"

		if interval.Stop > 32 && interval.Stop < 127 {
			str += fmt.Sprintf("%c", interval.Stop)
		} else {
			str += fmt.Sprintf("%d", interval.Stop)
		}
	}

	return str
}
