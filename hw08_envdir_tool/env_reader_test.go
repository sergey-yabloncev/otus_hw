package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadDirWithOutError(t *testing.T) {
	_, err := ReadDir("./testdata/env")
	require.NoError(t, err)
}

func TestReadBarEnv(t *testing.T) {
	env, _ := ReadDir("./testdata/env")

	require.Equal(t, "bar", env["BAR"].Value)
	require.Equal(t, false, env["BAR"].NeedRemove)
}

func TestReadEmptyEnv(t *testing.T) {
	env, _ := ReadDir("./testdata/env")

	require.Equal(t, "", env["EMPTY"].Value)
	require.Equal(t, false, env["EMPTY"].NeedRemove)
}

func TestReadFooEnv(t *testing.T) {
	env, _ := ReadDir("./testdata/env")

	require.Equal(t, "   foo\nwith new line", env["FOO"].Value)
	require.Equal(t, false, env["FOO"].NeedRemove)
}

func TestReadHelloEnv(t *testing.T) {
	env, _ := ReadDir("./testdata/env")

	require.Equal(t, "\"hello\"", env["HELLO"].Value)
	require.Equal(t, false, env["HELLO"].NeedRemove)
}

func TestReadUnsetEnv(t *testing.T) {
	env, _ := ReadDir("./testdata/env")

	require.Equal(t, "", env["UNSET"].Value)
	require.Equal(t, true, env["UNSET"].NeedRemove)
}
