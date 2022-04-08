package Builder

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"reflect"
	"strings"
	"thesisGoRemake/FABuilder"
	"thesisGoRemake/FABuilder/Builder/CharSet"
	Parsers2 "thesisGoRemake/FABuilder/Builder/CharSet/Parsers"
	"thesisGoRemake/FABuilder/Flags"
	"thesisGoRemake/FABuilder/NumberedLoops"
	parser "thesisGoRemake/FABuilder/PCREGrammar"
	"unicode"
)

type DFABuilderListener struct {
	*parser.BasePCREListener

	Graphs    []Graph
	CurrGraph *Graph

	err error

	Parser Parsers2.Parser

	SubGraphStack SubGraphStack

	captureGroupNames uint64
	CaptureGroupStack *FABuilder.CaptureGroupStack
}

func NewDFABuilderListener() *DFABuilderListener {
	return &DFABuilderListener{
		Graphs:            make([]Graph, 0),
		CurrGraph:         nil,
		Parser:            Parsers2.ASCII{}, //default
		SubGraphStack:     *NewSubGraphStack(),
		CaptureGroupStack: FABuilder.NewCaptureGroupStack(),
	}
}

func GenerateFromFile() {
	input, _ := antlr.NewFileStream("")
	lexer := parser.NewPCRELexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewPCREParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Parse()
	antlr.ParseTreeWalkerDefault.Walk(NewDFABuilderListener(), tree)
}

func (d *DFABuilderListener) EnterParse(ctx *parser.ParseContext) {

}

func (d *DFABuilderListener) EnterSingleRule(ctx *parser.SingleRuleContext) {

	d.CurrGraph = NewGraphWithFlags(ctx.GetText(), Flags.NewFlags(ctx.Directives().GetText()))

	if ctx.Caret() != nil {
		d.CurrGraph.Flags.PCRE_A = true
	}
}

func (d *DFABuilderListener) ExitSingleRule(*parser.SingleRuleContext) {
	//orphan node removal
	//Won't be applied for now
	//TODO implement orphan removal

	sub := d.SubGraphStack.Pop()
	d.CurrGraph.RootVertex = sub.Start

	//loop back
	if d.CurrGraph.Flags.StartOfString {
		errorV := d.CurrGraph.CreateVertex(false, false)
		d.CurrGraph.ErrorVertex = errorV

		restartEdge := d.CurrGraph.CreateEdge(errorV, d.CurrGraph.RootVertex, CharSet.NewCharSet(CharSet.CharInterval{}))
		if d.CurrGraph.Flags.PCRE_G {
			interval, err := d.Parser.ParseCharInterval("\n")
			if err != nil {
				d.err = err
				return
			}
			restartEdge.C.Add(interval)
		}
	} else {
		for _, v := range d.CurrGraph.GetVertexMap() {
			if v.Id != d.CurrGraph.RootVertex.Id {
				for _, edge := range d.CurrGraph.RootVertex.Outgoing {
					d.CurrGraph.CopyEdge(v, edge.To, *edge)
				}
			}
		}
	}

	//TODO loopback for multiline and non multiline rules

	//Call EnforceDeterminism
	err := d.EnforceDeterminism()
	if err != nil {
		return
	}
	//TODO implement
}

func (d *DFABuilderListener) ExitAlternation(ctx *parser.AlternationContext) {
	startV, stopV := d.SubGraphStack.Pop().Decompose()

	allEpsilonIncoming := true
	for _, edge := range stopV.Incoming {
		if !edge.C.IsEmpty() {
			allEpsilonIncoming = false
			break
		}
	}

	var literalStop, epsilonStop *Vertex

	if allEpsilonIncoming {
		epsilonStop = stopV
	} else {
		literalStop = stopV
	}

	for d.SubGraphStack.Size() > 0 || d.SubGraphStack.Peek().Type == Expr {
		newSubgraph := d.SubGraphStack.Pop()
		newStart := newSubgraph.Start
		newStop := newSubgraph.Stop

		d.CurrGraph.MergeVertexBIntoA(startV, newStart)

		allEpsilonIncoming := true
		for _, edge := range newStop.GetIncoming() {
			if !edge.C.IsEmpty() {
				allEpsilonIncoming = false
				break
			}
		}

		if allEpsilonIncoming {
			if epsilonStop == nil {
				literalStop.Terminal = false
				epsilonStop = newStop
				d.CurrGraph.CreateEpsilonEdge(literalStop, epsilonStop)
			} else {
				d.CurrGraph.MergeVertexBIntoA(epsilonStop, newStop)
			}
		} else {
			if epsilonStop == nil {
				literalStop = newStop
				literalStop.Terminal = false
				d.CurrGraph.CreateEpsilonEdge(literalStop, epsilonStop)
			} else {
				d.CurrGraph.MergeVertexBIntoA(literalStop, newStop)
			}
		}
	}

	var finalStop *Vertex
	if epsilonStop == nil {
		finalStop = literalStop
	} else {
		finalStop = epsilonStop
	}

	d.SubGraphStack.Push(&SubGraph{
		Start: startV,
		Stop:  finalStop,
		Type:  Alternation,
		Ctx:   ctx.GetText(),
	})
}

