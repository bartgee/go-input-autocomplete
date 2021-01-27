// +build windows

package input_autocomplete

import (
	"syscall"
)

var (
	kernel32Dll    *syscall.LazyDLL  = syscall.NewLazyDLL("Kernel32.dll")
	setConsoleMode *syscall.LazyProc = kernel32Dll.NewProc("SetConsoleMode")
)

// func EnableVirtalTerminalWindows() error {
// 	var originalMode uint32
// 	// var vtInputSupported bool
// 	fmt.Println("before windows.Handle")
// 	stdout := windows.Handle(os.Stdout.Fd())
//
// 	fmt.Println("before windows.GetConsoleMode")
// 	if err := windows.GetConsoleMode(stdout, &originalMode); err != nil {
// 		fmt.Println("windows.GetConsoleMode error")
// 		return err
// 	}
//
// 	fmt.Println("before windows.SetConsoleMode")
// 	if err := windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING); err != nil {
// 		// vtInputSupported := true
// 		return nil
// 	}
// 	return windows.SetConsoleMode(stdout, originalMode)
// }

func EnableVirtualTerminalWindows() error {
	const ENABLE_VIRTUAL_TERMINAL_PROCESSING uint32 = 0x4

	var mode uint32
	err := syscall.GetConsoleMode(syscall.Stdout, &mode)
	if err != nil {
		return err
	}

	// if enable {
	mode |= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	// } else {
	// 	mode &^= ENABLE_VIRTUAL_TERMINAL_PROCESSING
	// }

	ret, _, err := setConsoleMode.Call(uintptr(syscall.Stdout), uintptr(mode))
	if ret == 0 {
		return err
	}

	return nil
}