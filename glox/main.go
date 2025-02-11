package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	args := os.Args
	if len(args) == 2 {
		if err := runFile(args[1]); err != nil {
			panic(err)
		}
	} else if len(args) > 2 {
		print("Usage: glox [input file]")
		os.Exit(1)
	} else {
		if err := runPrompt(); err != nil {
			panic(err)
		}
	}
}

func runFile(filePath string) error {
	print("Running GLOX with input file: " + filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("File %s does not exist", filePath))
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := run(string(data)); err != nil {
		return err
	}

	return nil
}

func run(input string) error {
	println(input)
	return nil
}

func runPrompt() error {
	print("Running GLOX in REPL mode")

	for {
		print("> ")
		reader, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		input := strings.TrimSpace(string(reader))
		if input == "exit" {
			break
		}
		if err := run(input); err != nil {
			return err
		}
	}
	print("GLOX exiting")
	return nil
}
