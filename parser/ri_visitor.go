// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Ri

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// A complete Visitor for a parse tree produced by RiParser.
type RiVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RiParser#printExpr.
	VisitPrintExpr(ctx *PrintExprContext) interface{}

	// Visit a parse tree produced by RiParser#assgin.
	VisitAssgin(ctx *AssginContext) interface{}

	// Visit a parse tree produced by RiParser#blank.
	VisitBlank(ctx *BlankContext) interface{}

	// Visit a parse tree produced by RiParser#parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by RiParser#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by RiParser#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by RiParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by RiParser#int.
	VisitInt(ctx *IntContext) interface{}

	// Visit a parse tree produced by RiParser#prog.
	VisitProg(ctx *ProgContext) interface{}
}
