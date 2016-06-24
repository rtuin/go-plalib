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
)

func RunTargetByName(targetName string, targets []Target, stopRunning bool, params []string) *PlaError {
	target, error := FindTargetByTargetName(targetName, targets)
	if error != nil {
		log.Errorf("%v", *error)
		return error
	}

	// if len(params) < len(target.Parameters) {
	// missingParameters := target.Parameters[len(params):]
	// err := fmt.Sprintf("\x1b[31;2mCannot run \"%v\": Parameter \x1b[31;1m\x1b[31;4m%v\x1b[0m\x1b[31;2m not provided.\x1b[0m\n", target.Name, strings.Join(missingParameters, ", "))
	// fmt.Printf(err)
	// return errors.New(err)
	// }

	log.Infof("Running target \"%v\"", targetName)

	var failure = target.Run(params, stopRunning)
	if failure {
		return &PlaError{
			errors.New(fmt.Sprintf("Error running target \"%v\". Please check the logs for more information.", targetName)),
			TARGET_RUN_ERROR,
		}
	}
	return nil
}

func FindTargetByTargetName(targetName string, targets []Target) (Target, *PlaError) {
	for targetIndex := range targets {
		if targets[targetIndex].Name == targetName {
			return targets[targetIndex], nil
		}
	}

	return Target{}, &PlaError{
		errors.New(fmt.Sprintf("Invalid value: Target \"%v\" not present in Plafile.yml.", targetName)),
		TARGET_NOT_FOUND,
	}
}
