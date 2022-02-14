package utils

import "testing"

func TestFileExists(t *testing.T) {
	path := "../testdata/mobydick.txt"
	want := true

	res, err := FileExists(path)
	if !res || err != nil {
		t.Fatalf(`FileExists("%s") = '%v'; want '%v'`, path, res, want)
	}

}

func TestFileDoesNotExist(t *testing.T) {
	path := "../testdata/missing.txt"
	want := false

	res, err := FileExists(path)
	if res || err == nil {
		t.Fatalf(`FileExists("%s") = '%v'; want '%v'`, path, res, want)
	}
}
