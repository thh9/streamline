// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // Ri

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

type BaseRiVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRiVisitor) VisitPrintExpr(ctx *PrintExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitAssgin(ctx *AssginContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitBlank(ctx *BlankContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}
