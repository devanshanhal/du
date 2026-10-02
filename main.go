package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: du <command>")
		return
	}

	command := os.Args[1]

	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("job failed:")
		fmt.Println(string(output))
		fmt.Println(err)
		return
	}

	fmt.Println(string(output))
}
