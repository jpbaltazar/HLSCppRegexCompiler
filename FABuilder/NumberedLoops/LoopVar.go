package NumberedLoops

import "fmt"

type VarType uint64

const (
	Bound   VarType = iota // n, m
	Exact                  // n, n
	Unbound                //n, -1
)

type VarId uint64

type LoopVar struct {
	Id VarId

	CaptureGroups []string

	Min int64
	Max int64

	Type VarType
}

func (v *LoopVar) GetLoopVarType() VarType {
	if v.Max == -1 {
		return Unbound
	} else if v.Max == v.Min {
		return Exact
	}
	return Bound
}

func (v *LoopVar) ToString() string {
	c := ""
	switch v.GetLoopVarType() {
	case Exact:
		c = fmt.Sprintf("%d", v.Min)
		break
	case Bound:
		c = fmt.Sprintf("%d, %d", v.Min, v.Max)
		break
	case Unbound:
		c = fmt.Sprintf("%d,", v.Min)
		break
	}

	return fmt.Sprintf("Var{%s}", c)
}
