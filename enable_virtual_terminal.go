// +build windows

package input_autocomplete

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func EnableVirtalTerminalWindows() error {
	var originalMode uint32
	fmt.Println("before windows.Handle")
	stdout := windows.Handle(os.Stdout.Fd())

	fmt.Println("before windows.GetConsoleMode")
	if err := windows.GetConsoleMode(stdout, &originalMode); err != nil {
		fmt.Println("windows.GetConsoleMode error")
		return err
	}

	fmt.Println("before windows.SetConsoleMode")
	return windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}
