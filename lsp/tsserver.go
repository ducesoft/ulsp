// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated for LSP. DO NOT EDIT.

package lsp

// Code generated from protocol/metaModel.json at ref release/protocol/3.17.4-next.0 (hash 5c6ec4f537f304aa1ad645b5fd2bbb757fc40ed1).
// https://github.com/microsoft/vscode-languageserver-node/blob/release/protocol/3.17.4-next.0/protocol/metaModel.json
// LSP metaData.version = 3.17.0.

import (
	"encoding/json"
	"github.com/ducesoft/ulsp/jsonrpc2"
)

type Server interface {
	Progress(ctx Context, conn *jsonrpc2.Conn, params *ProgressParams) error                                                       // $/progress
	SetTrace(ctx Context, conn *jsonrpc2.Conn, params *SetTraceParams) error                                                       // $/setTrace
	IncomingCalls(ctx Context, conn *jsonrpc2.Conn, params *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error) // callHierarchy/incomingCalls
	OutgoingCalls(ctx Context, conn *jsonrpc2.Conn, params *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error) // callHierarchy/outgoingCalls
	ResolveCodeAction(ctx Context, conn *jsonrpc2.Conn, params *CodeAction) (*CodeAction, error)                                   // codeAction/resolve
	ResolveCodeLens(ctx Context, conn *jsonrpc2.Conn, params *CodeLens) (*CodeLens, error)                                         // codeLens/resolve
	ResolveCompletionItem(ctx Context, conn *jsonrpc2.Conn, params *CompletionItem) (*CompletionItem, error)                       // completionItem/resolve
	ResolveDocumentLink(ctx Context, conn *jsonrpc2.Conn, params *DocumentLink) (*DocumentLink, error)                             // documentLink/resolve
	Exit(ctx Context, conn *jsonrpc2.Conn) error                                                                                   // exit
	Initialize(ctx Context, conn *jsonrpc2.Conn, params *ParamInitialize) (*InitializeResult, error)                               // initialize
	Initialized(ctx Context, conn *jsonrpc2.Conn, params *InitializedParams) error                                                 // initialized
	Resolve(ctx Context, conn *jsonrpc2.Conn, params *InlayHint) (*InlayHint, error)                                               // inlayHint/resolve
	DidChangeNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidChangeNotebookDocumentParams) error                     // notebookDocument/didChange
	DidCloseNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidCloseNotebookDocumentParams) error                       // notebookDocument/didClose
	DidOpenNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidOpenNotebookDocumentParams) error                         // notebookDocument/didOpen
	DidSaveNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidSaveNotebookDocumentParams) error                         // notebookDocument/didSave
	Shutdown(ctx Context, conn *jsonrpc2.Conn) error                                                                               // shutdown
	CodeAction(ctx Context, conn *jsonrpc2.Conn, params *CodeActionParams) ([]CodeAction, error)                                   // textDocument/codeAction
	CodeLens(ctx Context, conn *jsonrpc2.Conn, params *CodeLensParams) ([]CodeLens, error)                                         // textDocument/codeLens
	ColorPresentation(ctx Context, conn *jsonrpc2.Conn, params *ColorPresentationParams) ([]ColorPresentation, error)              // textDocument/colorPresentation
	Completion(ctx Context, conn *jsonrpc2.Conn, params *CompletionParams) (*CompletionList, error)                                // textDocument/completion
	Declaration(ctx Context, conn *jsonrpc2.Conn, params *DeclarationParams) (*Or_textDocument_declaration, error)                 // textDocument/declaration
	Definition(ctx Context, conn *jsonrpc2.Conn, params *DefinitionParams) ([]Location, error)                                     // textDocument/definition
	Diagnostic(ctx Context, conn *jsonrpc2.Conn, params *string) (*string, error)                                                  // textDocument/diagnostic
	DidChange(ctx Context, conn *jsonrpc2.Conn, params *DidChangeTextDocumentParams) error                                         // textDocument/didChange
	DidClose(ctx Context, conn *jsonrpc2.Conn, params *DidCloseTextDocumentParams) error                                           // textDocument/didClose
	DidOpen(ctx Context, conn *jsonrpc2.Conn, params *DidOpenTextDocumentParams) error                                             // textDocument/didOpen
	DidSave(ctx Context, conn *jsonrpc2.Conn, params *DidSaveTextDocumentParams) error                                             // textDocument/didSave
	DocumentColor(ctx Context, conn *jsonrpc2.Conn, params *DocumentColorParams) ([]ColorInformation, error)                       // textDocument/documentColor
	DocumentHighlight(ctx Context, conn *jsonrpc2.Conn, params *DocumentHighlightParams) ([]DocumentHighlight, error)              // textDocument/documentHighlight
	DocumentLink(ctx Context, conn *jsonrpc2.Conn, params *DocumentLinkParams) ([]DocumentLink, error)                             // textDocument/documentLink
	DocumentSymbol(ctx Context, conn *jsonrpc2.Conn, params *DocumentSymbolParams) ([]interface{}, error)                          // textDocument/documentSymbol
	FoldingRange(ctx Context, conn *jsonrpc2.Conn, params *FoldingRangeParams) ([]FoldingRange, error)                             // textDocument/foldingRange
	Formatting(ctx Context, conn *jsonrpc2.Conn, params *DocumentFormattingParams) ([]TextEdit, error)                             // textDocument/formatting
	Hover(ctx Context, conn *jsonrpc2.Conn, params *HoverParams) (*Hover, error)                                                   // textDocument/hover
	Implementation(ctx Context, conn *jsonrpc2.Conn, params *ImplementationParams) ([]Location, error)                             // textDocument/implementation
	InlayHint(ctx Context, conn *jsonrpc2.Conn, params *InlayHintParams) ([]InlayHint, error)                                      // textDocument/inlayHint
	InlineValue(ctx Context, conn *jsonrpc2.Conn, params *InlineValueParams) ([]InlineValue, error)                                // textDocument/inlineValue
	LinkedEditingRange(ctx Context, conn *jsonrpc2.Conn, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error)           // textDocument/linkedEditingRange
	Moniker(ctx Context, conn *jsonrpc2.Conn, params *MonikerParams) ([]Moniker, error)                                            // textDocument/moniker
	OnTypeFormatting(ctx Context, conn *jsonrpc2.Conn, params *DocumentOnTypeFormattingParams) ([]TextEdit, error)                 // textDocument/onTypeFormatting
	PrepareCallHierarchy(ctx Context, conn *jsonrpc2.Conn, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error)        // textDocument/prepareCallHierarchy
	PrepareRename(ctx Context, conn *jsonrpc2.Conn, params *PrepareRenameParams) (*PrepareRename2Gn, error)                        // textDocument/prepareRename
	PrepareTypeHierarchy(ctx Context, conn *jsonrpc2.Conn, params *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error)        // textDocument/prepareTypeHierarchy
	RangeFormatting(ctx Context, conn *jsonrpc2.Conn, params *DocumentRangeFormattingParams) ([]TextEdit, error)                   // textDocument/rangeFormatting
	References(ctx Context, conn *jsonrpc2.Conn, params *ReferenceParams) ([]Location, error)                                      // textDocument/references
	Rename(ctx Context, conn *jsonrpc2.Conn, params *RenameParams) (*WorkspaceEdit, error)                                         // textDocument/rename
	SelectionRange(ctx Context, conn *jsonrpc2.Conn, params *SelectionRangeParams) ([]SelectionRange, error)                       // textDocument/selectionRange
	SemanticTokensFull(ctx Context, conn *jsonrpc2.Conn, params *SemanticTokensParams) (*SemanticTokens, error)                    // textDocument/semanticTokens/full
	SemanticTokensFullDelta(ctx Context, conn *jsonrpc2.Conn, params *SemanticTokensDeltaParams) (interface{}, error)              // textDocument/semanticTokens/full/delta
	SemanticTokensRange(ctx Context, conn *jsonrpc2.Conn, params *SemanticTokensRangeParams) (*SemanticTokens, error)              // textDocument/semanticTokens/range
	SignatureHelp(ctx Context, conn *jsonrpc2.Conn, params *SignatureHelpParams) (*SignatureHelp, error)                           // textDocument/signatureHelp
	TypeDefinition(ctx Context, conn *jsonrpc2.Conn, params *TypeDefinitionParams) ([]Location, error)                             // textDocument/typeDefinition
	WillSave(ctx Context, conn *jsonrpc2.Conn, params *WillSaveTextDocumentParams) error                                           // textDocument/willSave
	WillSaveWaitUntil(ctx Context, conn *jsonrpc2.Conn, params *WillSaveTextDocumentParams) ([]TextEdit, error)                    // textDocument/willSaveWaitUntil
	Subtypes(ctx Context, conn *jsonrpc2.Conn, params *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error)                   // typeHierarchy/subtypes
	Supertypes(ctx Context, conn *jsonrpc2.Conn, params *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error)               // typeHierarchy/supertypes
	WorkDoneProgressCancel(ctx Context, conn *jsonrpc2.Conn, params *WorkDoneProgressCancelParams) error                           // window/workDoneProgress/cancel
	DiagnosticWorkspace(ctx Context, conn *jsonrpc2.Conn, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error)   // workspace/diagnostic
	DidChangeConfiguration(ctx Context, conn *jsonrpc2.Conn, params *DidChangeConfigurationParams) error                           // workspace/didChangeConfiguration
	DidChangeWatchedFiles(ctx Context, conn *jsonrpc2.Conn, params *DidChangeWatchedFilesParams) error                             // workspace/didChangeWatchedFiles
	DidChangeWorkspaceFolders(ctx Context, conn *jsonrpc2.Conn, params *DidChangeWorkspaceFoldersParams) error                     // workspace/didChangeWorkspaceFolders
	DidCreateFiles(ctx Context, conn *jsonrpc2.Conn, params *CreateFilesParams) error                                              // workspace/didCreateFiles
	DidDeleteFiles(ctx Context, conn *jsonrpc2.Conn, params *DeleteFilesParams) error                                              // workspace/didDeleteFiles
	DidRenameFiles(ctx Context, conn *jsonrpc2.Conn, params *RenameFilesParams) error                                              // workspace/didRenameFiles
	ExecuteCommand(ctx Context, conn *jsonrpc2.Conn, params *ExecuteCommandParams) (interface{}, error)                            // workspace/executeCommand
	Symbol(ctx Context, conn *jsonrpc2.Conn, params *WorkspaceSymbolParams) ([]SymbolInformation, error)                           // workspace/symbol
	WillCreateFiles(ctx Context, conn *jsonrpc2.Conn, params *CreateFilesParams) (*WorkspaceEdit, error)                           // workspace/willCreateFiles
	WillDeleteFiles(ctx Context, conn *jsonrpc2.Conn, params *DeleteFilesParams) (*WorkspaceEdit, error)                           // workspace/willDeleteFiles
	WillRenameFiles(ctx Context, conn *jsonrpc2.Conn, params *RenameFilesParams) (*WorkspaceEdit, error)                           // workspace/willRenameFiles
	ResolveWorkspaceSymbol(ctx Context, conn *jsonrpc2.Conn, params *WorkspaceSymbol) (*WorkspaceSymbol, error)                    // workspaceSymbol/resolve
	NonstandardRequest(ctx Context, method string, params interface{}) (interface{}, error)
}

