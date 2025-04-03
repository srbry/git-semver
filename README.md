# git-semver

A simple CLI that will bump semantic versions based on git tags and
[conventional commits](https://www.conventionalcommits.org/en/v1.0.0/).

## Install

```sh
go install github.com/srbry/git-semver
```

## Usage

**Note**: At present, `git-semver` will not actually add the new git tag to
your repository, it will simply output the next version/ rule for you to use
as you wish.

### Show the next semantic version

```sh
cd <my-git-repo>
git-semver
```

### Show the bump rule that will be used

```sh
cd <my-git-repo>
git-semver --rule
```

**Note**: This is particularly powerful when combinded with a tool like [poetry](https://python-poetry.org/) which
includes its own version bumping tool

```sh
cd <my_poetry_repo>
poetry version $(git-semver --rule)
```

## Release rules

At present, `git-semver` is slightly opinionated but aligs with how conventional commits intend to use semver:
<https://www.conventionalcommits.org/en/v1.0.0/#summary>

The supported types are taken from:
<https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional>

Any commits that are either not conventional, or do not have an associated rule are not treated as releases.

As well as the rules below. Any type (e.g `ci!`) that appends a `!` is treated as a breaking change and therefore a `MAJOR` release. Additionally any commit that contains the string `BREAKING CHANGE` is treated as a breaking change and `MAJOR` release.

| Commit type | Release |
| ----------- | ------- |
| feat        | `MINOR` |
| fix         | `PATCH` |
| perf        | `PATCH` |
| refactor    | `PATCH` |
| deps        | `PATCH` |
| build       | N/A     |
| chore       | N/A     |
| ci          | N/A     |
| docs        | N/A     |
| revert      | N/A     |
| style       | N/A     |
| test        | N/A     |

Please raise an issue if you disagree with any of our release rules, we would like to encourage conversation.

## Built With

  - [Contributor Covenant](https://www.contributor-covenant.org/) - Used
    for the Code of Conduct
  - [Creative Commons](https://creativecommons.org/) - Used to choose
    the license

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code
of conduct, and the process for submitting pull requests to us.

## License

This project is licensed under the [CC0 1.0 Universal](LICENSE.md)
Creative Commons License - see the [LICENSE.md](LICENSE.md) file for
details
