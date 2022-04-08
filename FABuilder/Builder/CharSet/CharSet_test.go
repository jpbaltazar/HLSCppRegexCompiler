package CharSet

import (
	"fmt"
	"testing"
)

func TestConditions_isEmpty(t *testing.T) {
	c := CharSet{}

	if !c.IsEmpty() {
		t.Errorf("IsEmpty reported non empty\n")
	}
}

func TestConditions_isSingleChar(t *testing.T) {
	var intervals []CharInterval
	intervals = append(intervals, CharInterval{1, 1})
	c := CharSet{intervals}

	if !c.IsSingleChar() {
		t.Errorf("IsSingleChar did not report true\n")
	}
}

func TestConditions_sort(t *testing.T) {
	var intervals []CharInterval
	intervals = append(intervals, CharInterval{1, 126})
	intervals = append(intervals, CharInterval{0, 255})
	c := CharSet{Intervals: intervals}

	//Start sorting
	c.Sort()
	if intervals[0].Start == 0 && intervals[0].Stop == 126 {
		t.Errorf("Wrong order\n")
	}
	if intervals[1].Start == 1 && intervals[1].Stop == 255 {
		t.Errorf("Wrong order\n")
	}

	//Stop sorting

	intervals[0] = CharInterval{2, 254}
	intervals[1] = CharInterval{2, 255}
	c.Sort()

	if intervals[0].Stop != 254 {
		t.Errorf("Doesn't sort by Stop integer if Start integer is the same")
	}
}

func TestConditions_minimize(t *testing.T) {
	var intervals []CharInterval

	//intersection
	intervals = append(intervals, CharInterval{1, 5})
	intervals = append(intervals, CharInterval{2, 7})
	//expected to merge 1:5 and 2:7 to 1:7

	//concatenation
	intervals = append(intervals, CharInterval{10, 15})
	intervals = append(intervals, CharInterval{15, 20})
	//expected to merge 10:15 and 15:20 to 10:20

	//no intersection
	intervals = append(intervals, CharInterval{30, 34})
	intervals = append(intervals, CharInterval{36, 40})
	//expected to keep the same Intervals

	c := CharSet{intervals}

	c.Minimize()

	if c.Intervals[0].Start != 1 && c.Intervals[0].Stop != 7 {
		t.Errorf("Wrong merge: got %d:%d, expected %d:%d\n",
			c.Intervals[0].Start, c.Intervals[0].Stop, 1, 7)
	}

	if c.Intervals[1].Start != 10 && c.Intervals[1].Stop != 20 {
		t.Errorf("Wrong merge: got %d:%d, expected %d:%d\n",
			c.Intervals[1].Start, c.Intervals[1].Stop, 10, 20)
	}

	if c.Intervals[2].Start != 30 && c.Intervals[2].Stop != 34 &&
		c.Intervals[3].Start != 36 && c.Intervals[3].Stop != 40 {

		t.Errorf("Unexpected merge: got %d:%d and %d:%d, expected %d:%d and %d:%d\n",
			c.Intervals[2].Start, c.Intervals[2].Stop,
			c.Intervals[3].Start, c.Intervals[3].Stop,
			30, 34,
			36, 40)
	}

	if len(c.Intervals) != 4 {
		t.Errorf("Intervals wrongly merged, had %d Intervals, expected %d\n", len(c.Intervals), 4)
	}
}

func TestConditions_intersect(t *testing.T) {
	thisCS := CharSet{
		[]CharInterval{
			{0, 5},
			{7, 15},
		},
	}

	otherCS := CharSet{
		[]CharInterval{
			{3, 6},
			{8, 20},
		},
	}

	thisCS, otherCS, intersection := thisCS.Intersect(otherCS)

	expectedThisCS := CharSet{
		[]CharInterval{
			{0, 2},
			{7, 7},
		},
	}

	if !thisCS.EqualTo(expectedThisCS) {
		thisString := ""
		for _, t := range thisCS.Intervals {
			thisString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		expectedString := ""
		for _, t := range expectedThisCS.Intervals {
			expectedString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		t.Errorf("Interval1 doesn't contain the expected values: got (%s) vs expected (%s)", thisString, expectedString)
	}

	expectedOtherCS := CharSet{
		[]CharInterval{
			{6, 6},
			{16, 20},
		},
	}

	if !otherCS.EqualTo(expectedOtherCS) {
		thisString := ""
		for _, t := range otherCS.Intervals {
			thisString += fmt.Sprintf("%d-%d", t.Start, t.Stop)
		}

		expectedString := ""
		for _, t := range expectedOtherCS.Intervals {
			expectedString += fmt.Sprintf("%d-%d", t.Start, t.Stop)
		}

		t.Errorf("Interval2 doesn't contain the expected values: got (%s) vs expected (%s)", thisString, expectedString)
	}

	expectedIntersection := CharSet{
		[]CharInterval{
			{3, 5},
			{8, 15},
		},
	}

	if !intersection.EqualTo(expectedIntersection) {
		thisString := ""
		for _, t := range intersection.Intervals {
			thisString += fmt.Sprintf("%d-%d", t.Start, t.Stop)
		}

		expectedString := ""
		for _, t := range expectedIntersection.Intervals {
			expectedString += fmt.Sprintf("%d-%d", t.Start, t.Stop)
		}

		t.Errorf("Intersection doesn't contain the expected values: got (%s) vs expected (%s)", thisString, expectedString)
	}

}

func TestConditions_intersectSkippingCriteria(t *testing.T) {
	thisCS := CharSet{
		Intervals: []CharInterval{
			{0, 5},
		},
	}

	otherCS := CharSet{
		Intervals: []CharInterval{
			{1, 2},
			{4, 5},
		},
	}

	thisCS, otherCS, intersection := thisCS.Intersect(otherCS)

	expectedThisCS := CharSet{
		Intervals: []CharInterval{
			{0, 0},
			{3, 3},
		},
	}

	if !thisCS.EqualTo(expectedThisCS) {
		thisString := ""
		for _, t := range thisCS.Intervals {
			thisString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		expectedString := ""
		for _, t := range expectedThisCS.Intervals {
			expectedString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		fmt.Printf("Interval1 doesn't contain the expected values: got (%s) vs expected (%s)", thisString, expectedString)
		return
	}

	expectedOtherCS := CharSet{
		[]CharInterval{},
	}

	if !otherCS.EqualTo(expectedOtherCS) {
		thisString := ""
		for _, t := range otherCS.Intervals {
			thisString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		expectedString := ""
		for _, t := range expectedOtherCS.Intervals {
			expectedString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		fmt.Printf("Interval2 doesn't contain the expected values: got (%s) vs expected (%s)", thisString, expectedString)
		return
	}

	expectedIntersection := CharSet{
		[]CharInterval{
			{1, 2},
			{4, 5},
		},
	}

	if !intersection.EqualTo(expectedIntersection) {
		thisString := ""
		for _, t := range intersection.Intervals {
			thisString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		expectedString := ""
		for _, t := range expectedIntersection.Intervals {
			expectedString += fmt.Sprintf("%d-%d;", t.Start, t.Stop)
		}

		fmt.Printf("Intersection doesn't contain the expected values: got (%s) vs expected (%s)", thisString, expectedString)
		return
	}
}
