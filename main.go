package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Runner describes how to run a file
type Runner struct {
	// Exts is a slice of extensions
	Exts []string
	// Run constructs a slice of commands from a given file
	Run func(file string) []string
}

var runners = []Runner{
	{
		[]string{"js"},
		func(file string) []string {
			return []string{
				fmt.Sprintf("node %s", file),
			}
		},
	},
	{
		[]string{"py"},
		func(file string) []string {
			return []string{
				fmt.Sprintf("python %s", file),
			}
		},
	},
	{
		[]string{"go"},
		func(file string) []string {
			return []string{
				fmt.Sprintf("go run %s", file),
			}
		},
	},
	{
		[]string{"bf"},
		func(file string) []string {
			return []string{
				fmt.Sprintf("brainfuck %s", file),
			}
		},
	},
	{
		[]string{"cpp"},
		func(file string) []string {
			return []string{
				fmt.Sprintf("g++ %s -o abc123.exe", file),
				"./abc123.exe",
			}
		},
	},
	{
		[]string{"c"},
		func(file string) []string {
			return []string{
				fmt.Sprintf("gcc %s -o abc123.exe", file),
				"./abc123.exe",
			}
		},
	},
}

func main() {
	var path string

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdin := bufio.NewReader(os.Stdin)
		text, _ := stdin.ReadString('\n')
		path = strings.TrimSpace(text)
	} else {
		if len(os.Args) > 1 {
			path = os.Args[1]
		}
	}

	if stat, err := os.Stat(path); os.IsNotExist(err) || stat.IsDir() {
		fmt.Println("Provided path is not pointing to a file.")
	} else {
		ext := filepath.Ext(path)[1:]

		for _, runner := range runners {
			for _, ex := range runner.Exts {
				if ex == ext {
					for _, command := range runner.Run(path) {
						parts := strings.Fields(command)
						cmd := exec.Command(parts[0], parts[1:]...)
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						err := cmd.Run()
						if err != nil {
							log.Fatalf("Failed to run the commands: %s", err)
						}
					}
				}
			}
		}
	}
}