func (d *DFABuilderListener) EnterExpr(*parser.ExprContext) {
	d.SubGraphStack.Push(&SubGraph{
		Type: SEPARATOR,
	})
}

func (d *DFABuilderListener) ExitExpr(ctx *parser.ExprContext) {
	var start, stop *Vertex

	if d.SubGraphStack.Peek().Type != SEPARATOR {
		subgraph := d.SubGraphStack.Pop()
		start = subgraph.Start
		stop = subgraph.Stop

		for d.SubGraphStack.Size() > 0 &&
			d.SubGraphStack.Peek().Type == Atom {
			subgraph = d.SubGraphStack.Pop()

			newStart := subgraph.Start
			newStop := subgraph.Stop

			allEpsilon := true
			for _, edge := range newStop.GetIncoming() {
				if !edge.C.IsEmpty() {
					allEpsilon = false
					break
				}
			}

			if allEpsilon {
				for _, newStopInEdge := range newStop.GetIncoming() {
					for _, startOutEdge := range start.GetOutgoing() {
						d.CurrGraph.CopyEdge(newStopInEdge.From, startOutEdge.To, *startOutEdge)
					}
				}

				d.CurrGraph.RemoveVertex(newStop)
				d.CurrGraph.RemoveVertex(start)
			} else {
				d.CurrGraph.MergeVertexBIntoA(start, newStop)
			}

			start = newStart
		}
	} else {
		start := d.CurrGraph.CreateDefaultVertex()
		stop := d.CurrGraph.CreateVertex(true, false)

		d.CurrGraph.CreateEpsilonEdge(start, stop)
	}

	if d.SubGraphStack.Size() > 0 && d.SubGraphStack.Peek().Type == SEPARATOR {
		d.SubGraphStack.Pop()
	}

	d.SubGraphStack.Push(&SubGraph{
		Start: start,
		Stop:  stop,
		Type:  Expr,
		Ctx:   ctx.GetText(),
	})
}

func (d *DFABuilderListener) ExitAtom(ctx *parser.AtomContext) {

	//automatically convert literals and others into atoms
	if ctx.GetChild(0).GetChildCount() != 0 {
		d.SubGraphStack.Peek().Type = Atom
	} else {
		start := d.CurrGraph.CreateDefaultVertex()
		stop := d.CurrGraph.CreateTerminalVertex()

		e := d.CurrGraph.CreateEpsilonEdge(start, stop)

		switch ctx.GetText() {
		case ".":
			e.C = d.Parser.CreateNewLineCharSet()
			d.Parser.InvertCharSet(e.C) //inversion of '\n' set
			break
			//TODO Word boundaries
			//case "\\B":
			//break
		case "$":
			stop.EndOfSubject = true
			if d.CurrGraph.Flags.PERL_m {
				e.C = d.Parser.CreateNewLineNullCharSet() //
			} else {
				e.C = d.Parser.CreateNullCharSet()
			}
			break
		}

		d.SubGraphStack.Push(&SubGraph{
			Start: start,
			Stop:  stop,
			Type:  Atom,
			Ctx:   ctx.GetText(),
		})
	}
}

func (d *DFABuilderListener) EnterCapture(ctx *parser.CaptureContext) {
	childCount := ctx.GetChildCount()

	if childCount == 3 { //'(' alternation ')'
		d.CaptureGroupStack.Push(fmt.Sprintf("%d", d.captureGroupNames))
		d.captureGroupNames++
	} else { // '(' '?' '<' name '>' alternation ')'
		d.CaptureGroupStack.Push(ctx.GetChildOfType(0, reflect.TypeOf(parser.NameContext{})).GetText())
	}
}

