package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestPersonReader(t *testing.T) {
	cmd := exec.Command("./person", "testdata/people.json")
	got, err := cmd.CombinedOutput()
	if err != nil {
		t.Error(err)
	}
	want, err := ioutil.ReadFile(filepath.Join("testdata/people.txt"))
	if bytes.Compare(got, want) != 0 {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}

}
