package main_test

import (
	"fmt"

	"github.com/Masterminds/semver"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	git_semver "git-semver"
)

var _ = Describe("#Bump", func() {
	var _ = DescribeTable("Calculate version bumping rule",
		func(commitMessages []string, rule git_semver.Rule) {
			Expect(git_semver.BumpRule(commitMessages)).To(Equal(rule))
		},
		Entry("Features should be minor", []string{"feat: Example"}, git_semver.Minor),
		Entry("Fixes should be patch", []string{"fix: Example"}, git_semver.Patch),
		Entry("Refactor should be patch", []string{"refactor: Example"}, git_semver.Patch),
		Entry("Perf should be patch", []string{"perf: Example"}, git_semver.Patch),
		Entry("Build should be none", []string{"build: Example"}, git_semver.None),
		Entry("Ci should be none", []string{"ci: Example"}, git_semver.None),
		Entry("Chore should be none", []string{"chore: Example"}, git_semver.None),
		Entry("Docs should be none", []string{"docs: Example"}, git_semver.None),
		Entry("Revert should be none", []string{"revert: Example"}, git_semver.None),
		Entry("Style should be none", []string{"style: Example"}, git_semver.None),
		Entry("Test should be none", []string{"test: Example"}, git_semver.None),
		Entry("Breaking changes should be major", []string{"refactor: Example\n\nBREAKING CHANGE: breaks everything"}, git_semver.Major),
		Entry("Any type with exclamation should be major", []string{"refactor!: Example"}, git_semver.Major),
		Entry("The highest possible version should be returned, minor, major", []string{"feat: example", "refactor!: Example"}, git_semver.Major),
		Entry("The highest possible version should be returned, minor, patch", []string{"feat: example", "fix: Example"}, git_semver.Minor),
		Entry("The highest possible version should be returned, patch, minor", []string{"fix: example", "feat: Example"}, git_semver.Minor),
		Entry("The highest possible version should be returned, major, minor", []string{"feat!: example", "fix: Example"}, git_semver.Major),
		Entry("The highest possible version should be returned, major, patch, minor", []string{"feat!: example", "fix: example", "feat: Example"}, git_semver.Major),
		Entry("Unknown types are handled", []string{"foo: example"}, git_semver.None),
		Entry("Non conventional commits return no version", []string{"example"}, git_semver.None),
		Entry("Commit prefixes are supported", []string{"[JRA-123] fix: Example"}, git_semver.Patch),
	)
})

var _ = DescribeTable("Bump rules as strings",
	func(rule git_semver.Rule, expected string) {
		Expect(rule.String()).To(Equal(expected))
	},
	Entry("Major", git_semver.Major, "MAJOR"),
	Entry("Minor", git_semver.Minor, "MINOR"),
	Entry("Patch", git_semver.Patch, "PATCH"),
	Entry("None", git_semver.None, ""),
)

var _ = DescribeTable("Bump version based on rule",
	func(rule git_semver.Rule, version string, expected string) {
		versionSemver, err := semver.NewVersion(version)
		if err != nil {
			Fail(fmt.Sprintf("Version was invalid: %s", version))
		}
		expectedVersion, err := semver.NewVersion(expected)
		if err != nil {
			Fail(fmt.Sprintf("Expected version was invalid: %s", expectedVersion))
		}
		Expect(rule.Bump(versionSemver)).To(Equal(expectedVersion))
	},
	Entry("Major", git_semver.Major, "1.3.5", "2.0.0"),
	Entry("Minor", git_semver.Minor, "1.3.5", "1.4.0"),
	Entry("Patch", git_semver.Patch, "1.3.5", "1.3.6"),
	Entry("None", git_semver.None, "1.3.5", "1.3.5"),
)