func (d *DFABuilderListener) ExitCapture(*parser.CaptureContext) {
	d.CaptureGroupStack.Pop()
}

func (d *DFABuilderListener) ExitBackreference(ctx *parser.BackreferenceContext) {
	ref := ""

	//TODO test d
	if !ctx.Backreference_or_octal().IsEmpty() { //backreference_or_octal
		//backslash digit
		ref = strings.TrimPrefix(ctx.GetText(), "\\") //haven't found a better way to do d
	} else { //anything else, can be named groups, etc...
		if !ctx.Number().IsEmpty() {
			//named by number
			ref = ctx.Number().GetText()
		} else if !ctx.Name().IsEmpty() {
			//unnamed
			ref = ctx.Name().GetText()
		}
	}

	start := d.CurrGraph.CreateDefaultVertex()
	stop := d.CurrGraph.CreateDefaultVertex()

	d.CurrGraph.CreateBackrefEdge(start, stop, ref)

	epsilonStop := d.CurrGraph.CreateDefaultVertex()

	d.CurrGraph.CreateEpsilonEdge(start, epsilonStop)
	d.CurrGraph.CreateEpsilonEdge(stop, epsilonStop)

	d.SubGraphStack.Push(&SubGraph{
		Start: start,
		Stop:  epsilonStop,
		Type:  Atom,
		Ctx:   ctx.GetText(),
	})
}

func (d *DFABuilderListener) ExitCharacter_class(ctx *parser.Character_classContext) {
	charStart := d.CurrGraph.CreateDefaultVertex()
	charStop := d.CurrGraph.CreateTerminalVertex()

	e := d.CurrGraph.CreateEpsilonEdge(charStart, charStop)

	stack := NewEdgeStack()
	for d.SubGraphStack.Size() > 0 && d.SubGraphStack.Peek().Type == Cc_atom {
		ccAtomSubgraph := d.SubGraphStack.Pop()
		copyOfEdge := Edge{}
		copyOfEdge = *ccAtomSubgraph.Start.Outgoing[ccAtomSubgraph.Stop.Id]

		d.CurrGraph.RemoveVertex(ccAtomSubgraph.Start)
		d.CurrGraph.RemoveVertex(ccAtomSubgraph.Stop)

		stack.Push(&copyOfEdge)
	}

	if ctx.GetChildCount() > 1 && ctx.GetText()[1] == "^"[0] {
		d.Parser.InvertCharSet(e.C)
	}

	for stack.Size() > 0 {
		e.C.AddMultiple(stack.Pop().C.Intervals)
	}

	d.SubGraphStack.Push(&SubGraph{
		Start: charStart,
		Stop:  charStop,
		Type:  Atom,
		Ctx:   ctx.GetText(),
	})
}

func (d *DFABuilderListener) processCCLiteral(ctx antlr.RuleContext) *SubGraph {
	if ctx.GetChild(0).GetChild(0).GetChildCount() > 0 {
		return d.SubGraphStack.Pop()
	} else {
		ccLiteralStart := d.CurrGraph.CreateDefaultVertex()
		ccLiteralStop := d.CurrGraph.CreateDefaultVertex()

		//can be disregarded as the text is predictable from the parsing
		interval, _ := d.Parser.ParseCharInterval(ctx.GetText())

		d.CurrGraph.CreateEdge(ccLiteralStart, ccLiteralStop, &CharSet.CharSet{
			Intervals: []CharSet.CharInterval{interval},
		})

		return &SubGraph{
			Start: ccLiteralStart,
			Stop:  ccLiteralStop,
			Type:  Cc_atom,
			Ctx:   ctx.GetText(),
		}
	}
}

