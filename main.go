package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	config := Config{}
	config.init()
	if len(os.Args) < 2 {
		log.Fatal("please provide folder")
	}
	repo := os.Args[1]
	pathToRepo := path.Join(getHomeDir(), config.rootDir, repo)

	path := path.Clean(strings.Replace(pathToRepo, "~", "", 1))

	_, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	openRepo(path)
}

func openRepo(repo string) {
	cmd := exec.Command("code", repo)
	err := cmd.Run()
	handleError(err)
}
