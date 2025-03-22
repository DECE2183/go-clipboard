# go-clipboard

![logo](golang-clipboard.png)

A cross-platform clipboard utility for [Go](go.dev).

It offers abilities for copying and pasting plain text.

## installation

```
go get github.com/dece2183/go-clipboard
```

## documentation

[https://pkg.go.dev/github.com/dece2183/go-clipboard](https://pkg.go.dev/github.com/dece2183/go-clipboard)

## platforms

| OS | Supported copy tools | Supported paste tools |
|----------|----------|----------|
| Darwin | `pbcopy` | `pbpaste` |
| Windows | `clip.exe` | `powershell` |
| Linux/FreeBSD/NetBSD/OpenBSD/Dragonfly| X11: `xsel`, `xclip` <br> Wayland: `wl-copy` | X11: `xsel`, `xclip` <br> Wayland: `wl-paste` |
| Solaris | X11: `xsel`, `xclip`| X11: `xsel`, `xclip` |
| Android (via Termux) | `termux-clipboard-set`| `termux-clipboard-get` |

## examples

### copy ([examples/copy/copy.go](examples/copy/copy.go))

```go
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
```

### paste ([examples/paste/paste.go](examples/paste/paste.go))

```go
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
```
