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

type Client interface {
	LogTrace(ctx Context, conn *jsonrpc2.Conn, params *LogTraceParams) error                                           // $/logTrace
	Progress(ctx Context, conn *jsonrpc2.Conn, params *ProgressParams) error                                           // $/progress
	RegisterCapability(ctx Context, conn *jsonrpc2.Conn, params *RegistrationParams) error                             // client/registerCapability
	UnregisterCapability(ctx Context, conn *jsonrpc2.Conn, params *UnregistrationParams) error                         // client/unregisterCapability
	Event(ctx Context, conn *jsonrpc2.Conn, params *interface{}) error                                                 // telemetry/event
	PublishDiagnostics(ctx Context, conn *jsonrpc2.Conn, params *PublishDiagnosticsParams) error                       // textDocument/publishDiagnostics
	LogMessage(ctx Context, conn *jsonrpc2.Conn, params *LogMessageParams) error                                       // window/logMessage
	ShowDocument(ctx Context, conn *jsonrpc2.Conn, params *ShowDocumentParams) (*ShowDocumentResult, error)            // window/showDocument
	ShowMessage(ctx Context, conn *jsonrpc2.Conn, params *ShowMessageParams) error                                     // window/showMessage
	ShowMessageRequest(ctx Context, conn *jsonrpc2.Conn, params *ShowMessageRequestParams) (*MessageActionItem, error) // window/showMessageRequest
	WorkDoneProgressCreate(ctx Context, conn *jsonrpc2.Conn, params *WorkDoneProgressCreateParams) error               // window/workDoneProgress/create
	ApplyEdit(ctx Context, conn *jsonrpc2.Conn, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error)   // workspace/applyEdit
	CodeLensRefresh(ctx Context, conn *jsonrpc2.Conn) error                                                            // workspace/codeLens/refresh
	Configuration(ctx Context, conn *jsonrpc2.Conn, params *ParamConfiguration) ([]LSPAny, error)                      // workspace/configuration
	DiagnosticRefresh(ctx Context, conn *jsonrpc2.Conn) error                                                          // workspace/diagnostic/refresh
	InlayHintRefresh(ctx Context, conn *jsonrpc2.Conn) error                                                           // workspace/inlayHint/refresh
	InlineValueRefresh(ctx Context, conn *jsonrpc2.Conn) error                                                         // workspace/inlineValue/refresh
	SemanticTokensRefresh(ctx Context, conn *jsonrpc2.Conn) error                                                      // workspace/semanticTokens/refresh
	WorkspaceFolders(ctx Context, conn *jsonrpc2.Conn) ([]WorkspaceFolder, error)                                      // workspace/workspaceFolders
}

