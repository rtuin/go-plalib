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
