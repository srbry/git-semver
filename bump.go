package main

import (
	"regexp"
	"strings"

	"golang.org/x/exp/slices"

	"github.com/Masterminds/semver"
	"github.com/leodido/go-conventionalcommits"
	"github.com/leodido/go-conventionalcommits/parser"
)

type Rule int64

const (
	Major Rule = iota
	Minor
	Patch
	None
)

var PatchTypes = []string{"fix", "perf", "refactor"}
var MinorTypes = []string{"feat"}

func BumpRule(commitMessages []string) Rule {
	var rule Rule = None
	commitParser := parser.NewMachine([]conventionalcommits.MachineOption{
		conventionalcommits.WithTypes(conventionalcommits.TypesConventional),
		conventionalcommits.WithBestEffort(),
	}...)
	for _, message := range commitMessages {
		conventionalCommit, _ := commitParser.Parse([]byte(cleanMessagePrefix(message)))
		if conventionalCommit == nil {
			continue
		}
		if conventionalCommit.IsBreakingChange() {
			return Major
		}
		commitType := conventionalCommit.(*conventionalcommits.ConventionalCommit).Type
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
