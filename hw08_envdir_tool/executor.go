package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	command := cmd[0]
	args := cmd[1:]
	exCmd := exec.Command(command, args...)
	exCmd.Env = os.Environ()
	exCmd.Stdout = os.Stdout
	exCmd.Stdin = os.Stdin
	exCmd.Stderr = os.Stderr

	for envKey, envValue := range env {
		var strEnv string
		if envValue.NeedRemove {
			strEnv = fmt.Sprintf("%s=", envKey)
		} else {
			strEnv = fmt.Sprintf("%s=%s", envKey, envValue.Value)
		}

		exCmd.Env = append(exCmd.Env, strEnv)
	}

	if err := exCmd.Run(); err != nil {
		exitError := &exec.ExitError{}
		if errors.As(err, &exitError) {
			return exitError.ExitCode()
		}
	}

	return 0
}
