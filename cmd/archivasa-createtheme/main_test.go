package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func createTempDir() (string, error) {
	dir, err := ioutil.TempDir("", "test-*")
	return dir, err
}

func removeTempDir(dir string) {
	os.RemoveAll(dir)
}

func TestBuildPath(t *testing.T) {
	files := map[string]string{
		filepath.Join("one", "one.txt"):  "hello",
		filepath.Join("one", "two.txt"):  "one",
		filepath.Join("/two", "one.txt"): "two",
	}
	const dir1 = "final"
	cases := map[string]string{
		filepath.Join(dir1, "one", "one.txt"): "hello",
		filepath.Join(dir1, "one", "two.txt"): "one",
		filepath.Join(dir1, "two", "one.txt"): "two",
	}
	results := buildPaths(dir1, files)
	for path, content := range results {
		if cases[path] != content {
			t.Errorf("%v not expected.", path)
		}
	}
}

func TestCopyFiles(t *testing.T) {
	dir, err := createTempDir()
	if err != nil {
		t.Fatal(err)
	}
	removeTempDir(dir)
}
