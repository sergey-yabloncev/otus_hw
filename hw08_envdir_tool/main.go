package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Errorf("Count args can't be <3, count your args: %#v", len(os.Args))
		return
	}

	env, err := ReadDir(os.Args[1])

	if err != nil {
		fmt.Errorf("failed read dir with environments: %#v", err.Error())
		return
	}

	os.Exit(RunCmd(os.Args[2:], env))
}
