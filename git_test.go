package main_test

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/storer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	git_semver "github.com/srbry/git-semver"
)

var _ = DescribeTable("Find the latest semver tag",
	func(tags storer.ReferenceIter, tag string) {
		Expect(git_semver.LatestSemverTag(tags)).To(Equal(tag))
	},
	Entry("It returns that latest semver tag",
		storer.NewReferenceSliceIter([]*plumbing.Reference{
			plumbing.NewReferenceFromStrings("1.2.1", "hash"),
			plumbing.NewReferenceFromStrings("1.1.1", "hash"),
			plumbing.NewReferenceFromStrings("1.3.1", "hash"),
		}),
		"1.3.1"),
	Entry("It returns that latest semver tag, ignoring non semvers",
		storer.NewReferenceSliceIter([]*plumbing.Reference{
			plumbing.NewReferenceFromStrings("1.2.1", "hash"),
			plumbing.NewReferenceFromStrings("1.1.1", "hash"),
			plumbing.NewReferenceFromStrings("foo", "hash"),
		}),
		"1.2.1"),
	Entry("It returns that latest semver tag, handling 'v' prefixes",
		storer.NewReferenceSliceIter([]*plumbing.Reference{
			plumbing.NewReferenceFromStrings("1.2.1", "hash"),
			plumbing.NewReferenceFromStrings("v3.6.2", "hash"),
			plumbing.NewReferenceFromStrings("1.1.1", "hash"),
		}),
		"v3.6.2"),
	Entry("It handles no existing semvers",
		storer.NewReferenceSliceIter([]*plumbing.Reference{
			plumbing.NewReferenceFromStrings("foo", "hash"),
			plumbing.NewReferenceFromStrings("bar", "hash"),
			plumbing.NewReferenceFromStrings("baz", "hash"),
		}),
		""),
)

var _ = Describe("#CommitMessagesSince", func() {
	var (
		commits = []*object.Commit{
			{
				Hash:    plumbing.NewHash("1111111111111111"),
				Message: "first commit",
			},
			{
				Hash:    plumbing.NewHash("2222222222222222"),
				Message: "second commit",
			},
			{
				Hash:    plumbing.NewHash("3333333333333333"),
				Message: "third commit",
			},
		}
		commitIter *FakeCommitIter
	)

	BeforeEach(func() {
		commitIter = &FakeCommitIter{Commits: commits}
	})

	Context("when there is a matching ref", func() {
		It("returns the commits up to the matching ref", func() {
			commitMessages := git_semver.CommitMessagesSince(commits[1].Hash, commitIter)
			Expect(commitMessages).To(Equal([]string{"first commit"}))
		})
	})

	Context("when there is no matching ref", func() {
		It("returns all the commits", func() {
			commitMessages := git_semver.CommitMessagesSince(plumbing.NewHash("no match"), commitIter)
			Expect(commitMessages).To(Equal([]string{
				"first commit",
				"second commit",
				"third commit",
			}))
		})
	})
})
