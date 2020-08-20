package main

// Package is called aw
import (
	"example.com/banana/bitbucketRepositories/bitbucketClient"
	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	for _, repo := range bitbucketClient.GetRepositories() {
		wf.NewItem(repo.Name).
			Arg(repo.Url).
			Subtitle(repo.Url).
			UID(repo.Name).
			Valid(true)
	}
	wf.SendFeedback()
}

func main() {
	wf.Args()
	wf.Run(run)
}
