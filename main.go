package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("shyell>")
		cmd_string, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error")
			continue
		}

		cmd_string = strings.TrimSpace(cmd_string)

		if cmd_string == "exit" {
			fmt.Println("Exiting...")
			break
		}

		parts := strings.Split(cmd_string, " ")

		switch parts[0] { // 'cd' to path but not more than one.
		case "cd":
			if len(parts) < 2 {
				fmt.Errorf("Error")
			}

			os.Chdir(parts[1])
		}

		cmd := parts[0]
		args := parts[1:]

		run_cmd := exec.Command(cmd, args...)
		run_cmd.Stderr = os.Stderr
		run_cmd.Stdout = os.Stdout
		run_cmd.Run()

	}

}
