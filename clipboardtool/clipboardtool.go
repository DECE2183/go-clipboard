// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"github.com/dece2183/go-clipboard/clipboardtool/command"
)

// Tool encapsulates the details of a shell command.
type Tool struct {
	name string   // Name of the command or executable
	args []string // Arguments required for the operation
}

func (t *Tool) Command() *command.Command {
	return command.New(t.name, t.args...)
}

// clipboardTool combines Copy and Paste tools to provide a unified interface
// for clipboard operations. It abstracts the underlying command-line tools used
// to interact with the system clipboard.
type ClipboardTool struct {
	CopyTool  *Tool // Tool to copy content to the clipboard
	PasteTool *Tool // Tool to paste content from the clipboard
}

// New initializes and returns a new instance of clipboardTool.
// It determines the appropriate tools to use based on the current system environment
// and returns an error if no suitable tools are found.
func New(primary bool) *ClipboardTool {
	return newClipboardTool(primary)
}
