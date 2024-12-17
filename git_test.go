package main_test

import (
	"github.com/go-git/go-git/v5/plumbing"
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
