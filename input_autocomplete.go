package input_autocomplete

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"runtime"
)

func keyboardListener(input *Input) error {
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			return err
		}

		switch key {
		case keyboard.KeyEnter:
			fmt.Println("")
			return nil
		case keyboard.KeyArrowLeft:
			input.MoveCursorLeft()
		case keyboard.KeyArrowRight:
			input.MoveCursorRight()
		case keyboard.KeyBackspace:
			input.RemoveChar()
		case keyboard.KeyBackspace2:
			input.RemoveChar()
		case keyboard.KeyTab:
			input.Autocomplete()

		default:
			input.AddChar(char)
		}
	}
}

func Read(text string) (string, error) {
	if err := keyboard.Open(); err != nil {
		return "", err
	}

	defer keyboard.Close()

	os := runtime.GOOS
	if os == "windows" {
		// version := getWindowsVersion()
		// if !strings.Contains(version[1], "2012") {
			if err := EnableVirtualTerminalWindows(); err != nil {
				return "", err
			}
		// }
	}

	fmt.Println("before NewInput")
	input := NewInput(text)

	fmt.Println("before input.Print")
	input.Print()

	fmt.Println("before keyboardListener")
	if err := keyboardListener(input); err != nil {
		return "", err
	}

	fmt.Println("before input.GetCurrentText")
	return input.GetCurrentText(), nil
}