func (d *DFABuilderListener) ExitCc_atom(ctx *parser.Cc_atomContext) {
	//has either:
	// 1 child:
	//  shared_atom
	//  cc_literal
	//  backreference_or_octal //doesn't happen in snort
	// 3 children:
	//  cc_literal '-' cc_literal

	if ctx.GetChildCount() == 1 {
		d.SubGraphStack.Peek().Type = Cc_atom
	} else if ctx.GetChildCount() == 3 {
		var subgraph1, subgraph2 *SubGraph

		ccLiteral1Ctx := ctx.GetChildOfType(0, reflect.TypeOf(parser.Cc_literalContext{}))
		ccLiteral2Ctx := ctx.GetChildOfType(1, reflect.TypeOf(parser.Cc_literalContext{}))

		var ccLiteral1, ccLiteral2 uint64
		if ctx.GetChild(0).GetChild(0).GetChildCount() > 0 {
			subgraph1 = d.SubGraphStack.Pop()
			//start and stop are the same
			ccLiteral1 = subgraph1.Start.Outgoing[subgraph1.Stop.Id].C.Intervals[0].Start
		} else {
			ccLiteral1, _ = d.Parser.ParseCCLiteral(ccLiteral1Ctx.GetText())
		}

		if ctx.GetChild(2).GetChild(0).GetChildCount() > 0 {
			subgraph2 := d.SubGraphStack.Pop()
			//start and stop are the same
			ccLiteral2 = subgraph2.Start.Outgoing[subgraph2.Stop.Id].C.Intervals[0].Start
		} else {
			ccLiteral2, _ = d.Parser.ParseCCLiteral(ccLiteral2Ctx.GetText())
		}

		newLiteralStart := d.CurrGraph.CreateDefaultVertex()
		newLiteralStop := d.CurrGraph.CreateTerminalVertex()

		if ccLiteral2 < ccLiteral1 {
			tmp := ccLiteral1
			ccLiteral1 = ccLiteral2
			ccLiteral2 = tmp
		}

		newEdge := d.CurrGraph.CreateEdge(newLiteralStart, newLiteralStop, &CharSet.CharSet{
			Intervals: []CharSet.CharInterval{{ccLiteral1, ccLiteral2}},
		})

		if d.CurrGraph.Flags.PERL_i {
			start := ccLiteral1
			stop := ccLiteral2

			minimization := false

			//check if it needs uppercase interval
			if !((start < 'a' && stop < 'a') || (start > 'z' && stop > 'z')) {
				var uppercaseStart, uppercaseStop uint64
				if 'a' > start {
					uppercaseStart = 'a'
				} else {
					uppercaseStart = start
				}

				if 'z' < stop {
					uppercaseStop = 'z'
				} else {
					uppercaseStop = stop
				}

				uppercaseStart = uint64(unicode.ToUpper(rune(uppercaseStart)))
				uppercaseStop = uint64(unicode.ToUpper(rune(uppercaseStop)))

				newEdge.C.Add(CharSet.CharInterval{
					Start: uppercaseStart,
					Stop:  uppercaseStop,
				})

				minimization = true
			}

			//check if it needs lowercase interval
			if !((start < 'A' && stop < 'A') || (start > 'Z' && stop > 'Z')) {
				var lowercaseStart, lowercaseStop uint64
				if 'A' > start {
					lowercaseStart = 'A'
				} else {
					lowercaseStart = start
				}

				if 'Z' < stop {
					lowercaseStop = 'Z'
				} else {
					lowercaseStop = stop
				}

				lowercaseStart = uint64(unicode.ToLower(rune(lowercaseStart)))
				lowercaseStop = uint64(unicode.ToLower(rune(lowercaseStop)))

				newEdge.C.Add(CharSet.CharInterval{
					Start: lowercaseStart,
					Stop:  lowercaseStop,
				})

				minimization = true
			}

			if minimization {
				newEdge.C.Minimize()
			}
		}

		d.CurrGraph.RemoveVertex(subgraph1.Start)
		d.CurrGraph.RemoveVertex(subgraph1.Stop)

		d.CurrGraph.RemoveVertex(subgraph2.Start)
		d.CurrGraph.RemoveVertex(subgraph2.Stop)

		d.SubGraphStack.Push(&SubGraph{
			Start: newLiteralStart,
			Stop:  newLiteralStop,
			Type:  Cc_atom,
			Ctx:   ctx.GetText(),
		})
	}

	//TODO finish
}

