package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("Hello Bitrise CLI Go Toolkit!!")

	// print Go version
	fmt.Println()
	fmt.Print("Go version: ")
	{
		cmd := exec.Command("go", "version")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}

	// input test
	fmt.Println()
	fmt.Println("Input test_input: ", os.Getenv("test_input"))

	// output test
	fmt.Println()
	fmt.Println("Generating output:")
	{
		cmd := exec.Command("envman", "add",
			"--key", "TEST_OUTPUT",
			"--value", "Test output value")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	fmt.Println()
}
