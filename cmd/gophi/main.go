package main

import (
	"encoding/json"
	"fmt"
	"gophi/cli"
	"os"
	"os/exec"
)

type Task struct {
	Name    string `json:"name"`
	Command string `json:"command"`
	Icon    string `json:"icon"`
}

func main() {
	cfg, err := cli.ParseArgs()
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred while parsing command arguments: %v\n", err)
		os.Exit(1)
	}

	configFile := cfg.SourceJSONFilePath

	jsonData, err := os.ReadFile(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file '%s': %v\n", configFile, err)
		os.Exit(1)
	}

	var tasks []Task
	if err := json.Unmarshal(jsonData, &tasks); err != nil {
		fmt.Fprintf(os.Stderr, "Error unmarshaling JSON from '%s': %v\n", configFile, err)
		os.Exit(1)
	}

	if cfg.ExecCmd == "" {
		for _, task := range tasks {
			if task.Name == "" {
				continue
			}
			icon := task.Icon
			if icon == "" {
				icon = "system-run"
			}
			fmt.Printf("%s\x00icon\x1f%s\n", task.Name, icon)
		}
		os.Exit(0)
	} else {
		selectedTaskName := cfg.ExecCmd

		var selectedTask *Task
		for i := range tasks {
			if tasks[i].Name == selectedTaskName {
				selectedTask = &tasks[i]
				break
			}
		}

		if selectedTask == nil {
			fmt.Fprintf(os.Stderr, "Error: Task '%s' not found in config file. No command to execute.\n", selectedTaskName)
			os.Exit(1)
		}

		commandToExecute := selectedTask.Command
		if commandToExecute == "" {
			commandToExecute = selectedTask.Name
		}

		cmd := exec.Command("bash", "-c", commandToExecute)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err = cmd.Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error starting command '%s': %v\n", commandToExecute, err)
			os.Exit(1)
		}
		os.Exit(0)
	}
}
