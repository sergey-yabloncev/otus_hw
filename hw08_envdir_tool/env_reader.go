package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
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

		environment[d.Name()] = setEnvValue(path)

		return nil
	}

	if err := filepath.WalkDir(dir, walkDirFunc); err != nil {
		fmt.Errorf("Error read dir: %#v", err)
		return nil, err
	}

	return environment, nil
}

// Create env data
func setEnvValue(path string) EnvValue {
	needRemove := false

	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Printf("Error open file: %#v ", err)
		os.Exit(1)
	}

	size, err := getFileSize(*file)
	if err != nil {
		fmt.Printf("Error get file size: %#v ", err)
		os.Exit(1)
	}

	if size == 0 {
		needRemove = true
	}

	reader := bufio.NewReader(file)

	value, err := reader.ReadBytes('\n')

	if err != nil && err != io.EOF {
		fmt.Printf("Error read file %v: %#v ", path, err)
		os.Exit(1)
	}

	return EnvValue{
		sanitize(value),
		needRemove,
	}
}

func getFileSize(file os.File) (int64, error) {
	size, er := file.Stat()

	if er != nil {
		return 0, er
	}

	return size.Size(), nil
}

// sanitize env value
func sanitize(envValue []byte) string {
	envValue = bytes.TrimRight(envValue,"\n")
	envValue = bytes.TrimRight(envValue," ")
	envValue = bytes.ReplaceAll(envValue, []byte{0x00}, []byte{'\n'})

	return string(envValue)
}
