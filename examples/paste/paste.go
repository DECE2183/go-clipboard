// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Copyright (c) 2025 Timofey Korolik. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dece2183/go-clipboard"
)

var primaryFlag = flag.Bool("primary", false, "Use primary buffer on Linux")

func main() {
	flag.Parse()
	c := clipboard.New(clipboard.ClipboardOptions{Primary: *primaryFlag})

	text, err := c.PasteText()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("text from clipboard: %v\n", text)
}
