package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	var _, filename, _, _ = runtime.Caller(0)
	var execFile, _ = filepath.Abs(filename)
	var execPath = filepath.Dir(execFile)
	var scanner = bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(":")
		scanner.Scan()
		var execLine = scanner.Text()

		if strings.HasPrefix(execLine, "exit") {
			fmt.Println("GoTerminal is finished")
			return
		}

		var args = strings.Fields(execLine)
		var command = execPath + "/" + args[0] + ".go"
		args = args[1:]

		if _, err := os.Stat(command); os.IsNotExist(err) {
			fmt.Println("GoTerminal: " + command + ": command is not exist")
			continue
		}

		Execute(command, args)
	}

}

func Execute(command string, args []string) {
	var execArgs = []string{"run", command}
	execArgs = append(execArgs, args[:]...)
	var cmd = exec.Command("go", execArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("GoTerminal: ", command, ": failed execute:\n\r\t ", err)
	}

	fmt.Println(cmd.Path)
}
