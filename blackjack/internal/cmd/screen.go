package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Screen interface {
	Printf(msg string, args ...interface{})
	Input() string
	Reset()
}

type screen struct {
}

func NewScreen() Screen {
	return &screen{}
}

func (x screen) Printf(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (x screen) Input() string {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		x.Printf("can't read your input, please try again:")
		_, err = fmt.Scanln(&input)
		if err != nil {
			panic(err)
		}
	}
	return strings.TrimSpace(input)
}

func (x screen) Reset() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear") //Linux example, its tested
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		// do nothing here
	}
}
