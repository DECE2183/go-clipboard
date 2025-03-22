//go:build windows

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import (
	"os/exec"
)

const (
	// clip is the name of the Windows command-line tool used to copy text to the clipboard.
	clip = "clip.exe"
	// powershell is used to access advanced Windows functionality,
	// including the ability to paste text from the clipboard.
	powershell = "powershell"
)

var (
	// copyTool is a preconfigured CopyTool for Windows using the clip utility.
	copyTool = Tool{
		name: clip,
	}
	// pasteTool is a preconfigured PasteTool for Windows using PowerShell commands.
	pasteTool = Tool{
		name: powershell,
		args: []string{"Get-Clipboard"},
	}
	// lookPath is a variable holding the exec.LookPath function,
	// used to check for the presence of a command in the system's PATH.
	lookPath = exec.LookPath
)

// newClipboardTool checks the availability of clipboard utilities
// and initializes a new clipboardTool.
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

// isToolAvailable verifies the presence of a clipboard utility in the system's PATH.
func isToolAvailable(toolName string) bool {
	if _, err := lookPath(toolName); err != nil {
		return false
	}
	return true
}
