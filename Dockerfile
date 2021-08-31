FROM golang:alpine3.14 as builder

RUN mkdir build

WORKDIR build

COPY go.mod go.sum ./
COPY pkg pkg
COPY cmd cmd

RUN go build -o prq ./cmd/prq/*.go

FROM alpine:3.14.2

COPY --from=builder /go/build/prq /usr/bin/prq

ENTRYPOINT ["prq"]
CMD "--help"



