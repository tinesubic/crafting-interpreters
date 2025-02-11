package main

import (
	"crafting-interpreters/internal"
	"crafting-interpreters/internal/loxerr"
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
	fmt.Println("Running GLOX with input file: " + filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("File %s does not exist", filePath))
	}

	glox := internal.Glox{}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	if err := glox.Run(string(data)); err != nil {
		return err
	}

	return nil
}

func runPrompt() error {
	print("Running GLOX in REPL mode")

	glox := internal.Glox{}
	for {
		fmt.Print("> ")
		glox.ResetError()
		reader, err := io.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		input := strings.TrimSpace(string(reader))
		if input == "" {
			break
		}
		if err := glox.Run(input); err != nil {
			if errors.Is(err, loxerr.GloxError{}) {
				continue
			}
			return err
		}
	}
	print("\nGLOX exiting")
	return nil
}
