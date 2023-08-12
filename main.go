package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func process(command string) (string, error) {
	parts := strings.Fields(command)
	cmd := exec.Command(parts[0], parts[1:]...)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func main() {
	for true {
		fmt.Print("\033[33m~> \033[0m")
		var command string; fmt.Scanln(&command)
		fmt.Println(process(command))
	}
}
