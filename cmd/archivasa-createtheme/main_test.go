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

const dir1 = "final"

type TestConfig struct {
	files map[string]string
	cases map[string]string
}

func giveMeConf() TestConfig {
	return TestConfig{
		files: map[string]string{
			filepath.Join("one", "one.txt"):  "hello",
			filepath.Join("one", "two.txt"):  "one",
			filepath.Join("/two", "one.txt"): "two",
		},
		cases: map[string]string{
			filepath.Join(dir1, "one", "one.txt"): "hello",
			filepath.Join(dir1, "one", "two.txt"): "one",
			filepath.Join(dir1, "two", "one.txt"): "two",
		},
	}
}

func TestBuildPath(t *testing.T) {
	config := giveMeConf()
	t.Log("Config:\n")
	t.Log(config)
	results := buildPaths(dir1, config.files)
	for path, content := range results {
		if config.cases[path] != content {
			t.Errorf("%v not expected. Got: %v; with content: %v", path, config.cases[path], content)
		}
	}
}

func TestMakeDirs(t *testing.T) {
	dir, err := createTempDir()
	if err != nil {
		t.Fatal(err)
	}
	removeTempDir(dir)
}

func TestCopyFiles(t *testing.T) {
	dir, err := createTempDir()
	if err != nil {
		t.Fatal(err)
	}
	removeTempDir(dir)
}
