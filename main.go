package main

import (
	"flag"
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
)

var repo string

func main() {
	flag.StringVar(&repo, "repo", "", "repository url")
	flag.Parse()

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repo,
	})
	youSilly(err)

	ref, err := r.Head()
	youSilly(err)

	it, err := r.Log(&git.LogOptions{
		From: ref.Hash(),
		Order: git.LogOrderCommitterTime,
	})
	youSilly(err)

	commits := make([]string, 0)

	err = it.ForEach(func(c *object.Commit) error {
		commits = append(commits, c.Hash.String()[:8])
		return nil
	})
	youSilly(err)

	for i, j := 0, len(commits)-1; i < j; i, j = i+1, j-1 {
		commits[i], commits[j] = commits[j], commits[i]
	}

	for _, h := range commits {
		fmt.Println(h)
	}

}

func youSilly(err error) {
	if err != nil {
		panic(err)
	}
}
