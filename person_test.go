package person

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestReadPerson(t *testing.T) {
	got := ReadPeople()
	wantBytes, err := ioutil.ReadFile(filepath.Join("testdata/people.txt"))
	want := string(wantBytes)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}