func serverDispatch(ctx Context, server Server, conn *jsonrpc2.Conn, r Request) (any, error) {
	switch r.Method() {
	case "$/progress":
		var params ProgressParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.Progress(ctx, conn, &params)
		return nil, err
	case "$/setTrace":
		var params SetTraceParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.SetTrace(ctx, conn, &params)
		return nil, err
	case "callHierarchy/incomingCalls":
		var params CallHierarchyIncomingCallsParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.IncomingCalls(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "callHierarchy/outgoingCalls":
		var params CallHierarchyOutgoingCallsParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.OutgoingCalls(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "codeAction/resolve":
		var params CodeAction
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ResolveCodeAction(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "codeLens/resolve":
		var params CodeLens
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ResolveCodeLens(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "completionItem/resolve":
		var params CompletionItem
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ResolveCompletionItem(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "documentLink/resolve":
		var params DocumentLink
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ResolveDocumentLink(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "exit":
		err := server.Exit(ctx, conn)
		return nil, err
	case "initialize":
		var params ParamInitialize
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Initialize(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "initialized":
		var params InitializedParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.Initialized(ctx, conn, &params)
		return nil, err
	case "inlayHint/resolve":
		var params InlayHint
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Resolve(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "notebookDocument/didChange":
		var params DidChangeNotebookDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidChangeNotebookDocument(ctx, conn, &params)
		return nil, err
	case "notebookDocument/didClose":
		var params DidCloseNotebookDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidCloseNotebookDocument(ctx, conn, &params)
		return nil, err
	case "notebookDocument/didOpen":
		var params DidOpenNotebookDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidOpenNotebookDocument(ctx, conn, &params)
		return nil, err
	case "notebookDocument/didSave":
		var params DidSaveNotebookDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidSaveNotebookDocument(ctx, conn, &params)
		return nil, err
	case "shutdown":
		err := server.Shutdown(ctx, conn)
		return nil, err
	case "textDocument/codeAction":
		var params CodeActionParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.CodeAction(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/codeLens":
		var params CodeLensParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.CodeLens(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/colorPresentation":
		var params ColorPresentationParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ColorPresentation(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/completion":
		var params CompletionParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Completion(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/declaration":
		var params DeclarationParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Declaration(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/definition":
		var params DefinitionParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Definition(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/diagnostic":
		var params string
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Diagnostic(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/didChange":
		var params DidChangeTextDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidChange(ctx, conn, &params)
		return nil, err
	case "textDocument/didClose":
		var params DidCloseTextDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidClose(ctx, conn, &params)
		return nil, err
	case "textDocument/didOpen":
		var params DidOpenTextDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidOpen(ctx, conn, &params)
		return nil, err
	case "textDocument/didSave":
		var params DidSaveTextDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidSave(ctx, conn, &params)
		return nil, err
	case "textDocument/documentColor":
		var params DocumentColorParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.DocumentColor(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/documentHighlight":
		var params DocumentHighlightParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.DocumentHighlight(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/documentLink":
		var params DocumentLinkParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.DocumentLink(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/documentSymbol":
		var params DocumentSymbolParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.DocumentSymbol(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/foldingRange":
		var params FoldingRangeParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.FoldingRange(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/formatting":
		var params DocumentFormattingParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Formatting(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/hover":
		var params HoverParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Hover(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/implementation":
		var params ImplementationParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Implementation(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/inlayHint":
		var params InlayHintParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.InlayHint(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/inlineValue":
		var params InlineValueParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.InlineValue(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/linkedEditingRange":
		var params LinkedEditingRangeParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.LinkedEditingRange(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/moniker":
		var params MonikerParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Moniker(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/onTypeFormatting":
		var params DocumentOnTypeFormattingParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.OnTypeFormatting(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/prepareCallHierarchy":
		var params CallHierarchyPrepareParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.PrepareCallHierarchy(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/prepareRename":
		var params PrepareRenameParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.PrepareRename(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/prepareTypeHierarchy":
		var params TypeHierarchyPrepareParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.PrepareTypeHierarchy(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/rangeFormatting":
		var params DocumentRangeFormattingParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.RangeFormatting(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/references":
		var params ReferenceParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.References(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/rename":
		var params RenameParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Rename(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/selectionRange":
		var params SelectionRangeParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.SelectionRange(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/semanticTokens/full":
		var params SemanticTokensParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.SemanticTokensFull(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/semanticTokens/full/delta":
		var params SemanticTokensDeltaParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.SemanticTokensFullDelta(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/semanticTokens/range":
		var params SemanticTokensRangeParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.SemanticTokensRange(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/signatureHelp":
		var params SignatureHelpParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.SignatureHelp(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/typeDefinition":
		var params TypeDefinitionParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.TypeDefinition(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "textDocument/willSave":
		var params WillSaveTextDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.WillSave(ctx, conn, &params)
		return nil, err
	case "textDocument/willSaveWaitUntil":
		var params WillSaveTextDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.WillSaveWaitUntil(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "typeHierarchy/subtypes":
		var params TypeHierarchySubtypesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Subtypes(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "typeHierarchy/supertypes":
		var params TypeHierarchySupertypesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Supertypes(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "window/workDoneProgress/cancel":
		var params WorkDoneProgressCancelParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.WorkDoneProgressCancel(ctx, conn, &params)
		return nil, err
	case "workspace/diagnostic":
		var params WorkspaceDiagnosticParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.DiagnosticWorkspace(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/didChangeConfiguration":
		var params DidChangeConfigurationParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidChangeConfiguration(ctx, conn, &params)
		return nil, err
	case "workspace/didChangeWatchedFiles":
		var params DidChangeWatchedFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidChangeWatchedFiles(ctx, conn, &params)
		return nil, err
	case "workspace/didChangeWorkspaceFolders":
		var params DidChangeWorkspaceFoldersParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidChangeWorkspaceFolders(ctx, conn, &params)
		return nil, err
	case "workspace/didCreateFiles":
		var params CreateFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidCreateFiles(ctx, conn, &params)
		return nil, err
	case "workspace/didDeleteFiles":
		var params DeleteFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidDeleteFiles(ctx, conn, &params)
		return nil, err
	case "workspace/didRenameFiles":
		var params RenameFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := server.DidRenameFiles(ctx, conn, &params)
		return nil, err
	case "workspace/executeCommand":
		var params ExecuteCommandParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ExecuteCommand(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/symbol":
		var params WorkspaceSymbolParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.Symbol(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/willCreateFiles":
		var params CreateFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.WillCreateFiles(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/willDeleteFiles":
		var params DeleteFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.WillDeleteFiles(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/willRenameFiles":
		var params RenameFilesParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.WillRenameFiles(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspaceSymbol/resolve":
		var params WorkspaceSymbol
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := server.ResolveWorkspaceSymbol(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, ErrMethodNotFound
	}
}

func (s *serverDispatcher) Progress(ctx Context, params *ProgressParams) error {
	return s.sender.Notify(ctx, "$/progress", params)
}
func (s *serverDispatcher) SetTrace(ctx Context, params *SetTraceParams) error {
	return s.sender.Notify(ctx, "$/setTrace", params)
}
func (s *serverDispatcher) IncomingCalls(ctx Context, params *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error) {
	var result []CallHierarchyIncomingCall
	if err := s.sender.Call(ctx, "callHierarchy/incomingCalls", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) OutgoingCalls(ctx Context, params *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error) {
	var result []CallHierarchyOutgoingCall
	if err := s.sender.Call(ctx, "callHierarchy/outgoingCalls", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) ResolveCodeAction(ctx Context, params *CodeAction) (*CodeAction, error) {
	var result *CodeAction
	if err := s.sender.Call(ctx, "codeAction/resolve", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) ResolveCodeLens(ctx Context, params *CodeLens) (*CodeLens, error) {
	var result *CodeLens
	if err := s.sender.Call(ctx, "codeLens/resolve", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) ResolveCompletionItem(ctx Context, params *CompletionItem) (*CompletionItem, error) {
	var result *CompletionItem
	if err := s.sender.Call(ctx, "completionItem/resolve", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) ResolveDocumentLink(ctx Context, params *DocumentLink) (*DocumentLink, error) {
	var result *DocumentLink
	if err := s.sender.Call(ctx, "documentLink/resolve", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Exit(ctx Context) error {
	return s.sender.Notify(ctx, "exit", nil)
}
func (s *serverDispatcher) Initialize(ctx Context, params *ParamInitialize) (*InitializeResult, error) {
	var result *InitializeResult
	if err := s.sender.Call(ctx, "initialize", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Initialized(ctx Context, params *InitializedParams) error {
	return s.sender.Notify(ctx, "initialized", params)
}
func (s *serverDispatcher) Resolve(ctx Context, params *InlayHint) (*InlayHint, error) {
	var result *InlayHint
	if err := s.sender.Call(ctx, "inlayHint/resolve", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) DidChangeNotebookDocument(ctx Context, params *DidChangeNotebookDocumentParams) error {
	return s.sender.Notify(ctx, "notebookDocument/didChange", params)
}
func (s *serverDispatcher) DidCloseNotebookDocument(ctx Context, params *DidCloseNotebookDocumentParams) error {
	return s.sender.Notify(ctx, "notebookDocument/didClose", params)
}
func (s *serverDispatcher) DidOpenNotebookDocument(ctx Context, params *DidOpenNotebookDocumentParams) error {
	return s.sender.Notify(ctx, "notebookDocument/didOpen", params)
}
func (s *serverDispatcher) DidSaveNotebookDocument(ctx Context, params *DidSaveNotebookDocumentParams) error {
	return s.sender.Notify(ctx, "notebookDocument/didSave", params)
}
func (s *serverDispatcher) Shutdown(ctx Context) error {
	return s.sender.Call(ctx, "shutdown", nil, nil)
}
func (s *serverDispatcher) CodeAction(ctx Context, params *CodeActionParams) ([]CodeAction, error) {
	var result []CodeAction
	if err := s.sender.Call(ctx, "textDocument/codeAction", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) CodeLens(ctx Context, params *CodeLensParams) ([]CodeLens, error) {
	var result []CodeLens
	if err := s.sender.Call(ctx, "textDocument/codeLens", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) ColorPresentation(ctx Context, params *ColorPresentationParams) ([]ColorPresentation, error) {
	var result []ColorPresentation
	if err := s.sender.Call(ctx, "textDocument/colorPresentation", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Completion(ctx Context, params *CompletionParams) (*CompletionList, error) {
	var result *CompletionList
	if err := s.sender.Call(ctx, "textDocument/completion", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Declaration(ctx Context, params *DeclarationParams) (*Or_textDocument_declaration, error) {
	var result *Or_textDocument_declaration
	if err := s.sender.Call(ctx, "textDocument/declaration", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Definition(ctx Context, params *DefinitionParams) ([]Location, error) {
	var result []Location
	if err := s.sender.Call(ctx, "textDocument/definition", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Diagnostic(ctx Context, params *string) (*string, error) {
	var result *string
	if err := s.sender.Call(ctx, "textDocument/diagnostic", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) DidChange(ctx Context, params *DidChangeTextDocumentParams) error {
	return s.sender.Notify(ctx, "textDocument/didChange", params)
}
func (s *serverDispatcher) DidClose(ctx Context, params *DidCloseTextDocumentParams) error {
	return s.sender.Notify(ctx, "textDocument/didClose", params)
}
func (s *serverDispatcher) DidOpen(ctx Context, params *DidOpenTextDocumentParams) error {
	return s.sender.Notify(ctx, "textDocument/didOpen", params)
}
func (s *serverDispatcher) DidSave(ctx Context, params *DidSaveTextDocumentParams) error {
	return s.sender.Notify(ctx, "textDocument/didSave", params)
}
func (s *serverDispatcher) DocumentColor(ctx Context, params *DocumentColorParams) ([]ColorInformation, error) {
	var result []ColorInformation
	if err := s.sender.Call(ctx, "textDocument/documentColor", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) DocumentHighlight(ctx Context, params *DocumentHighlightParams) ([]DocumentHighlight, error) {
	var result []DocumentHighlight
	if err := s.sender.Call(ctx, "textDocument/documentHighlight", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) DocumentLink(ctx Context, params *DocumentLinkParams) ([]DocumentLink, error) {
	var result []DocumentLink
	if err := s.sender.Call(ctx, "textDocument/documentLink", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) DocumentSymbol(ctx Context, params *DocumentSymbolParams) ([]interface{}, error) {
	var result []interface{}
	if err := s.sender.Call(ctx, "textDocument/documentSymbol", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) FoldingRange(ctx Context, params *FoldingRangeParams) ([]FoldingRange, error) {
	var result []FoldingRange
	if err := s.sender.Call(ctx, "textDocument/foldingRange", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Formatting(ctx Context, params *DocumentFormattingParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.sender.Call(ctx, "textDocument/formatting", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Hover(ctx Context, params *HoverParams) (*Hover, error) {
	var result *Hover
	if err := s.sender.Call(ctx, "textDocument/hover", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Implementation(ctx Context, params *ImplementationParams) ([]Location, error) {
	var result []Location
	if err := s.sender.Call(ctx, "textDocument/implementation", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) InlayHint(ctx Context, params *InlayHintParams) ([]InlayHint, error) {
	var result []InlayHint
	if err := s.sender.Call(ctx, "textDocument/inlayHint", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) InlineValue(ctx Context, params *InlineValueParams) ([]InlineValue, error) {
	var result []InlineValue
	if err := s.sender.Call(ctx, "textDocument/inlineValue", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) LinkedEditingRange(ctx Context, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
	var result *LinkedEditingRanges
	if err := s.sender.Call(ctx, "textDocument/linkedEditingRange", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Moniker(ctx Context, params *MonikerParams) ([]Moniker, error) {
	var result []Moniker
	if err := s.sender.Call(ctx, "textDocument/moniker", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) OnTypeFormatting(ctx Context, params *DocumentOnTypeFormattingParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.sender.Call(ctx, "textDocument/onTypeFormatting", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) PrepareCallHierarchy(ctx Context, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error) {
	var result []CallHierarchyItem
	if err := s.sender.Call(ctx, "textDocument/prepareCallHierarchy", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) PrepareRename(ctx Context, params *PrepareRenameParams) (*PrepareRename2Gn, error) {
	var result *PrepareRename2Gn
	if err := s.sender.Call(ctx, "textDocument/prepareRename", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) PrepareTypeHierarchy(ctx Context, params *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error) {
	var result []TypeHierarchyItem
	if err := s.sender.Call(ctx, "textDocument/prepareTypeHierarchy", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) RangeFormatting(ctx Context, params *DocumentRangeFormattingParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.sender.Call(ctx, "textDocument/rangeFormatting", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) References(ctx Context, params *ReferenceParams) ([]Location, error) {
	var result []Location
	if err := s.sender.Call(ctx, "textDocument/references", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Rename(ctx Context, params *RenameParams) (*WorkspaceEdit, error) {
	var result *WorkspaceEdit
	if err := s.sender.Call(ctx, "textDocument/rename", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) SelectionRange(ctx Context, params *SelectionRangeParams) ([]SelectionRange, error) {
	var result []SelectionRange
	if err := s.sender.Call(ctx, "textDocument/selectionRange", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) SemanticTokensFull(ctx Context, params *SemanticTokensParams) (*SemanticTokens, error) {
	var result *SemanticTokens
	if err := s.sender.Call(ctx, "textDocument/semanticTokens/full", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) SemanticTokensFullDelta(ctx Context, params *SemanticTokensDeltaParams) (interface{}, error) {
	var result interface{}
	if err := s.sender.Call(ctx, "textDocument/semanticTokens/full/delta", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) SemanticTokensRange(ctx Context, params *SemanticTokensRangeParams) (*SemanticTokens, error) {
	var result *SemanticTokens
	if err := s.sender.Call(ctx, "textDocument/semanticTokens/range", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) SignatureHelp(ctx Context, params *SignatureHelpParams) (*SignatureHelp, error) {
	var result *SignatureHelp
	if err := s.sender.Call(ctx, "textDocument/signatureHelp", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) TypeDefinition(ctx Context, params *TypeDefinitionParams) ([]Location, error) {
	var result []Location
	if err := s.sender.Call(ctx, "textDocument/typeDefinition", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) WillSave(ctx Context, params *WillSaveTextDocumentParams) error {
	return s.sender.Notify(ctx, "textDocument/willSave", params)
}
func (s *serverDispatcher) WillSaveWaitUntil(ctx Context, params *WillSaveTextDocumentParams) ([]TextEdit, error) {
	var result []TextEdit
	if err := s.sender.Call(ctx, "textDocument/willSaveWaitUntil", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Subtypes(ctx Context, params *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error) {
	var result []TypeHierarchyItem
	if err := s.sender.Call(ctx, "typeHierarchy/subtypes", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Supertypes(ctx Context, params *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error) {
	var result []TypeHierarchyItem
	if err := s.sender.Call(ctx, "typeHierarchy/supertypes", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) WorkDoneProgressCancel(ctx Context, params *WorkDoneProgressCancelParams) error {
	return s.sender.Notify(ctx, "window/workDoneProgress/cancel", params)
}
func (s *serverDispatcher) DiagnosticWorkspace(ctx Context, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error) {
	var result *WorkspaceDiagnosticReport
	if err := s.sender.Call(ctx, "workspace/diagnostic", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) DidChangeConfiguration(ctx Context, params *DidChangeConfigurationParams) error {
	return s.sender.Notify(ctx, "workspace/didChangeConfiguration", params)
}
func (s *serverDispatcher) DidChangeWatchedFiles(ctx Context, params *DidChangeWatchedFilesParams) error {
	return s.sender.Notify(ctx, "workspace/didChangeWatchedFiles", params)
}
func (s *serverDispatcher) DidChangeWorkspaceFolders(ctx Context, params *DidChangeWorkspaceFoldersParams) error {
	return s.sender.Notify(ctx, "workspace/didChangeWorkspaceFolders", params)
}
func (s *serverDispatcher) DidCreateFiles(ctx Context, params *CreateFilesParams) error {
	return s.sender.Notify(ctx, "workspace/didCreateFiles", params)
}
func (s *serverDispatcher) DidDeleteFiles(ctx Context, params *DeleteFilesParams) error {
	return s.sender.Notify(ctx, "workspace/didDeleteFiles", params)
}
func (s *serverDispatcher) DidRenameFiles(ctx Context, params *RenameFilesParams) error {
	return s.sender.Notify(ctx, "workspace/didRenameFiles", params)
}
func (s *serverDispatcher) ExecuteCommand(ctx Context, params *ExecuteCommandParams) (interface{}, error) {
	var result interface{}
	if err := s.sender.Call(ctx, "workspace/executeCommand", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) Symbol(ctx Context, params *WorkspaceSymbolParams) ([]SymbolInformation, error) {
	var result []SymbolInformation
	if err := s.sender.Call(ctx, "workspace/symbol", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) WillCreateFiles(ctx Context, params *CreateFilesParams) (*WorkspaceEdit, error) {
	var result *WorkspaceEdit
	if err := s.sender.Call(ctx, "workspace/willCreateFiles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) WillDeleteFiles(ctx Context, params *DeleteFilesParams) (*WorkspaceEdit, error) {
	var result *WorkspaceEdit
	if err := s.sender.Call(ctx, "workspace/willDeleteFiles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) WillRenameFiles(ctx Context, params *RenameFilesParams) (*WorkspaceEdit, error) {
	var result *WorkspaceEdit
	if err := s.sender.Call(ctx, "workspace/willRenameFiles", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) ResolveWorkspaceSymbol(ctx Context, params *WorkspaceSymbol) (*WorkspaceSymbol, error) {
	var result *WorkspaceSymbol
	if err := s.sender.Call(ctx, "workspaceSymbol/resolve", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *serverDispatcher) NonstandardRequest(ctx Context, method string, params interface{}) (interface{}, error) {
	var result interface{}
	if err := s.sender.Call(ctx, method, params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

type AbsServer struct {
}

func (that *AbsServer) Progress(ctx Context, conn *jsonrpc2.Conn, params *ProgressParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) SetTrace(ctx Context, conn *jsonrpc2.Conn, params *SetTraceParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) IncomingCalls(ctx Context, conn *jsonrpc2.Conn, params *CallHierarchyIncomingCallsParams) ([]CallHierarchyIncomingCall, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) OutgoingCalls(ctx Context, conn *jsonrpc2.Conn, params *CallHierarchyOutgoingCallsParams) ([]CallHierarchyOutgoingCall, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) ResolveCodeAction(ctx Context, conn *jsonrpc2.Conn, params *CodeAction) (*CodeAction, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) ResolveCodeLens(ctx Context, conn *jsonrpc2.Conn, params *CodeLens) (*CodeLens, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) ResolveCompletionItem(ctx Context, conn *jsonrpc2.Conn, params *CompletionItem) (*CompletionItem, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) ResolveDocumentLink(ctx Context, conn *jsonrpc2.Conn, params *DocumentLink) (*DocumentLink, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Exit(ctx Context, conn *jsonrpc2.Conn) error {
	return ErrMethodNotFound
}

func (that *AbsServer) Initialize(ctx Context, conn *jsonrpc2.Conn, params *ParamInitialize) (*InitializeResult, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Initialized(ctx Context, conn *jsonrpc2.Conn, params *InitializedParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) Resolve(ctx Context, conn *jsonrpc2.Conn, params *InlayHint) (*InlayHint, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) DidChangeNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidChangeNotebookDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidCloseNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidCloseNotebookDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidOpenNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidOpenNotebookDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidSaveNotebookDocument(ctx Context, conn *jsonrpc2.Conn, params *DidSaveNotebookDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) Shutdown(ctx Context, conn *jsonrpc2.Conn) error {
	return ErrMethodNotFound
}

func (that *AbsServer) CodeAction(ctx Context, conn *jsonrpc2.Conn, params *CodeActionParams) ([]CodeAction, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) CodeLens(ctx Context, conn *jsonrpc2.Conn, params *CodeLensParams) ([]CodeLens, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) ColorPresentation(ctx Context, conn *jsonrpc2.Conn, params *ColorPresentationParams) ([]ColorPresentation, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Completion(ctx Context, conn *jsonrpc2.Conn, params *CompletionParams) (*CompletionList, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Declaration(ctx Context, conn *jsonrpc2.Conn, params *DeclarationParams) (*Or_textDocument_declaration, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Definition(ctx Context, conn *jsonrpc2.Conn, params *DefinitionParams) ([]Location, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Diagnostic(ctx Context, conn *jsonrpc2.Conn, params *string) (*string, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) DidChange(ctx Context, conn *jsonrpc2.Conn, params *DidChangeTextDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidClose(ctx Context, conn *jsonrpc2.Conn, params *DidCloseTextDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidOpen(ctx Context, conn *jsonrpc2.Conn, params *DidOpenTextDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidSave(ctx Context, conn *jsonrpc2.Conn, params *DidSaveTextDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DocumentColor(ctx Context, conn *jsonrpc2.Conn, params *DocumentColorParams) ([]ColorInformation, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) DocumentHighlight(ctx Context, conn *jsonrpc2.Conn, params *DocumentHighlightParams) ([]DocumentHighlight, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) DocumentLink(ctx Context, conn *jsonrpc2.Conn, params *DocumentLinkParams) ([]DocumentLink, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) DocumentSymbol(ctx Context, conn *jsonrpc2.Conn, params *DocumentSymbolParams) ([]interface{}, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) FoldingRange(ctx Context, conn *jsonrpc2.Conn, params *FoldingRangeParams) ([]FoldingRange, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Formatting(ctx Context, conn *jsonrpc2.Conn, params *DocumentFormattingParams) ([]TextEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Hover(ctx Context, conn *jsonrpc2.Conn, params *HoverParams) (*Hover, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Implementation(ctx Context, conn *jsonrpc2.Conn, params *ImplementationParams) ([]Location, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) InlayHint(ctx Context, conn *jsonrpc2.Conn, params *InlayHintParams) ([]InlayHint, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) InlineValue(ctx Context, conn *jsonrpc2.Conn, params *InlineValueParams) ([]InlineValue, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) LinkedEditingRange(ctx Context, conn *jsonrpc2.Conn, params *LinkedEditingRangeParams) (*LinkedEditingRanges, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Moniker(ctx Context, conn *jsonrpc2.Conn, params *MonikerParams) ([]Moniker, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) OnTypeFormatting(ctx Context, conn *jsonrpc2.Conn, params *DocumentOnTypeFormattingParams) ([]TextEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) PrepareCallHierarchy(ctx Context, conn *jsonrpc2.Conn, params *CallHierarchyPrepareParams) ([]CallHierarchyItem, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) PrepareRename(ctx Context, conn *jsonrpc2.Conn, params *PrepareRenameParams) (*PrepareRename2Gn, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) PrepareTypeHierarchy(ctx Context, conn *jsonrpc2.Conn, params *TypeHierarchyPrepareParams) ([]TypeHierarchyItem, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) RangeFormatting(ctx Context, conn *jsonrpc2.Conn, params *DocumentRangeFormattingParams) ([]TextEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) References(ctx Context, conn *jsonrpc2.Conn, params *ReferenceParams) ([]Location, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Rename(ctx Context, conn *jsonrpc2.Conn, params *RenameParams) (*WorkspaceEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) SelectionRange(ctx Context, conn *jsonrpc2.Conn, params *SelectionRangeParams) ([]SelectionRange, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) SemanticTokensFull(ctx Context, conn *jsonrpc2.Conn, params *SemanticTokensParams) (*SemanticTokens, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) SemanticTokensFullDelta(ctx Context, conn *jsonrpc2.Conn, params *SemanticTokensDeltaParams) (interface{}, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) SemanticTokensRange(ctx Context, conn *jsonrpc2.Conn, params *SemanticTokensRangeParams) (*SemanticTokens, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) SignatureHelp(ctx Context, conn *jsonrpc2.Conn, params *SignatureHelpParams) (*SignatureHelp, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) TypeDefinition(ctx Context, conn *jsonrpc2.Conn, params *TypeDefinitionParams) ([]Location, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) WillSave(ctx Context, conn *jsonrpc2.Conn, params *WillSaveTextDocumentParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) WillSaveWaitUntil(ctx Context, conn *jsonrpc2.Conn, params *WillSaveTextDocumentParams) ([]TextEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Subtypes(ctx Context, conn *jsonrpc2.Conn, params *TypeHierarchySubtypesParams) ([]TypeHierarchyItem, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Supertypes(ctx Context, conn *jsonrpc2.Conn, params *TypeHierarchySupertypesParams) ([]TypeHierarchyItem, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) WorkDoneProgressCancel(ctx Context, conn *jsonrpc2.Conn, params *WorkDoneProgressCancelParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DiagnosticWorkspace(ctx Context, conn *jsonrpc2.Conn, params *WorkspaceDiagnosticParams) (*WorkspaceDiagnosticReport, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) DidChangeConfiguration(ctx Context, conn *jsonrpc2.Conn, params *DidChangeConfigurationParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidChangeWatchedFiles(ctx Context, conn *jsonrpc2.Conn, params *DidChangeWatchedFilesParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidChangeWorkspaceFolders(ctx Context, conn *jsonrpc2.Conn, params *DidChangeWorkspaceFoldersParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidCreateFiles(ctx Context, conn *jsonrpc2.Conn, params *CreateFilesParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidDeleteFiles(ctx Context, conn *jsonrpc2.Conn, params *DeleteFilesParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) DidRenameFiles(ctx Context, conn *jsonrpc2.Conn, params *RenameFilesParams) error {
	return ErrMethodNotFound
}

func (that *AbsServer) ExecuteCommand(ctx Context, conn *jsonrpc2.Conn, params *ExecuteCommandParams) (interface{}, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) Symbol(ctx Context, conn *jsonrpc2.Conn, params *WorkspaceSymbolParams) ([]SymbolInformation, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) WillCreateFiles(ctx Context, conn *jsonrpc2.Conn, params *CreateFilesParams) (*WorkspaceEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) WillDeleteFiles(ctx Context, conn *jsonrpc2.Conn, params *DeleteFilesParams) (*WorkspaceEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) WillRenameFiles(ctx Context, conn *jsonrpc2.Conn, params *RenameFilesParams) (*WorkspaceEdit, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) ResolveWorkspaceSymbol(ctx Context, conn *jsonrpc2.Conn, params *WorkspaceSymbol) (*WorkspaceSymbol, error) {
	return nil, ErrMethodNotFound
}

func (that *AbsServer) NonstandardRequest(ctx Context, method string, params interface{}) (interface{}, error) {
	return nil, ErrMethodNotFound
}
