package CharSet

import (
	"fmt"
	"sort"
)

// CharInterval
// [Start:Stop];
// single char interval is ['a':'a']
type CharInterval struct {
	Start, Stop uint64
}

// CharSet :
//set of character Intervals
type CharSet struct {
	Intervals []CharInterval
}

func NewCharSet(interval CharInterval) *CharSet {
	s := new(CharSet)
	s.Intervals = []CharInterval{interval}
	return s
}

func NewEmptyCharSet() *CharSet {
	s := new(CharSet)
	s.Intervals = make([]CharInterval, 0)
	return s
}

func CopyCharSet(set CharSet) *CharSet {
	s := new(CharSet)
	s.Intervals = append(s.Intervals, set.Intervals...)
	return s
}

func (c CharSet) IsEmpty() bool {
	if c.Intervals == nil || len(c.Intervals) == 0 {
		return true
	}
	return false
}

func (c CharSet) IsSingleChar() bool {
	if len(c.Intervals) == 1 && c.Intervals[0].Start == c.Intervals[0].Stop {
		return true
	}
	return false
}

func (c CharSet) Sort() { //first by Start, then by Stop
	sort.Slice(c.Intervals, func(i, j int) bool {
		return c.Intervals[i].Start < c.Intervals[j].Start ||
			(c.Intervals[i].Start == c.Intervals[j].Start && c.Intervals[i].Stop < c.Intervals[j].Stop)
	})
}

func (c *CharSet) Minimize() { //minimize number of Intervals if possible
	if len(c.Intervals) == 0 || len(c.Intervals) == 1 {
		return
	}

	c.Sort()

	var newIntervals []CharInterval
	newIntervals = append(newIntervals, c.Intervals[0])
	currIntervalIndex := 0
	for i := 1; i < len(c.Intervals); i++ {
		if newIntervals[currIntervalIndex].Stop >= c.Intervals[i].Start { //either intersection or concatenation
			newIntervals[currIntervalIndex].Stop = c.Intervals[i].Stop
		} else {
			//go to next interval
			currIntervalIndex++
			newIntervals = append(newIntervals, c.Intervals[i])
		}
	}

	c.Intervals = newIntervals
}

func (c *CharSet) Add(interval CharInterval) {
	c.Intervals = append(c.Intervals, interval)
}

func (c *CharSet) AddMultiple(intervals []CharInterval) {
	for _, interval := range intervals {
		c.Add(interval)
	}
}

func (c *CharSet) EqualTo(other CharSet) bool {
	c.Minimize()
	other.Minimize()

	if len(c.Intervals) != len(other.Intervals) {
		return false
	}

	for i := 0; i < len(c.Intervals); i++ {
		if c.Intervals[i] != other.Intervals[i] {
			return false
		}
	}

	return true
}

func (c *CharSet) GetCharSetList() []uint64 {
	list := make([]uint64, 0)

	for _, interval := range c.Intervals {
		intermediaryList := make([]uint64, interval.Stop-interval.Start+1)

		i := 0
		for c := interval.Start; c <= interval.Stop; c++ {
			intermediaryList[i] = c
			i++
		}

		list = append(list, intermediaryList...)
	}

	return list
}

