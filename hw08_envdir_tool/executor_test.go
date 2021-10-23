package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	result := RunCmd([]string{"ls", "-a"}, nil)
	require.Equal(t, 0, result)
}
