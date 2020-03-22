package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// Runner describes how to run a file
type Runner struct {
	// Exts is a slice of extensions
	Exts []string `yaml:"exts,flow"`
	// Run constructs a slice of commands.
	// $or_file is the source file
	Run []string `yaml:"run"`
}

func runnersConfigPath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	runnersPath := path.Join(configDir, "omnirun", "runners.yaml")

	return runnersPath, nil
}

func loadRunners() ([]Runner, error) {
	runners := []Runner{}

	runnersPath, err := runnersConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(runnersPath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(runnersPath), os.ModePerm)
		d, _ := yaml.Marshal(&initialRunners)
		ioutil.WriteFile(runnersPath, d, os.ModePerm)
	}

	data, err := ioutil.ReadFile(runnersPath)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &runners)
	if err != nil {
		return nil, err
	}

	return runners, nil
}

func main() {
	runners, err := loadRunners()
	check(err, "Failed to load runners")

	var fileToRun string
	action := parseArgs()

	switch action {
	case cmdActionConfig:
		p, _ := runnersConfigPath()
		fmt.Println(p)
		os.Exit(0)
	case cmdActionFromStdin:
		stdin := bufio.NewReader(os.Stdin)
		text, _ := stdin.ReadString('\n')
		fileToRun = strings.TrimSpace(text)
	case cmdActionFileInput:
		fileToRun = os.Args[1]
	}

	if stat, err := os.Stat(fileToRun); os.IsNotExist(err) || stat.IsDir() {
		errExitWith("Provided path is not pointing to a file.")
	} else {
		ext := filepath.Ext(fileToRun)
		if ext != "" {
			ext = ext[1:]
		}
		foundAny := false

		for _, runner := range runners {
			for _, ex := range runner.Exts {
				if ex == ext {
					foundAny = true
					for _, command := range runner.Run {
						command = strings.ReplaceAll(command, "$or_file", fileToRun)
						parts := strings.Fields(command)
						cmd := exec.Command(parts[0], parts[1:]...)
						cmd.Stdin = os.Stdin
						cmd.Stdout = os.Stdout
						cmd.Stderr = os.Stderr
						err := cmd.Run()
						check(err, "Omnirun failed to run the following command: "+command)
					}
				}
			}
		}

		if !foundAny {
			errExitWith("Extension '" + ext + "' does not have a defined runner.")
		}
	}
}
