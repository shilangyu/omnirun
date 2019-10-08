package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

// Runner describes how to run a file
type Runner struct {
	// Exts is a slice of extensions
	Exts []string `json:"exts"`
	// Run constructs a slice of commands.
	// $or_file is the source file
	Run []string `json:"run"`
}

func loadRunners() ([]Runner, error) {
	runners := []Runner{}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	runnersPath := path.Join(configDir, "omnirun", "runners.json")
	if _, err := os.Stat(runnersPath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(runnersPath), 0644)
		ioutil.WriteFile(runnersPath, []byte(runnersJSON), 0644)
	}

	data, err := ioutil.ReadFile(runnersPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &runners)
	if err != nil {
		return nil, err
	}

	return runners, nil
}

func main() {
	var path string

	runners, err := loadRunners()

	if err != nil {
		log.Fatalf("Failed to load runners.json: %s\n", err)
	}

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
					for _, command := range runner.Run {
						command = strings.ReplaceAll(command, "$or_file", path)
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
