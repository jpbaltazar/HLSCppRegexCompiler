package NumberedLoops

import "testing"

func TestLoopInteractions_solve(t *testing.T) {
	v := LoopVar{
		0,
		[]string{},
		0,
		0,
		Bound,
	}

	l := LoopInteractions{
		interactions: []LoopInteraction{
			{&v, AddToVar, 1},
		},
	}

	l.Solve()

	if len(l.interactions) != 1 {
		t.Errorf("Failed to add interaction")
	}

	l.AddInteraction(LoopInteraction{&v, SetVar, 0})

	l.Solve()

	if len(l.interactions) != 1 {
		t.Errorf("Failed to solve interactions")
	}

	if l.interactions[0].Type != SetVar && l.interactions[0].Value == 0 {
		t.Errorf("Solved iteration wrong, expected type %d value %d, got t %d val %d\n",
			SetVar, 0, l.interactions[0].Type, l.interactions[0].Value)
	}
}

func TestLoopInteractions_solve2(t *testing.T) {
	v1 := LoopVar{
		0,
		nil,
		0,
		0,
		Bound,
	}

	v2 := LoopVar{
		1,
		nil,
		0,
		0,
		Bound,
	}

	l := LoopInteractions{
		interactions: []LoopInteraction{
			{&v1, AddToVar, 1},
			{&v2, AddToVar, 3},
		},
	}

	l.Solve()

	if len(l.interactions) != 2 {
		t.Errorf("Merged interactions of different variables")
	}
}
