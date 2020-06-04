package config

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/dotoscat/archivasa/pkg/util"
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
	Cwd          string
}

// ReadConfigFile reads a config file from the current working directory
func Read(cwd string) Config {
	config := Config{}
	configPath := filepath.Join(cwd, "config.txt")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalln(configPath, "not found. Aborting", configNotFound)
	}
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("Error opening config file: ", err)
	}
	pairs := util.ReadConfigFromString(string(configFile))
	title, _ := pairs.GetString("title")
	config.Title = title
	ppp, _ := pairs.GetInt("postsperpage")
	config.PostsPerPage = ppp
	config.Cwd = cwd
	return config
}
