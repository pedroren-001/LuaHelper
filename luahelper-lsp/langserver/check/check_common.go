package check

import (
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/check/common"
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/check/compiler/lexer"
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/check/results"
)

// CommonFuncParam 公共的代码传参
type CommonFuncParam struct {
	fileResult    *results.FileResult
	fi            *common.FuncInfo
	scope         *common.ScopeInfo
	loc           lexer.Location
	secondProject *results.SingleProjectResult
	thirdStruct   *results.AnalysisThird
}

// SliceInsert2 字符串切片拼接
func SliceInsert2(s *[]string, index int, value string) {
	rear := append([]string{}, (*s)[index:]...)
	*s = append(append((*s)[:index], value), rear...)
}
