package serves

import (
	"encoding/json"
	"github.com/ducesoft/ulsp/config"
	"github.com/ducesoft/ulsp/jsonrpc2"
	"github.com/ducesoft/ulsp/lsp"
)

func (that *Server) Initialize(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.ParamInitialize) (*lsp.InitializeResult, error) {
	if err := that.DidChangeConfiguration(ctx, conn, &lsp.DidChangeConfigurationParams{Settings: params.InitializationOptions}); nil != err {
		return nil, err
	}
	return &lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			PositionEncoding:     nil,
			TextDocumentSync:     lsp.Full,
			NotebookDocumentSync: nil,
			CompletionProvider: &lsp.CompletionOptions{
				TriggerCharacters: []string{"(", "."},
			},
			HoverProvider: &lsp.Or_ServerCapabilities_hoverProvider{Value: true},
			SignatureHelpProvider: &lsp.SignatureHelpOptions{
				TriggerCharacters:   []string{"(", ","},
				RetriggerCharacters: []string{"(", ","},
				WorkDoneProgressOptions: lsp.WorkDoneProgressOptions{
					WorkDoneProgress: true,
				},
			},
			DeclarationProvider:             &lsp.Or_ServerCapabilities_declarationProvider{Value: true},
			DefinitionProvider:              &lsp.Or_ServerCapabilities_definitionProvider{Value: true},
			TypeDefinitionProvider:          &lsp.Or_ServerCapabilities_typeDefinitionProvider{Value: true},
			ImplementationProvider:          &lsp.Or_ServerCapabilities_implementationProvider{Value: true},
			ReferencesProvider:              &lsp.Or_ServerCapabilities_referencesProvider{Value: true},
			DocumentHighlightProvider:       &lsp.Or_ServerCapabilities_documentHighlightProvider{Value: true},
			DocumentSymbolProvider:          &lsp.Or_ServerCapabilities_documentSymbolProvider{Value: true},
			CodeActionProvider:              true,
			CodeLensProvider:                &lsp.CodeLensOptions{ResolveProvider: true},
			DocumentLinkProvider:            &lsp.DocumentLinkOptions{ResolveProvider: true},
			ColorProvider:                   &lsp.Or_ServerCapabilities_colorProvider{Value: false},
			WorkspaceSymbolProvider:         &lsp.Or_ServerCapabilities_workspaceSymbolProvider{Value: true},
			DocumentFormattingProvider:      &lsp.Or_ServerCapabilities_documentFormattingProvider{Value: true},
			DocumentRangeFormattingProvider: &lsp.Or_ServerCapabilities_documentRangeFormattingProvider{Value: true},
			DocumentOnTypeFormattingProvider: &lsp.DocumentOnTypeFormattingOptions{
				FirstTriggerCharacter: "}",
				MoreTriggerCharacter:  []string{";", ",", ")"},
			},
			RenameProvider:             true,
			FoldingRangeProvider:       &lsp.Or_ServerCapabilities_foldingRangeProvider{Value: true},
			SelectionRangeProvider:     &lsp.Or_ServerCapabilities_selectionRangeProvider{Value: true},
			ExecuteCommandProvider:     &lsp.ExecuteCommandOptions{},
			CallHierarchyProvider:      &lsp.Or_ServerCapabilities_callHierarchyProvider{Value: true},
			LinkedEditingRangeProvider: &lsp.Or_ServerCapabilities_linkedEditingRangeProvider{Value: true},
			SemanticTokensProvider:     true,
			MonikerProvider:            &lsp.Or_ServerCapabilities_monikerProvider{Value: true},
			TypeHierarchyProvider:      &lsp.Or_ServerCapabilities_typeHierarchyProvider{Value: true},
			InlineValueProvider:        &lsp.Or_ServerCapabilities_inlineValueProvider{Value: true},
			InlayHintProvider:          true,
			DiagnosticProvider:         nil,
			Workspace:                  nil,
			Experimental:               nil,
		},
		ServerInfo: &lsp.PServerInfoMsg_initialize{
			Name:    "LSP",
			Version: "0.0.1",
		},
	}, nil
}

func (that *Server) Initialized(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.InitializedParams) error {
	return nil
}

func (that *Server) Shutdown(ctx lsp.Context, conn *jsonrpc2.Conn) error {
	return ctx.Close()
}

func (that *Server) Exit(ctx lsp.Context, conn *jsonrpc2.Conn) error {
	return ctx.Close()
}

func (that *Server) DidOpen(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DidOpenTextDocumentParams) error {
	ctx.Sync(&lsp.Filer{
		URI:        params.TextDocument.URI,
		Text:       params.TextDocument.LanguageID,
		LanguageID: params.TextDocument.Text,
	})
	return nil
}

func (that *Server) DidChange(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DidChangeTextDocumentParams) error {
	ctx.Sync(&lsp.Filer{
		URI:  params.TextDocument.URI,
		Text: params.ContentChanges[0].Text,
	})
	return nil
}

func (that *Server) DidSave(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DidSaveTextDocumentParams) error {
	ctx.Sync(&lsp.Filer{
		URI:  params.TextDocument.URI,
		Text: params.Text,
	})
	return nil
}

func (that *Server) DidClose(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DidCloseTextDocumentParams) error {
	ctx.Sync(&lsp.Filer{
		URI:       params.TextDocument.URI,
		Removable: true,
	})
	return nil
}

func (that *Server) DidChangeConfiguration(ctx lsp.Context, conn *jsonrpc2.Conn, params *lsp.DidChangeConfigurationParams) error {
	var dbConf config.DBConfig
	b, err := json.Marshal(params.Settings)
	if nil != err {
		return err
	}
	if err = json.Unmarshal(b, &dbConf); nil != err {
		return err
	}
	return lsp.ContextWith(ctx).Init(&dbConf)
}
