package lib

import (
	"errors"
	"os"
)

var (
	ErrCouldNotSeek  = errors.New("could not seek the given file to the given offset")
	ErrInvalidOffset = errors.New("the given offset is not valid")
	ErrNoDataToRead  = errors.New("there is no new data to read")
	ErrCouldNotRead  = errors.New("could not read data form file")
)

func AppendDataToFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}

	return nil
}

func GetDataFromFile(filename string, offset int64) ([]byte, error) {
	f, err := os.OpenFile(filename, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	inf, err := f.Stat()
	if err != nil {
		return nil, err
	}

	sz := inf.Size() - offset
	if sz < 0 {
		return nil, ErrInvalidOffset
	} else if sz == 0 {
		return nil, ErrNoDataToRead
	}

	if n, err := f.Seek(offset, 0); err != nil || n != offset {
		return nil, ErrCouldNotSeek
	}

	data := make([]byte, sz)
	if n, err := f.Read(data); int64(n) != sz || err != nil {
		return nil, ErrCouldNotRead
	}

	return data, nil
}
