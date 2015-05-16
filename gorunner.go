package main

import (
	"./dal"
	"bytes"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"os/exec"
	"strings"
	"time"
)

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func readSensors(s chan []byte, ctrl chan int) {
	ticker := time.NewTicker(time.Millisecond * 60000)

	for t := range ticker.C {
		// Create an *exec.Cmd
		fmt.Printf("\n==>readSensors. tick at %v \n", t)
		cmd := exec.Command("/usr/bin/node", "test.js")

		// Stdout buffer
		cmdOutput := &bytes.Buffer{}

		// Only output the command's stdout
		// Attach buffer to command
		cmd.Stdout = cmdOutput

		// Execute command
		printCommand(cmd)
		err := cmd.Run() // will wait for command to return
		if err != nil {
			printError(err)
			ctrl <- 1
		}
		s <- cmdOutput.Bytes()
	}
}

func consumeSensorValues(s chan []byte, ctrl chan int) {
	//var dat SensorValue
	for {
		result := <-s
		fmt.Printf("\n==>got sensor value. %v \n", time.Now())
		printOutput(result)

		readout, err := dal.NewReadoutFromJson(result)
		if err != nil {
			//panic(err)
			printError(err)
			ctrl <- 1
			continue
		}
		fmt.Println(readout)
		if err := readout.Insert(); err != nil {
			printError(err)
			ctrl <- 1
			continue
		}

		ctrl <- 0
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}

func main() {
	fmt.Printf("==>main\n")
	c := make(chan []byte)
	ctrl := make(chan int)
	go readSensors(c, ctrl)
	go consumeSensorValues(c, ctrl)

	for {
		result := <-ctrl
		fmt.Printf("==> Exit: %v\n", result)
	}
}
