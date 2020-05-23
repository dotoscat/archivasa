package main

import (
	"log"
	"os"
	"path/filepath"
)

const configNotFound = `
	Create a 'config.txt' file in your current working directory, with these content (for example):

	title = MyBlog
	postsperpage = 3
`

// Config represents a config file
type Config struct {
	Title        string
	PostsPerPage int
}

// ReadConfigFile reads a config file from the current working directory
func ReadConfigFile(cwd string) Config {
	configPath := filepath.Join(cwd, "config.txt")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalln(configPath, "not found. Aborting", configNotFound)
	}
	return Config{}
}
