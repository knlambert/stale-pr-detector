# PRQ

PRQ stands for Pull Requests Query. The command line only supports a stale feature, but it could be extended with more
if required.

# Usage

## Build with Golang 1.17

```bash
go build -o prq ./cmd/prq/*.go
mv prq /usr/local/bin/prq
```

Or directly :

```bash
go run cmd/prq/*.go stale \
  --repositories=https://github.com/emissary-ingress/emissary \
  --no-activity-since=14d --format=text
```

To run the test suite :

```bash
go test ./pkg/... -v -covermode=atomic -coverprofile cover.out
```

Execute this command to see coverage:

```bash
go tool cover -html=cover.out
```

## Build with docker

```bash
docker build . -t prq
docker run -t prq --help
```

To provide a github token, use the `GITHUB_ACCESS_TOKEN` env var.
For example with docker:

```bash
docker run \
  -e GITHUB_ACCESS_TOKEN=xxxx \
  -t prq [...]
```

## Commands

### help

To get the list of commands and options :

```bash
prq --help
```

### stale

The stale command is designed to extract PRs with no activity.

All the PRs from one repository, using JSON format, with 14 days without any activity. All that with the label `size/XS`
.

```bash
prq stale \
  --labels="stale" \
  --repositories=https://github.com/emissary-ingress/emissary \
  --no-activity-since=14d \
  --format=json | jq
```

# Possible improvements

* Better retry with an exponential backoff.
* Better test coverage (errors, ...).
* More git vendors.
* Other commands than stale for other use cases.
* Multi-thread crawling for better performance.
