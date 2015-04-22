package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type SensorValue struct {
	Uuid string
	Date string
	Temp float64
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func readSensors(s chan []byte, ctrl chan int) {
	//TODO tick every 2 mins
	ticker := time.NewTicker(time.Millisecond * 7000)

	for t := range ticker.C {
		// Create an *exec.Cmd
		fmt.Printf("==>readSensors. tick at %v \n", t)
		cmd := exec.Command("/usr/bin/node", "test.js")
		// Stdout buffer
		cmdOutput := &bytes.Buffer{}
		// Attach buffer to command
		// Only output the command's stdout
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
	var dat SensorValue
	//var dat map[string]interface{}
	for {
		result := <-s
		printOutput(result)

		if err := json.Unmarshal(result, &dat); err != nil {
			panic(err)
		}
		fmt.Println(dat)

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
