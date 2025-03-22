//go:build darwin

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"os/exec"
)

const (
	// pbcopy is the name of the macOS command-line tool used to copy text to the clipboard.
	pbcopy = "pbcopy"
	// pbpaste is the name of the macOS command-line tool used to paste text from the clipboard.
	pbpaste = "pbpaste"
)

var (
	// copyTool is a preconfigured CopyTool for macOS using the pbcopy utility.
	copyTool = Tool{
		name: pbcopy,
	}
	// pasteTool is a preconfigured PasteTool for macOS using the pbpaste utility.
	pasteTool = Tool{
		name: pbpaste,
	}
	// lookPath is a variable holding the exec.LookPath function,
	// used to check for the presence of a command in the system's PATH.
	lookPath = exec.LookPath
)

// newClipboardTool initializes a new clipboardTool instance by
// checking the availability of clipboard utilities.
func newClipboardTool(primary bool) *ClipboardTool {
	cbtool := &ClipboardTool{}
	if isAvailable := isToolAvailable(copyTool.name); isAvailable {
		cbtool.CopyTool = &copyTool
	}
	if isAvailable := isToolAvailable(pasteTool.name); isAvailable {
		cbtool.PasteTool = &pasteTool
	}
	return cbtool
}

// isToolAvailable checks if a clipboard utility tool
// is available in the system's PATH.
func isToolAvailable(toolName string) bool {
	if _, err := lookPath(toolName); err != nil {
		return false
	}
	return true
}
