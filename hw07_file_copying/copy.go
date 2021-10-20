package main

import (
	"errors"
	"io"
	"io/fs"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrInfoFile              = errors.New("unable to get information about the file")
	ErrNotFoundFile          = errors.New("file reading error")
	ErrCopy                  = errors.New("copy error")
)

type Helper struct {
	from *os.File
	to   *os.File
	info fs.FileInfo
}

func (h *Helper) Open(fromPath string, offset int64) error {
	f, err := os.Open(fromPath)
	if err != nil {
		return ErrNotFoundFile
	}

	fi, err := f.Stat()
	if err != nil {
		return ErrInfoFile
	}

	if fi.IsDir() || fi.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset >= fi.Size() {
		return ErrOffsetExceedsFileSize
	}

	f.Seek(offset, 0)

	h.from = f
	h.info = fi

	return nil
}

func (h *Helper) Create(toPath string) error {
	out, err := os.Create(toPath)
	if err != nil {
		return ErrUnsupportedFile
	}

	h.to = out

	return nil
}

func (h *Helper) Copy(limit int64) error {
	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(h.from)

	if _, err := io.CopyN(h.to, barReader, limit); err != nil {
		return ErrCopy
	}

	bar.Finish()

	return nil
}

func Copy(fromPath string, toPath string, offset, limit int64) error {
	helper := Helper{}
	err := helper.Open(fromPath, offset)
	defer helper.from.Close()
	defer helper.to.Close()

	if err != nil {
		return err
	}

	countToCopy := helper.info.Size() - offset

	if limit > 0 && countToCopy > limit {
		countToCopy = limit
	}

	err = helper.Create(toPath)

	if err != nil {
		return err
	}

	err = helper.Copy(countToCopy)

	if err != nil {
		return err
	}

	return nil
}
