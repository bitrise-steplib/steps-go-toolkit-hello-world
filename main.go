package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/colorstring"
)

func main() {
	fmt.Println(colorstring.Green("Hello Bitrise CLI Go Toolkit!!"))

	// print Go version
	fmt.Println()
	fmt.Print(colorstring.Blue("Go version: "))
	{
		cmd := exec.Command("go", "version")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	// input test
	fmt.Println()
	fmt.Println(colorstring.Blue("Input test_input: "), os.Getenv("test_input"))

	// output test
	fmt.Println()
	fmt.Println(colorstring.Blue("Generating output ..."))
	{
		cmd := exec.Command("envman", "add",
			"--key", "TEST_OUTPUT",
			"--value", "Test output value")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	fmt.Println()
	fmt.Println(colorstring.Green("DONE"))
}
