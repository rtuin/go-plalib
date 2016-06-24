// The MIT License (MIT)

// Copyright (c) 2016 Richard Tuin

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package plalib

import (
	"bytes"
	// "fmt"
	"os/exec"
	"strings"
)

type Runnable interface {
	Run(params []string, stopRunning bool) bool
}

type Target struct {
	Name       string
	Commands   []Runnable
	Parameters []string
}

type Command struct {
	Command string
}

func (t Target) Run(params []string, stopRunning bool) bool {
	for i := range t.Commands {
		stopRunning = t.Commands[i].Run(params, stopRunning)
	}
	return stopRunning
}

func (c Command) Run(params []string, stopRunning bool) bool {
	if stopRunning {
		// fmt.Printf("\x1b[37;2m    . %v\x1b[0m\n", c.Command)
		return true
	}

	commandString := c.Command
	// if len(params) > 0 {
	// 	for index := range params {
	// 		commandString = strings.Replace(commandString, fmt.Sprintf("%%%v%%", target.Parameters[index]), params[index], -1)
	// 	}
	// }

	// fmt.Printf("    ⌛ %v", c.Command)

	cmd := exec.Command("sh", "-c", commandString)
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Errorf("Error running: %v", c.Command)
		strErrLines := strings.Split(stdErr.String(), "\n")
		if len(stdErr.String()) == 0 {
			strErrLines = []string{"[no output]"}
		}

		for lineIndex := range strErrLines {
			log.Errorf("%s", strErrLines[lineIndex])
		}
		return true
	}
	log.Infof("Success running: %v", c.Command)

	return false
}
