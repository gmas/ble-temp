package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

//TODO run every 2 mins
func readSensors(s chan []byte, ctrl chan int) {
	// Create an *exec.Cmd
	fmt.Printf("==>readSensors\n")
	cmd := exec.Command("/usr/bin/node", "test.js")

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	printCommand(cmd)
	err := cmd.Run() // will wait for command to return
	if err != nil {
		printError(err)
		ctrl <- 1
	}

	// Only output the commands stdout
	//printOutput(cmdOutput.Bytes())
	s <- cmdOutput.Bytes()
}

func printValues(s chan []byte, ctrl chan int) {
	for {
		result := <-s
		printOutput(result)
		ctrl <- 0
	}
}

func main() {
	fmt.Printf("==>main\n")
	c := make(chan []byte)
	ctrl := make(chan int)
	go readSensors(c, ctrl)
	go printValues(c, ctrl)

	for {
		result := <-ctrl
		fmt.Printf("==> Exit: %v\n", result)
	}
}
