package NumberedLoops

import "fmt"

type LoopConditionType int

const (
	Complete LoopConditionType = iota
	Incomplete
)

type LoopCondition struct {
	Var  LoopVar
	Type LoopConditionType
}

func (l *LoopCondition) ToString() string {
	str := l.Var.ToString()

	if l.Type == Incomplete {
		str += fmt.Sprintf("<%d", l.Var.Min)
	} else {
		switch l.Var.Type {
		case Exact:
			str += fmt.Sprintf("==%d", l.Var.Min)
			break
		case Bound:
			str += fmt.Sprintf("<=%d", l.Var.Max)
			break
		case Unbound:
			str += fmt.Sprintf(">=%d", l.Var.Min)
			break
		}
	}

	return str
}

type LoopConditions struct {
	C []LoopCondition
}

func NewLoopConditions() *LoopConditions {
	l := new(LoopConditions)
	*l = LoopConditions{
		make([]LoopCondition, 0),
	}

	return l
}

func CopyLoopConditions(conditions LoopConditions) *LoopConditions {
	l := new(LoopConditions)
	*l = LoopConditions{
		make([]LoopCondition, 0),
	}

	l.C = append(l.C, conditions.C...)

	return l
}

func (l *LoopConditions) Add(condition LoopCondition) {
	l.C = append(l.C, condition)
}

func (l *LoopConditions) AddAll(conditions []LoopCondition) {
	l.C = append(l.C, conditions...)
}

func (l *LoopConditions) Size() int {
	return len(l.C)
}

func (l *LoopConditions) Clear() {
	l.C = make([]LoopCondition, 0)
}

func (l *LoopConditions) IsEmpty() bool {
	return len(l.C) == 0
}

func (l *LoopConditions) IsEqual(other *LoopConditions) bool {
	if len(l.C) != len(other.C) {
		return false
	}

	for _, l1 := range l.C {
		found := false
		for _, l2 := range other.C {
			if l1.Type == l2.Type && l1.Var.Id == l2.Var.Id { //the variables should be the exact same
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

func (l *LoopConditions) IsCompatible(other *LoopConditions) bool {
	for _, condition1 := range l.C {
		for _, condition2 := range other.C {
			if condition1.Var.Id == condition2.Var.Id && condition1.Type != condition2.Type {
				return false
			}
		}
	}

	return true
}

// useful for visitor only
func (l *LoopConditions) Evaluate(VarTable map[VarId]int64) bool {
	for _, condition := range l.C {
		//Variable cases
		// (x, y)
		// (x, x)
		// (x,-1) //TODO remove case

		// Condition types
		//		Complete (i < x)
		//		Incomplete (i >= y)
		//			Exception (-1) ->

		if !(condition.Type == Incomplete && VarTable[condition.Var.Id] < condition.Var.Min) &&
			!(condition.Type == Complete && VarTable[condition.Var.Id] >= condition.Var.Max) {
			return false
		}
	}

	return true
}

func (l *LoopConditions) GetConditions() []LoopCondition {
	return append([]LoopCondition{}, l.C...)
}
