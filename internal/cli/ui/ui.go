package ui

import (
    _ "embed"
    "fmt"
    "strings"

    "github.com/Cod2rDude/bgs/internal/cli/color"
)

// Variables

//go:embed assets/banner.txt
var banner string

// Public Functions
func Log(append int, option string, message string) {
    switch option {
    case "warning":
        fmt.Println(strings.Repeat(" ", append) + color.Paint(color.Orange, "[WARNING] ") + color.Paint(color.Reset, message))
    case "error":
        fmt.Println(strings.Repeat(" ", append) + color.Paint(color.Red, "[ERROR] ") + color.Paint(color.Reset, message))
    default:
        fmt.Println(strings.Repeat(" ", append) + color.Paint(color.Blue, "[INFO] ") + color.Paint(color.Reset, message))
    }
}

func Startup() {

}
