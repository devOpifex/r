package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}

type Commands []Command

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Printf("No command specified\n")
		os.Exit(1)
	}

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory\n")
		os.Exit(1)
	}

	path := home + "/.rconfig"

	_, err = os.Stat(path)
	if err != nil {
		fmt.Printf("File %s does not exist\n", path)
		os.Exit(1)
	}

	buffer, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file %s\n", path)
		os.Exit(1)
	}

	cmds := Commands{}
	err = json.Unmarshal(buffer, &cmds)
	if err != nil {
		fmt.Printf("Error unmarshalling file %s\n", path)
		os.Exit(1)
	}

	for _, arg := range args {
		for _, cmd := range cmds {
			if cmd.Name == arg {
				fmt.Printf("%s\n", cmd.Command)
				c := exec.Command("Rscript", "-e", cmd.Command)
				c.Stdout = os.Stdout
				c.Stderr = os.Stderr

				err := c.Run()
				if err != nil {
					fmt.Printf("Error running command %s\n", cmd.Command)
					os.Exit(1)
				}
			}
		}
	}
}
