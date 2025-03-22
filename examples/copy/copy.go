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

	var text string
	args := flag.Args()
	if len(args) > 0 {
		text = args[len(args)-1]
	} else {
		text = "some text"
	}

	if err := c.CopyText(text); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("text \"%s\" was copied into clipboard. Paste it elsewhere.\n", text)
}