func clientDispatch(ctx Context, client Client, conn *jsonrpc2.Conn, r Request) (any, error) {
	switch r.Method() {
	case "$/logTrace":
		var params LogTraceParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.LogTrace(ctx, conn, &params)
		return nil, err
	case "$/progress":
		var params ProgressParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.Progress(ctx, conn, &params)
		return nil, err
	case "client/registerCapability":
		var params RegistrationParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.RegisterCapability(ctx, conn, &params)
		return nil, err
	case "client/unregisterCapability":
		var params UnregistrationParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.UnregisterCapability(ctx, conn, &params)
		return nil, err
	case "telemetry/event":
		var params interface{}
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.Event(ctx, conn, &params)
		return nil, err
	case "textDocument/publishDiagnostics":
		var params PublishDiagnosticsParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.PublishDiagnostics(ctx, conn, &params)
		return nil, err
	case "window/logMessage":
		var params LogMessageParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.LogMessage(ctx, conn, &params)
		return nil, err
	case "window/showDocument":
		var params ShowDocumentParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := client.ShowDocument(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "window/showMessage":
		var params ShowMessageParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.ShowMessage(ctx, conn, &params)
		return nil, err
	case "window/showMessageRequest":
		var params ShowMessageRequestParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := client.ShowMessageRequest(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "window/workDoneProgress/create":
		var params WorkDoneProgressCreateParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		err := client.WorkDoneProgressCreate(ctx, conn, &params)
		return nil, err
	case "workspace/applyEdit":
		var params ApplyWorkspaceEditParams
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := client.ApplyEdit(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/codeLens/refresh":
		err := client.CodeLensRefresh(ctx, conn)
		return nil, err
	case "workspace/configuration":
		var params ParamConfiguration
		if err := json.Unmarshal(r.Params(), &params); err != nil {
			return nil, ErrCodeParseFn(err)
		}
		resp, err := client.Configuration(ctx, conn, &params)
		if err != nil {
			return nil, err
		}
		return resp, nil
	case "workspace/diagnostic/refresh":
		err := client.DiagnosticRefresh(ctx, conn)
		return nil, err
	case "workspace/inlayHint/refresh":
		err := client.InlayHintRefresh(ctx, conn)
		return nil, err
	case "workspace/inlineValue/refresh":
		err := client.InlineValueRefresh(ctx, conn)
		return nil, err
	case "workspace/semanticTokens/refresh":
		err := client.SemanticTokensRefresh(ctx, conn)
		return nil, err
	case "workspace/workspaceFolders":
		resp, err := client.WorkspaceFolders(ctx, conn)
		if err != nil {
			return nil, err
		}
		return resp, nil
	default:
		return nil, ErrMethodNotFound
	}
}

func (s *clientDispatcher) LogTrace(ctx Context, params *LogTraceParams) error {
	return s.sender.Notify(ctx, "$/logTrace", params)
}
func (s *clientDispatcher) Progress(ctx Context, params *ProgressParams) error {
	return s.sender.Notify(ctx, "$/progress", params)
}
func (s *clientDispatcher) RegisterCapability(ctx Context, params *RegistrationParams) error {
	return s.sender.Call(ctx, "client/registerCapability", params, nil)
}
func (s *clientDispatcher) UnregisterCapability(ctx Context, params *UnregistrationParams) error {
	return s.sender.Call(ctx, "client/unregisterCapability", params, nil)
}
func (s *clientDispatcher) Event(ctx Context, params *interface{}) error {
	return s.sender.Notify(ctx, "telemetry/event", params)
}
func (s *clientDispatcher) PublishDiagnostics(ctx Context, params *PublishDiagnosticsParams) error {
	return s.sender.Notify(ctx, "textDocument/publishDiagnostics", params)
}
func (s *clientDispatcher) LogMessage(ctx Context, params *LogMessageParams) error {
	return s.sender.Notify(ctx, "window/logMessage", params)
}
func (s *clientDispatcher) ShowDocument(ctx Context, params *ShowDocumentParams) (*ShowDocumentResult, error) {
	var result *ShowDocumentResult
	if err := s.sender.Call(ctx, "window/showDocument", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) ShowMessage(ctx Context, params *ShowMessageParams) error {
	return s.sender.Notify(ctx, "window/showMessage", params)
}
func (s *clientDispatcher) ShowMessageRequest(ctx Context, params *ShowMessageRequestParams) (*MessageActionItem, error) {
	var result *MessageActionItem
	if err := s.sender.Call(ctx, "window/showMessageRequest", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) WorkDoneProgressCreate(ctx Context, params *WorkDoneProgressCreateParams) error {
	return s.sender.Call(ctx, "window/workDoneProgress/create", params, nil)
}
func (s *clientDispatcher) ApplyEdit(ctx Context, params *ApplyWorkspaceEditParams) (*ApplyWorkspaceEditResult, error) {
	var result *ApplyWorkspaceEditResult
	if err := s.sender.Call(ctx, "workspace/applyEdit", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) CodeLensRefresh(ctx Context) error {
	return s.sender.Call(ctx, "workspace/codeLens/refresh", nil, nil)
}
func (s *clientDispatcher) Configuration(ctx Context, params *ParamConfiguration) ([]LSPAny, error) {
	var result []LSPAny
	if err := s.sender.Call(ctx, "workspace/configuration", params, &result); err != nil {
		return nil, err
	}
	return result, nil
}
func (s *clientDispatcher) DiagnosticRefresh(ctx Context) error {
	return s.sender.Call(ctx, "workspace/diagnostic/refresh", nil, nil)
}
func (s *clientDispatcher) InlayHintRefresh(ctx Context) error {
	return s.sender.Call(ctx, "workspace/inlayHint/refresh", nil, nil)
}
func (s *clientDispatcher) InlineValueRefresh(ctx Context) error {
	return s.sender.Call(ctx, "workspace/inlineValue/refresh", nil, nil)
}
func (s *clientDispatcher) SemanticTokensRefresh(ctx Context) error {
	return s.sender.Call(ctx, "workspace/semanticTokens/refresh", nil, nil)
}
func (s *clientDispatcher) WorkspaceFolders(ctx Context) ([]WorkspaceFolder, error) {
	var result []WorkspaceFolder
	if err := s.sender.Call(ctx, "workspace/workspaceFolders", nil, &result); err != nil {
		return nil, err
	}
	return result, nil
}
