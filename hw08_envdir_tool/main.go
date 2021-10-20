package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Count args can't be <3, count your args: %#v", len(os.Args))
		return
	}

	env, err := ReadDir(os.Args[1])

	if err != nil {
		log.Fatalf("failed read dir with environments: %#v", err.Error())
		return
	}

	os.Exit(RunCmd(os.Args[2:], env))
}
