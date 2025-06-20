package langserver

import (
	"context"
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/check"
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/check/common"
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/log"
	"github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/lspcommon"
	lsp "github.com/pedroren-001/LuaHelper/luahelper-lsp/langserver/protocol"
)

// TextDocumentRename 批量更改名字
func (l *LspServer) TextDocumentRename(ctx context.Context, vs lsp.RenameParams) (edit lsp.WorkspaceEdit, err error) {
	// 判断打开的文件，是否是需要分析的文件
	comResult := l.beginFileRequest(vs.TextDocument.URI, vs.Position)
	if !comResult.result {
		return
	}

	if len(comResult.contents) == 0 || comResult.offset >= len(comResult.contents) {
		return
	}

	project := l.getAllProject()
	varStruct := check.GetVarStruct(comResult.contents, comResult.offset, comResult.pos.Line, comResult.pos.Character)
	if !varStruct.ValidFlag {
		log.Error("TextDocumentRename varStruct.ValidFlag not valid")
		return
	}

	// 去掉前缀后的名字
	referenVecs := project.FindReferences(comResult.strFile, &varStruct, common.CRSRename)
	edit.Changes = map[string][]lsp.TextEdit{}

	for _, referVarInfo := range referenVecs {
		retRange := lspcommon.LocToRange(&referVarInfo.Loc)
		uriStr := string(lspcommon.GetFileDocumentURI(referVarInfo.StrFile))

		if _, ok := edit.Changes[uriStr]; !ok {
			edit.Changes[uriStr] = []lsp.TextEdit{}
		}

		edit.Changes[uriStr] = append(edit.Changes[uriStr],
			lsp.TextEdit{
				Range:   retRange,
				NewText: vs.NewName,
			})
	}
	return
}
