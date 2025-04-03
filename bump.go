package main

import (
	"regexp"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/Masterminds/semver"
)

type Rule int64

const (
	Major Rule = iota
	Minor
	Patch
	None
)

var PatchTypes = []string{"fix", "perf", "refactor", "deps"}
var MinorTypes = []string{"feat"}

func BumpRule(commitMessages []string) Rule {
	var rule Rule = None

	for _, message := range commitMessages {
		isConvential, isBreakingChange, commitType := getCommitType(cleanMessagePrefix(message))

		if !isConvential {
			continue
		}
		if isBreakingChange {
			return Major
		}
		if slices.Contains(PatchTypes, commitType) && rule == None {
			rule = Patch
		}
		if slices.Contains(MinorTypes, commitType) {
			rule = Minor
		}
	}
	return rule
}

func (rule Rule) String() string {
	switch rule {
	case Major:
		return "major"
	case Minor:
		return "minor"
	case Patch:
		return "patch"
	}
	return ""
}

func (rule Rule) Bump(version *semver.Version) *semver.Version {
	switch rule {
	case Major:
		newVerion := version.IncMajor()
		return &newVerion
	case Minor:
		newVerion := version.IncMinor()
		return &newVerion
	case Patch:
		newVerion := version.IncPatch()
		return &newVerion
	}
	return version
}

func cleanMessagePrefix(message string) string {
	messageParts := strings.Split(message, " ")
	if match, _ := regexp.MatchString("\\[.*\\]", messageParts[0]); match {
		return strings.Join(messageParts[1:], " ")
	}
	return message
}

func getCommitType(commitMessage string) (bool, bool, string) {
	re := regexp.MustCompile(`(?P<commit_type>^[a-zA-Z]+)\(?`)
	groups := re.FindStringSubmatch(commitMessage)
	commitTypeIndex := re.SubexpIndex("commit_type")

	if commitTypeIndex == -1 {
		return false, false, ""
	}

	isBreakingChange := strings.Contains(commitMessage, "!") || strings.Contains(commitMessage, "BREAKING CHANGE")

	return true, isBreakingChange, groups[commitTypeIndex]
}
