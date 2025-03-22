// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package command

import (
	"fmt"
	"os/exec"
)

// Command is a wrapper around the os/exec.Cmd
// to execute system commands and process their input and output.
type Command struct {
	cmd *exec.Cmd
}

// New creates a new command instance with the specified exec.Cmd.
func New(name string, args ...string) *Command {
	return &Command{
		cmd: exec.Command(name, args...),
	}
}

// TextInput sends the provided text as input to the system command.
func (c *Command) TextInput(text string) error {
	in, err := c.cmd.StdinPipe()
	if err != nil {
		return fmt.Errorf("getting pipe for command: %w", err)
	}
	if err := c.cmd.Start(); err != nil {
		return fmt.Errorf("starting command: %w", err)
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return fmt.Errorf("writing input for command: %w", err)
	}
	if err := in.Close(); err != nil {
		return fmt.Errorf("closing input: %w", err)
	}
	if err := c.cmd.Wait(); err != nil {
		return fmt.Errorf("waiting for command: %w", err)
	}
	return nil
}

// TextOutput executes the command and returns its output as a string.
func (c *Command) TextOutput() (string, error) {
	out, err := c.cmd.Output()
	if err != nil {
		return "", fmt.Errorf("getting output for command: %w", err)
	}
	return string(out), nil
}
