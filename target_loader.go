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
	"errors"
	"fmt"
	"github.com/smallfish/simpleyaml"
	"io/ioutil"
	"strings"
)

func LoadTargets(filename string) ([]Target, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	yaml, err := simpleyaml.NewYaml(data)
	if err != nil {
		return nil, err
	}

	m, err := yaml.Map()
	if err != nil {
		return nil, err
	}

	var targets []Target

	for key, value := range m {
		value := value.([]interface{})
		commands := make([]Runnable, len(value))
		targetName := simplifyTargetName(key.(string))
		targetParams, targetParamErr := findParametersInTargetName(key.(string))
		if targetParamErr != nil {
			return nil, targetParamErr
		}

		for i := range value {
			commandString := value[i].(string)
			if commandString[0] == "="[0] {
				panic("Subtargets not implemented yet!")
			}
			commands[i] = Command{commandString}
			// commands[i] = Command{value[i].(string)}
		}
		targets = append(targets, Target{targetName, commands, targetParams})
	}

	// Routine for subtargets
	// for targetIndex := range targets {
	// 	for commandIndex := range targets[targetIndex].Commands {
	// 		currentCommand := targets[targetIndex].Commands[commandIndex].(string)
	// 		if currentCommand[0] != "="[0] {
	// 			continue
	// 		}

	// 		subTargetName := currentCommand[1:]
	// 		subTarget, err := FindTargetByTargetName(subTargetName, targets)
	// 		if err != nil {
	// 			fmt.Printf("Error: Using non-existent target \"%v\" as subtarget.\n", subTargetName)
	// 			fmt.Println("Valid targets are:")
	// 			for tIndex := range targets {
	// 				fmt.Println("  -", targets[tIndex].Name)
	// 			}
	// 			return nil, err
	// 		}

	// 		targets[targetIndex].Commands[commandIndex] = subTarget
	// 	}
	// }

	return targets, nil
}

func findParametersInTargetName(rawName string) ([]string, error) {
	paramStartIndex := strings.Index(rawName, "[")
	paramEndIndex := strings.Index(rawName, "]")
	if paramStartIndex == -1 {
		return make([]string, 0), nil
	}

	if paramEndIndex == -1 {
		return nil, errors.New(fmt.Sprintf("Incorrect syntax for target %v", rawName))
	}

	return strings.Split(rawName[paramStartIndex+1:paramEndIndex], ","), nil
}

func simplifyTargetName(rawName string) string {
	if rawName[0] == "="[0] {
		rawName = rawName[1:]
	}

	paramStartIndex := strings.Index(rawName, "[")
	if paramStartIndex == -1 {
		return rawName
	}
	return rawName[:paramStartIndex]
}
