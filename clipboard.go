// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboard

import (
	"errors"

	"github.com/dece2183/go-clipboard/clipboardtool"
)

var (
	ErrNoCopyToolFound  = errors.New("no clipboard copy tool available")
	ErrNoPasteToolFound = errors.New("no clipboard paste tool available")
)

// Clipboard wraps the basic clipboard operations.
type Clipboard struct {
	tool *clipboardtool.ClipboardTool
}

// exported flag container
type ClipboardOptions struct {
	Primary bool
}

// New creates and returns a new Clipboard instance that can be used
// to interact with the system clipboard.
func New(opts ...ClipboardOptions) *Clipboard {
	var primary bool
	if len(opts) == 1 {
		primary = opts[0].Primary
	}
	return &Clipboard{
		tool: clipboardtool.New(primary),
	}
}

// CopyText implements the Clipboard interface's CopyText method.
// It calls the copyText function to perform the actual operation.
func (cb *Clipboard) CopyText(s string) error {
	if cb.tool.CopyTool == nil {
		return ErrNoCopyToolFound
	}
	cmd := cb.tool.CopyTool.Command()
	return cmd.TextInput(s)
}

// PasteText implements the Clipboard interface's PasteText method.
// It calls the pasteText function to perform the actual operation.
func (cb *Clipboard) PasteText() (string, error) {
	if cb.tool.PasteTool == nil {
		return "", ErrNoPasteToolFound
	}
	cmd := cb.tool.PasteTool.Command()
	return cmd.TextOutput()
}
