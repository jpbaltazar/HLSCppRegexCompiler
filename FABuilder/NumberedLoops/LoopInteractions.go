package NumberedLoops

import "fmt"

type InteractionType int64

const (
	AddToVar InteractionType = iota
	SetVar
)

type LoopInteraction struct {
	Var *LoopVar

	Type  InteractionType
	Value int64
}

func (l *LoopInteraction) ToString() string {
	str := l.Var.ToString()
	switch l.Type {
	case AddToVar:
		if l.Value == 1 {
			str += "++"
		} else {
			str += fmt.Sprintf("+=%d", l.Value)
		}
	case SetVar:
		str += fmt.Sprintf("=%d", l.Value)
	}

	return str
}

type LoopInteractions struct {
	interactions []LoopInteraction
}

func NewLoopInteractions() *LoopInteractions {
	l := new(LoopInteractions)
	*l = LoopInteractions{
		make([]LoopInteraction, 0),
	}

	return l
}

func CopyLoopInteractions(interactions LoopInteractions) *LoopInteractions {
	l := new(LoopInteractions)
	l.interactions = append(l.interactions, interactions.interactions...)

	return l
}

func (l *LoopInteractions) IsEmpty() bool {
	return l.interactions == nil || len(l.interactions) == 0
}

func (l *LoopInteractions) Solve() {
	m := make(map[VarId]LoopInteraction)

	// TODO improvement:
	//if we start solving from the end, as soon as we hit a 'set'
	//we know that the other operations are useless

	for _, interaction := range l.interactions {
		if entry, ok := m[interaction.Var.Id]; ok {
			if interaction.Type == AddToVar {
				entry.Value += interaction.Value
			} else {
				entry.Value = interaction.Value
				entry.Type = SetVar
			}

			m[interaction.Var.Id] = entry
		} else { //doesn't exist yet
			m[interaction.Var.Id] = interaction
		}
	}

	l.interactions = make([]LoopInteraction, len(m))

	i := 0
	for _, interaction := range m {
		l.interactions[i] = interaction
		i++
	}
}

func (l *LoopInteractions) AddInteraction(i LoopInteraction) {
	l.interactions = append(l.interactions, i)
}

func (l *LoopInteractions) AddInteractions(i LoopInteractions) {
	l.interactions = append(l.interactions, i.interactions...)
}

func (l *LoopInteractions) GetInteractions() []LoopInteraction {
	return append([]LoopInteraction{}, l.interactions...)
}
