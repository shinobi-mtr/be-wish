package lib_test

import (
	"testing"

	"github.com/shinobi-mtr/be-wish/lib"
)

func TestAppendDataToFile(t *testing.T) {
	if err := lib.AppendDataToFile("../public/test.txt", []byte("this should work\n")); err != nil {
		t.Fatal(err)
	}

	if err := lib.AppendDataToFile("../public/test.txt", []byte("this should work\n")); err != nil {
		t.Fatal(err)
	}

	if err := lib.AppendDataToFile("/etc/hosts", []byte("127.0.0.1	localhost")); err == nil {
		t.Fatal("this /ets/hosts file should not be edited")
	}
}

func TestGetDataFromFile(t *testing.T) {
	_, err := lib.GetDataFromFile("../public/test.txt", 0)
	if err != nil {
		t.Fatal(err)
	}

	_, err = lib.GetDataFromFile("../public/test.txt", 10000)
	if err != lib.ErrInvalidOffset {
		t.Fatal(err)
	}
}