func (d *DFABuilderListener) EnterShared_atom(ctx *parser.Shared_atomContext) {
	literalStart := d.CurrGraph.CreateDefaultVertex()
	literalStop := d.CurrGraph.CreateTerminalVertex()

	literal := d.CurrGraph.CreateEpsilonEdge(literalStart, literalStop)

	switch ctx.GetText() {
	case "\\D":
		//the error can be ignored as it won't happen, because the case is tested
		interval, _ := d.Parser.ParseCharInterval("0-9")
		literal.C.Add(interval)
		d.Parser.InvertCharSet(literal.C)
		break
	case "\\d":
		interval, _ := d.Parser.ParseCharInterval("0-9")
		literal.C.Add(interval)
		break
	case "\\N": //not new line
		literal.C = d.Parser.CreateNewLineCharSet()
		d.Parser.InvertCharSet(literal.C)
		break
	case "\\W":
		intervals, _ := d.Parser.ParseCharIntervals([]string{"A-Z", "a-z", "0-9", "_"})
		literal.C.AddMultiple(intervals)

		d.Parser.InvertCharSet(literal.C)
		break
	case "\\w":
		intervals, _ := d.Parser.ParseCharIntervals([]string{"A-Z", "a-z", "0-9", "_"})
		literal.C.AddMultiple(intervals)
		break
	case "\\S":
		intervals, _ := d.Parser.ParseCharIntervals([]string{" ", "\\x0C", "\\x0A", "\\x0D", "\\x09", "\\x08"})
		literal.C.AddMultiple(intervals)

		d.Parser.InvertCharSet(literal.C)
		break
	case "\\s":
		intervals, _ := d.Parser.ParseCharIntervals([]string{" ", "\\x0C", "\\x0A", "\\x0D", "\\x09", "\\x08"})
		literal.C.AddMultiple(intervals)
		break
	}

	d.SubGraphStack.Push(&SubGraph{
		Start: literalStart,
		Stop:  literalStop,
		Type:  Atom,
		Ctx:   ctx.GetText(),
	})
}

func (d *DFABuilderListener) ExitLiteral(ctx *parser.LiteralContext) {
	if ctx.GetChild(0).GetChildCount() == 0 {
		start := d.CurrGraph.CreateDefaultVertex()
		stop := d.CurrGraph.CreateTerminalVertex()

		interval, err := d.Parser.ParseCharInterval(ctx.GetText())
		if err != nil {
			print(fmt.Sprintf("Error! Failed to parse %s", ctx.GetText()))
			return
		}

		d.CurrGraph.CreateEdge(start, stop, CharSet.NewCharSet(interval))

		d.SubGraphStack.Push(&SubGraph{
			Start: start,
			Stop:  stop,
			Type:  Literal,
			Ctx:   ctx.GetText(),
		})
	}
}

func (d *DFABuilderListener) ExitShared_literal(ctx *parser.Shared_literalContext) {
	start := d.CurrGraph.CreateDefaultVertex()
	stop := d.CurrGraph.CreateTerminalVertex()

	e := d.CurrGraph.CreateEpsilonEdge(start, stop)
	//TODO cases are forced
	switch ctx.GetText() {
	case "\\a":
		interval, _ := d.Parser.ParseCharInterval("\\x07")
		e.C.Add(interval)
		break
	case "\\e":
		interval, _ := d.Parser.ParseCharInterval("\\x1B")
		e.C.Add(interval)
		break
	case "\\f":
		interval, _ := d.Parser.ParseCharInterval("\\x0C")
		e.C.Add(interval)
		break
	case "\\n":
		interval, _ := d.Parser.ParseCharInterval("\\x0A")
		e.C.Add(interval)
		break
	case "\\r":
		interval, _ := d.Parser.ParseCharInterval("\\x0D")
		e.C.Add(interval)
		break
	case "\\t":
		interval, _ := d.Parser.ParseCharInterval("\\09")
		e.C.Add(interval)
		break
	default:
		txt := ctx.GetText()
		if txt[0] == "\\"[0] &&
			((txt[1] >= 'A' && txt[1] <= 'Z') ||
				(txt[1] >= 'a' && txt[1] <= 'z') ||
				(txt[1] >= '0' && txt[1] <= '9')) &&
			len(txt) == 2 { //skips the backslash for escaped characters
			txt = txt[1:]
		}

		if d.CurrGraph.Flags.PERL_i && //case-insensitive and is a letter
			((txt[0] >= 'A' && txt[0] <= 'Z') ||
				(txt[0] >= 'a' && txt[0] <= 'z')) {
			interval, err := d.Parser.ParseCharInterval(strings.ToLower(txt))
			if err != nil {
				return
			}

			e.C.Add(interval)

			interval, err = d.Parser.ParseCharInterval(strings.ToUpper(txt))
			if err != nil {
				return
			}

			e.C.Add(interval)
		} else { //covers all other cases
			interval, err := d.Parser.ParseCharInterval(txt)
			if err != nil {
				return
			}

			e.C.Add(interval)
		}
		break
	}

	d.SubGraphStack.Push(&SubGraph{
		Start: start,
		Stop:  stop,
		Type:  Literal,
		Ctx:   ctx.GetText(),
	})
}

