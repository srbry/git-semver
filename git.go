package main

import (
	"sort"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func LatestSemverTag(tags storer.ReferenceIter) string {
	var (
		semverTags     []*semver.Version
		semverMetadata map[*semver.Version]string
	)
	semverMetadata = make(map[*semver.Version]string)
	tags.ForEach(func(tag *plumbing.Reference) error {
		tagName := tag.Name().Short()
		version, err := semver.NewVersion(tagName)
		if err != nil {
			return nil
		}
		if version.Prerelease() != "" {
			return nil
		}
		semverTags = append(semverTags, version)
		semverMetadata[version] = tagName
		return nil
	})
	if len(semverTags) == 0 {
		return ""
	}
	sort.Sort(semver.Collection(semverTags))
	return semverMetadata[semverTags[len(semverTags)-1]]
}

func GetTagRef(repo *git.Repository, tag *plumbing.Reference) plumbing.Hash {
	tagRef := tag.Hash()
	obj, err := repo.TagObject(tagRef)
	if err != nil {
		return tagRef
	}
	tagCommit, err := obj.Commit()
	if err != nil {
		return tagRef
	}
	return tagCommit.Hash
}
