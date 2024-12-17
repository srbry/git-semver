package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/urfave/cli/v2"
)

func getGitInfo() ([]string, *plumbing.Reference, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, nil, err
	}
	tags, _ := repo.Tags()
	ref, err := repo.Tag(LatestSemverTag(tags))
	if err != nil {
		return nil, nil, err
	}
	tagRef := GetTagRef(repo, ref)
	commits, err := repo.Log(&git.LogOptions{From: tagRef})
	if err != nil {
		return nil, nil, err
	}
	var commitMessages []string

	commits.ForEach(func(commit *object.Commit) error {
		commitMessages = append(commitMessages, commit.Message)
		return nil
	})
	return commitMessages, ref, nil
}

func main() {
	app := &cli.App{
		Name:  "git-semver",
		Usage: "Bump semver versions based on git tags",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "rule",
				Value: false,
				Usage: "Output the bump rule instead of semver version",
			},
		},
		Action: func(cCtx *cli.Context) error {
			commitMessages, ref, err := getGitInfo()
			if err != nil {
				return err
			}
			rule := BumpRule(commitMessages)
			if cCtx.Bool("rule") {
				fmt.Println(rule.String())
			} else {
				version, _ := semver.NewVersion(ref.Name().Short())
				fmt.Println(rule.Bump(version).String())
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
