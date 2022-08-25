package main_test

import (
	"fmt"
	"testing"

	"github.com/go-git/go-git/v5/plumbing/object"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGitSemver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GitSemver Suite")
}

type FakeCommitIter struct {
	Commits     []*object.Commit
	ForcedError bool
	iter        int
}

func (f *FakeCommitIter) Next() (*object.Commit, error) {
	if f.iter >= len(f.Commits) || f.ForcedError {
		return nil, fmt.Errorf("No commits left")
	}
	commit := f.Commits[f.iter]
	f.iter += 1
	return commit, nil
}

func (f *FakeCommitIter) ForEach(commitFunc func(*object.Commit) error) error {
	for _, commit := range f.Commits {
		err := commitFunc(commit)
		if err != nil {
			return err
		}
	}
	return nil
}

func (f *FakeCommitIter) Close() {
	f.Commits = []*object.Commit{}
	f.iter = 0
}
