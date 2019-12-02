package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

const fileName = ".repo"

// Config ...
type Config struct {
	rootDir string
}

func (c *Config) init() {
	file := getConfig()
	parsedConfig := parseConfig(file)

	c.rootDir = parsedConfig["rootDir"]
}

func getHomeDir() string {
	homeDir, err := os.UserHomeDir()
	handleError(err)
	return homeDir
}

func getConfig() string {

	file, err := ioutil.ReadFile(path.Join(getHomeDir(), fileName))

	if err != nil {
		//TODO: make file
		log.Fatal(err)
	}
	return string(file)
}

func parseConfig(file string) map[string]string {
	options := strings.Split(strings.TrimSpace(file), "\n")
	config := make(map[string]string)
	for _, option := range options {
		x := strings.Split(option, "=")
		if len(x) > 1 {
			config[x[0]] = x[1]
		}

	}
	return config
}
