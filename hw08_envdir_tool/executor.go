package main

import (
	"fmt"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	// Place your code here.
	exCmd := exec.Command(cmd[0], cmd[1:]...)
	exCmd.Env = append(os.Environ())
	exCmd.Stdout = os.Stdout
	exCmd.Stdin = os.Stdin
	exCmd.Stderr = os.Stderr

	for envKey, envValue := range env {
		if envValue.NeedRemove {
			exCmd.Env = append(exCmd.Env, fmt.Sprintf("%s=", envKey))
		} else {
			exCmd.Env = append(exCmd.Env, fmt.Sprintf("%s=%s", envKey, envValue.Value))
		}
	}

	if err := exCmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode()
		}
	}

	return 0
}
