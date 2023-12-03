package ab5logging

import (
	"fmt"
	"time"
)

var (
	Red    = "\033[38;2;255;0;0m"
	Green  = "\033[38;2;0;255;0m"
	Blue   = "\033[38;2;0;0;255m"
	Yellow = "\033[38;2;255;255;0m"
	Reset  = "\033[39m"
)

func Log(v any) { // blue [*]
	fmt.Println(Reset, "[", Blue, "*", Reset, "]", v)
}

func Success(v any) { // green [/]
	fmt.Println(Reset, "[", Green, "/", Reset, "]", v)
}

func Warn(v any) { // yello [!]
	fmt.Println(Reset, "[", Yellow, "!", Reset, "]", v)
}

func Error(err error) { // red [X]
	fmt.Println(Reset, "[", Red, "X", Reset, "]", err)
}
