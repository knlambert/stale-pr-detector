#!/bin/sh

rm -rf pkg/mock

mockgen -destination=./pkg/mock/github/client_mock.go \
  -package=github_mock \
  -source=./pkg/git/github/client.go goGithubSearch
