package main

import (
	"bufio"
	"bytes"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	environment := make(Environment)

	walkDirFunc := func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() || strings.Contains(d.Name(), "=") {
			return nil
		}

		env, errSetEnv := setEnvValue(path)
		if errSetEnv != nil {
			return errSetEnv
		}

		environment[d.Name()] = env

		return nil
	}

	if err := filepath.WalkDir(dir, walkDirFunc); err != nil {
		return nil, err
	}

	return environment, nil
}

// Create env data.
func setEnvValue(path string) (EnvValue, error) {
	needRemove := false

	file, err := os.Open(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Cant close file")
		}
	}(file)

	if err != nil {
		return EnvValue{}, err
	}

	size, err := getFileSize(*file)
	if err != nil {
		return EnvValue{}, err
	}

	if size == 0 {
		needRemove = true
	}

	reader := bufio.NewReader(file)
	value, _ := reader.ReadBytes('\n')

	return EnvValue{
		sanitize(value),
		needRemove,
	}, nil
}

func getFileSize(file os.File) (int64, error) {
	size, er := file.Stat()

	if er != nil {
		return 0, er
	}

	return size.Size(), nil
}

// sanitize env value.
func sanitize(envValue []byte) string {
	envValue = bytes.TrimRight(envValue, "\n")
	envValue = bytes.TrimRight(envValue, " ")
	envValue = bytes.ReplaceAll(envValue, []byte{0x00}, []byte{'\n'})

	return string(envValue)
}
