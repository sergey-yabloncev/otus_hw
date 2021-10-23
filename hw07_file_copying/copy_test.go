package main

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

type test struct {
	from   string
	to     string
	offset int64
	limit  int64
	err    error
	result string
}

func compare(originalFilePath string, copyFilePath string) (bool, error) {
	originalFile, err := os.ReadFile(originalFilePath)
	if err != nil {
		return false, err
	}

	copyFile, err := os.ReadFile(copyFilePath)
	if err != nil {
		return false, err
	}

	return bytes.Equal(originalFile, copyFile), nil
}

func TestSuccess(t *testing.T) {
	f, err := os.CreateTemp("", "out.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(f.Name()) // clean up

	for _, tst := range [...]test{
		{
			from:   "testdata/input.txt",
			to:     "out.txt",
			result: "testdata/out_offset0_limit0.txt",
		},
		{
			from:   "testdata/input.txt",
			to:     "out.txt",
			limit:  10,
			result: "testdata/out_offset0_limit10.txt",
		},
		{
			from:   "testdata/input.txt",
			to:     "out.txt",
			limit:  1000,
			result: "testdata/out_offset0_limit1000.txt",
		},
		{
			from:   "testdata/input.txt",
			to:     "out.txt",
			limit:  10000,
			result: "testdata/out_offset0_limit10000.txt",
		},
		{
			from:   "testdata/input.txt",
			to:     "out.txt",
			limit:  1000,
			offset: 100,
			result: "testdata/out_offset100_limit1000.txt",
		},
		{
			from:   "testdata/input.txt",
			to:     "out.txt",
			limit:  1000,
			offset: 6000,
			result: "testdata/out_offset6000_limit1000.txt",
		},
	} {
		t.Run("ok", func(t *testing.T) {
			err := Copy(tst.from, tst.to, tst.offset, tst.limit)
			require.Nil(t, err)

			filesEqual, err := compare(tst.to, tst.result)
			require.Nil(t, err)

			require.True(t, filesEqual)
		})
	}
}

func TestFail(t *testing.T) {
	for _, tst := range [...]test{
		{
			from: "file.txt",
			err:  ErrNotFoundFile,
		},
		{
			from:   "testdata/input.txt",
			offset: 1000000,
			err:    ErrOffsetExceedsFileSize,
		},
		{
			from: "testdata",
			err:  ErrUnsupportedFile,
		},
	} {
		t.Run("fail", func(t *testing.T) {
			err := Copy(tst.from, tst.to, tst.offset, tst.limit)
			require.Equal(t, err, tst.err)
		})
	}
}
