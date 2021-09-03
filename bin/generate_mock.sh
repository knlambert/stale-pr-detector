#!/bin/sh

rm -rf pkg/mock

mockgen -destination=./pkg/mock/git.go \
  -package=pkg_mock \
  -source=./pkg/git.go GitClient

mockgen -destination=./pkg/mock/time.go \
  -package=pkg_mock \
  -source=./pkg/time.go TimeWrapper


mockgen -destination=./pkg/mock/formatter.go \
  -package=pkg_mock \
  -source=./pkg/formatter.go Formatter

mockgen -destination=./pkg/mock/git/github/client.go \
  -package=github_mock \
  -source=./pkg/git/github/client.go goGithubSearch timeWrapper