func (c CharSet) Intersect(other CharSet) (newThis CharSet, newOther CharSet, intersection CharSet) {
	c.Minimize()
	other.Minimize()

	notableMap := make(map[uint64]notable)

	for _, interval := range c.Intervals {
		if interval.Start == interval.Stop {
			notable := notableMap[interval.Start]
			notable.index = interval.Start
			notable.this = Single
			notableMap[interval.Start] = notable
		} else {
			notable := notableMap[interval.Start]
			notable.index = interval.Start
			notable.this = Open
			notableMap[interval.Start] = notable

			notable = notableMap[interval.Stop]
			notable.index = interval.Stop
			notable.this = Close
			notableMap[interval.Stop] = notable
		}
	}

	for _, interval := range other.Intervals {
		if interval.Start == interval.Stop {
			notable := notableMap[interval.Start]
			notable.index = interval.Start
			notable.other = Single
			notableMap[interval.Start] = notable
		} else {
			notable := notableMap[interval.Start]
			notable.index = interval.Start
			notable.other = Open
			notableMap[interval.Start] = notable

			notable = notableMap[interval.Stop]
			notable.index = interval.Stop
			notable.other = Close
			notableMap[interval.Stop] = notable
		}
	}

	notables := make([]notable, len(notableMap))
	i := 0
	for _, notable := range notableMap {
		notables[i] = notable
		i++
	}

	//sort by index
	sort.Slice(notables, func(i, j int) bool {
		return notables[i].index < notables[j].index
	})

	newThis = CharSet{}
	newOther = CharSet{}
	intersection = CharSet{}

	thisNumbers := make([]uint64, 0)
	otherNumbers := make([]uint64, 0)
	intersectionNumbers := make([]uint64, 0)

	thisOpen := false
	otherOpen := false

	for _, notable := range notables {
		if thisOpen && otherOpen { //TT
			intersectionNumbers = append(intersectionNumbers, notable.index)
			if notable.this == NaN && notable.other == Close {
				otherOpen = false

				thisNumbers = append(thisNumbers, notable.index+1)
			} else if notable.this == Close && notable.other == NaN {
				thisOpen = false
				otherNumbers = append(otherNumbers, notable.index+1)
			} else if notable.this == Close && notable.other == Close {
				thisOpen = false
				otherOpen = false
			} else {
				print("Error in construction!")
			}
		} else if thisOpen && !otherOpen { //TF
			if notable.this == Close {
				thisOpen = false

				switch notable.other {
				case NaN:
					thisNumbers = append(thisNumbers, notable.index)
				case Open:
					otherOpen = true

					thisNumbers = append(thisNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					otherNumbers = append(otherNumbers, notable.index+1)
				case Single:
					thisNumbers = append(thisNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
				default:
					print("Error in construction!")
				}
			} else if notable.this == NaN {
				switch notable.other {
				case Open:
					otherOpen = true

					thisNumbers = append(thisNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
				case Single:
					thisNumbers = append(thisNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					thisNumbers = append(thisNumbers, notable.index+1)
				default:
					print("Error in construction!")
				}
			} else {
				print("Error in construction!")
			}
		} else if !thisOpen && otherOpen { //FT
			if notable.this == Close {
				thisOpen = false

				switch notable.other {
				case NaN:
					otherNumbers = append(otherNumbers, notable.index)
				case Open:
					otherOpen = true

					otherNumbers = append(otherNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					otherNumbers = append(otherNumbers, notable.index+1)
				case Single:
					otherNumbers = append(otherNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
				default:
					print("Error in construction!")
				}
			} else if notable.this == NaN {
				switch notable.other {
				case Open:
					otherOpen = true

					otherNumbers = append(otherNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
				case Single:
					otherNumbers = append(otherNumbers, notable.index-1)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					otherNumbers = append(otherNumbers, notable.index+1)
				default:
					print("Error in construction!")
				}
			} else {
				print("Error in construction!")
			}
		} else { //FF
			if notable.this == NaN {
				switch notable.other {
				case Open:
					otherOpen = true

					otherNumbers = append(otherNumbers, notable.index)
				case Single:
					otherNumbers = append(otherNumbers, notable.index)
					otherNumbers = append(otherNumbers, notable.index)
				default:
					print("Error in construction!")
				}
			} else if notable.this == Open {
				thisOpen = true

				switch notable.other {
				case NaN:
					thisNumbers = append(otherNumbers, notable.index)
				case Open:
					otherOpen = true

					intersectionNumbers = append(intersectionNumbers, notable.index)
				case Single:
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)

					thisNumbers = append(thisNumbers, notable.index)
				default:
					print("Error in construction!")
				}
			} else if notable.this == Single {
				switch notable.other {
				case NaN:
					thisNumbers = append(thisNumbers, notable.index)
					thisNumbers = append(thisNumbers, notable.index)
				case Open:
					otherOpen = true

					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
					otherNumbers = append(otherNumbers, notable.index+1)
				case Single:
					intersectionNumbers = append(intersectionNumbers, notable.index)
					intersectionNumbers = append(intersectionNumbers, notable.index)
				default:
					print("Error in construction!")
				}
			} else {
				print("Error in construction!")
			}
		}
	}

	if len(thisNumbers)%2 > 0 {
		thisNumbers = append(thisNumbers, 255) //TODO technically cheating, bound should come from parser
		print("WARNING, UNEVEN NUMBER OF ThisNumbers\n")
	}
	if len(otherNumbers)%2 > 0 {
		otherNumbers = append(otherNumbers, 255)
		print("WARNING, UNEVEN NUMBER OF OtherNumbers\n")

	}
	if len(intersectionNumbers)%2 > 0 {
		intersectionNumbers = append(intersectionNumbers, 255)
		print("WARNING, UNEVEN NUMBER OF IntersectionNumbers\n")
	}

	for i := 0; i < len(thisNumbers)/2; i++ {
		newThis.Add(CharInterval{thisNumbers[2*i], thisNumbers[2*i+1]})
	}

	for i := 0; i < len(otherNumbers)/2; i++ {
		newOther.Add(CharInterval{otherNumbers[2*i], otherNumbers[2*i+1]})
	}

	for i := 0; i < len(intersectionNumbers)/2; i++ {
		intersection.Add(CharInterval{intersectionNumbers[2*i], intersectionNumbers[2*i+1]})
	}

	return newThis, newOther, intersection
}

func (c *CharSet) DebugString() string {

	str := ""
	for _, interval := range c.Intervals {
		str += fmt.Sprintf("[%d-%d];", interval.Start, interval.Stop)
	}

	return fmt.Sprintf("(%s)", str)
}

func (c *CharSet) Copy() *CharSet {
	return &CharSet{
		append([]CharInterval{}, c.Intervals...),
	}
}

// useful for visitor only
func (c CharSet) Contains(val uint64) bool {
	for _, interval := range c.Intervals {
		if interval.Start <= val && interval.Stop >= val {
			return true
		}
	}

	return false
}

type NotableType int

const (
	NaN NotableType = iota
	Open
	Close
	Single
)

type notable struct {
	index uint64
	this  NotableType
	other NotableType
}
