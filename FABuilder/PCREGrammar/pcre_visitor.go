// Code generated from /home/zed4805/GolandProjects/thesisGoRemake/Resources/Grammar/PCRE.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // PCRE

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by PCREParser.
type PCREVisitor interface {
	antlr.ParseTreeVisitor

	// Walk a parse tree produced by PCREParser#parse.
	VisitParse(ctx *ParseContext) interface{}

	// Walk a parse tree produced by PCREParser#singleRule.
	VisitSingleRule(ctx *SingleRuleContext) interface{}

	// Walk a parse tree produced by PCREParser#directives.
	VisitDirectives(ctx *DirectivesContext) interface{}

	// Walk a parse tree produced by PCREParser#alternation.
	VisitAlternation(ctx *AlternationContext) interface{}

	// Walk a parse tree produced by PCREParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Walk a parse tree produced by PCREParser#element.
	VisitElement(ctx *ElementContext) interface{}

	// Walk a parse tree produced by PCREParser#quantifier.
	VisitQuantifier(ctx *QuantifierContext) interface{}

	// Walk a parse tree produced by PCREParser#quantifier_type.
	VisitQuantifier_type(ctx *Quantifier_typeContext) interface{}

	// Walk a parse tree produced by PCREParser#character_class.
	VisitCharacter_class(ctx *Character_classContext) interface{}

	// Walk a parse tree produced by PCREParser#backreference.
	VisitBackreference(ctx *BackreferenceContext) interface{}

	// Walk a parse tree produced by PCREParser#backreference_or_octal.
	VisitBackreference_or_octal(ctx *Backreference_or_octalContext) interface{}

	// Walk a parse tree produced by PCREParser#capture.
	VisitCapture(ctx *CaptureContext) interface{}

	// Walk a parse tree produced by PCREParser#non_capture.
	VisitNon_capture(ctx *Non_captureContext) interface{}

	// Walk a parse tree produced by PCREParser#comment.
	VisitComment(ctx *CommentContext) interface{}

	// Walk a parse tree produced by PCREParser#option.
	VisitOption(ctx *OptionContext) interface{}

	// Walk a parse tree produced by PCREParser#option_flags.
	VisitOption_flags(ctx *Option_flagsContext) interface{}

	// Walk a parse tree produced by PCREParser#option_flag.
	VisitOption_flag(ctx *Option_flagContext) interface{}

	// Walk a parse tree produced by PCREParser#look_around.
	VisitLook_around(ctx *Look_aroundContext) interface{}

	// Walk a parse tree produced by PCREParser#subroutine_reference.
	VisitSubroutine_reference(ctx *Subroutine_referenceContext) interface{}

	// Walk a parse tree produced by PCREParser#conditional.
	VisitConditional(ctx *ConditionalContext) interface{}

	// Walk a parse tree produced by PCREParser#backtrack_control.
	VisitBacktrack_control(ctx *Backtrack_controlContext) interface{}

	// Walk a parse tree produced by PCREParser#newline_convention.
	VisitNewline_convention(ctx *Newline_conventionContext) interface{}

	// Walk a parse tree produced by PCREParser#callout.
	VisitCallout(ctx *CalloutContext) interface{}

	// Walk a parse tree produced by PCREParser#atom.
	VisitAtom(ctx *AtomContext) interface{}

	// Walk a parse tree produced by PCREParser#cc_atom.
	VisitCc_atom(ctx *Cc_atomContext) interface{}

	// Walk a parse tree produced by PCREParser#shared_atom.
	VisitShared_atom(ctx *Shared_atomContext) interface{}

	// Walk a parse tree produced by PCREParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Walk a parse tree produced by PCREParser#cc_literal.
	VisitCc_literal(ctx *Cc_literalContext) interface{}

	// Walk a parse tree produced by PCREParser#shared_literal.
	VisitShared_literal(ctx *Shared_literalContext) interface{}

	// Walk a parse tree produced by PCREParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Walk a parse tree produced by PCREParser#octal_char.
	VisitOctal_char(ctx *Octal_charContext) interface{}

	// Walk a parse tree produced by PCREParser#octal_digit.
	VisitOctal_digit(ctx *Octal_digitContext) interface{}

	// Walk a parse tree produced by PCREParser#digits.
	VisitDigits(ctx *DigitsContext) interface{}

	// Walk a parse tree produced by PCREParser#digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Walk a parse tree produced by PCREParser#name.
	VisitName(ctx *NameContext) interface{}

	// Walk a parse tree produced by PCREParser#alpha_nums.
	VisitAlpha_nums(ctx *Alpha_numsContext) interface{}

	// Walk a parse tree produced by PCREParser#non_close_parens.
	VisitNon_close_parens(ctx *Non_close_parensContext) interface{}

	// Walk a parse tree produced by PCREParser#non_close_paren.
	VisitNon_close_paren(ctx *Non_close_parenContext) interface{}

	// Walk a parse tree produced by PCREParser#letter.
	VisitLetter(ctx *LetterContext) interface{}
}
