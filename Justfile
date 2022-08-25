# This help screen
show-help:
  @just --list --unsorted

# Run the unit tests
test:
  go test -v ./...

# Install git-semver manually
install:
  go install .