//If any
func (d *DFABuilderListener) HandleQuestionMarkQuantifier(start, stop *Vertex, ctx string) {
	newStop := d.CurrGraph.CreateTerminalVertex()
	stop.Terminal = false

	d.CurrGraph.CreateEpsilonEdge(start, newStop)

	d.SubGraphStack.Push(&SubGraph{
		Start: start,
		Stop:  newStop,
		Type:  Atom,
		Ctx:   ctx,
	})

}

//at least one
func (d *DFABuilderListener) HandlePlusQuantifier(start, stop *Vertex, ctx string) {
	//  optimization: instead of an empty transition to the beginning,
	//                just copy the transitions to the next state

	for _, e := range start.Outgoing {
		d.CurrGraph.CopyEdge(stop, e.To, *e)
	}

	d.SubGraphStack.Push(&SubGraph{
		Start: start,
		Stop:  stop,
		Type:  Atom,
		Ctx:   ctx,
	})
}

//any number
func (d *DFABuilderListener) HandleAsteriskQuantifier(start, stop *Vertex, ctx string) {
	skipStop := d.CurrGraph.CreateTerminalVertex()
	stop.Terminal = false

	for _, e := range start.Outgoing {
		d.CurrGraph.CopyEdge(stop, e.To, *e)
	}

	//for it to be possible to skip
	d.CurrGraph.CreateEpsilonEdge(start, skipStop)

	d.SubGraphStack.Push(&SubGraph{
		Start: start,
		Stop:  skipStop,
		Type:  Atom,
		Ctx:   ctx,
	})
}

func (d *DFABuilderListener) HandleNumberedQuantifier(start, stop *Vertex, ctx string) {
	//cases:
	//{n} -> exactly
	//{n,} -> at least
	//{n, m} -> between

	minimum, maximum := d.ParseNumberedQuantifier(ctx)

	loopVar := NumberedLoops.LoopVar{
		Id:            d.CurrGraph.CurrGivenLoopVarId,
		CaptureGroups: append(make([]string, 0), d.CaptureGroupStack.GetAll()...),
		Min:           minimum,
		Max:           maximum,
	}

	d.CurrGraph.CurrGivenLoopVarId++

	for _, out := range start.Outgoing {
		out.LoopInteractions.AddInteraction(NumberedLoops.LoopInteraction{
			Var:   &loopVar,
			Type:  NumberedLoops.AddToVar,
			Value: 1,
		})
	}

	incompleteLoopVertex := d.CurrGraph.CreateDefaultVertex()

	incompleteEdge := d.CurrGraph.CreateEpsilonEdge(stop, incompleteLoopVertex)
	incompleteEdge.LoopConditions.Add(NumberedLoops.LoopCondition{
		Var:  loopVar,
		Type: NumberedLoops.Incomplete,
	})

	completeLoopVertex := d.CurrGraph.CreateDefaultVertex()

	completeEdge := d.CurrGraph.CreateEpsilonEdge(stop, completeLoopVertex)
	completeEdge.LoopConditions.Add(NumberedLoops.LoopCondition{
		Var:  loopVar,
		Type: NumberedLoops.Complete,
	})

	start, stop = d.SuperSetSolver(start)
}

func (d *DFABuilderListener) ExitQuantifier(ctx *parser.QuantifierContext) {
	start, stop := d.SubGraphStack.Pop().Decompose()
	quantifierType, quantifierModifier := d.SeparateTypeModifierFromQuantifier(ctx.GetText())

	if quantifierType == "" { //error
		//TODO add error handling

		return
	}

	if quantifierModifier != "" {
		print("Error: quantifier modifiers other than greedy (implicit modifier) are forbidden, treating as greedy...")
	}

	switch quantifierType {
	case "?":
		d.HandleQuestionMarkQuantifier(start, stop, quantifierType)
		break
	case "+":
		d.HandlePlusQuantifier(start, stop, quantifierType)
		break
	case "*":
		d.HandleAsteriskQuantifier(start, stop, quantifierType)
		break
	case "{":
		d.HandleNumberedQuantifier(start, stop, quantifierType)
		break
	}

}
