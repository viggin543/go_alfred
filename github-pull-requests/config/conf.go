package config

import (
	"flag"
	"io/ioutil"
	"os"
)

var ExecutedByAlfred bool
var PathToLocalRepositories string
var GithubUser string
var GithubPass string
var Org string

func init() {
	isAlfred := flag.Bool("alfred", false, "a boolean value ")
	flag.Parse()
	ExecutedByAlfred = *isAlfred
	GithubUser = getEnvVarOrPanic("GITHUB_USER")
	GithubPass = getEnvVarOrPanic("GITHUB_PASS")
	Org = getEnvVarOrPanic("GITHUB_ORG")
	PathToLocalRepositories = getEnvVarOrPanic("REPOS_DIR")
}

func getEnvVarOrPanic(s string) string {
	if os.Getenv(s) != "" {
		return  os.Getenv(s)
	} else {
		panic("please set REPOS_DIR env var")
	}
}


func LocalRepos() []string {
	files, _ := ioutil.ReadDir(PathToLocalRepositories)
	var repos []string
	for _, file := range files {
		repos = append(repos, file.Name())
	}
	return repos
}
