//go:build freebsd || linux || netbsd || openbsd || solaris || dragonfly

// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package clipboardtool

import "os/exec"

const (
	// xsel is a clipboard utility for X11.
	xsel = "xsel"
	// xclip is another clipboard utility for X11.
	xclip = "xclip"

	// wlcopy is a clipboard utility for the Wayland display server.
	wlcopy = "wl-copy"
	// wlpaste is a clipboard utility for the Wayland display server.
	wlpaste = "wl-paste"

	// termuxClipboardGet is a clipboard utility for Termux, an Android terminal emulator.
	termuxClipboardGet = "termux-clipboard-get"
	// termuxClipboardSet is a clipboard utility for Termux, an Android terminal emulator.
	termuxClipboardSet = "termux-clipboard-set"
)

var (
	// copyTools is a list of available CopyTool configurations for different environments.
	copyTools = []Tool{
		{
			name: wlcopy,
		},
		{
			name: xsel,
			args: []string{"--input", "--clipboard"},
		},
		{
			name: xclip,
			args: []string{"-in", "-selection", "clipboard"},
		},
		{
			name: termuxClipboardSet,
		},
	}
	// pasteTools is a list of available PasteTool configurations for different environments.
	pasteTools = []Tool{
		{
			name: wlpaste,
			args: []string{"--no-newline"},
		},
		{
			name: xsel,
			args: []string{"--output", "--clipboard"},
		},
		{
			name: xclip,
			args: []string{"-out", "-selection", "clipboard"},
		},
		{
			name: termuxClipboardGet,
		},
	}
	// same with primary selection
	copyToolsPrimary = []Tool{
		{
			name: wlcopy,
			args: []string{"--primary"},
		},
		{
			name: xsel,
			args: []string{"--input", "--primary"},
		},
		{
			name: xclip,
			args: []string{"-in", "-selection", "primary"},
		},
		{
			name: termuxClipboardSet,
		},
	}
	pasteToolsPrimary = []Tool{
		{
			name: wlpaste,
			args: []string{"--no-newline", "--primary"},
		},
		{
			name: xsel,
			args: []string{"--output", "--primary"},
		},
		{
			name: xclip,
			args: []string{"-out", "-selection", "primary"},
		},
		{
			name: termuxClipboardGet,
		},
	}
)

// newClipboardTool selects the first available pair of
// copy and paste tools from the predefined list.
func newClipboardTool(primary bool) *ClipboardTool {
	if primary {
		return &ClipboardTool{
			CopyTool:  findTools(copyToolsPrimary),
			PasteTool: findTools(pasteToolsPrimary),
		}
	} else {
		return &ClipboardTool{
			CopyTool:  findTools(copyTools),
			PasteTool: findTools(pasteTools),
		}
	}
}

// Search for a tool from array in a system PATH
// and return the first occurrence
func findTools(tools []Tool) *Tool {
	var err error
	for i := range tools {
		_, err = exec.LookPath(tools[i].name)
		if err == nil {
			return &tools[i]
		}
	}
	return nil
}
