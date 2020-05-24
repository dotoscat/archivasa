package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	config := Config{}
	configPath := filepath.Join(cwd, "config.txt")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalln(configPath, "not found. Aborting", configNotFound)
	}
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("Error opening config file: ", err)
	}
	lines := strings.Split(string(configFile), "\n")
	for _, row := range lines {
		fieldValue := strings.Split(row, "=")
		field := strings.TrimSpace(fieldValue[0])
		value := strings.TrimSpace(fieldValue[1])
		switch field {
		case "title":
			config.Title = value
		case "postsperpage":
			intValue, err := strconv.ParseInt(value, 10, 32)
			if err != nil {
				fmt.Println("Error with", field, ":", value, " cannot be converted to int")
			}
			config.PostsPerPage = int(intValue)
		}
	}
	return config
}
