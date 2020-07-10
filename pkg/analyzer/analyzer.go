package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "nomaprange",
	Doc:      "checks for range over maps",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.RangeStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		rangeStmt := node.(*ast.RangeStmt)
		what := rangeStmt.X
		t := pass.TypesInfo.TypeOf(what).Underlying()

		if _, ok := t.(*types.Map); ok {
			pass.Reportf(what.Pos(), "range iteration over map")
		}
	})

	return nil, nil
}
