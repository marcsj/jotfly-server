# use our recommended go version
ARG GO_VERSION=1.13.8

# build the executable
FROM golang:${GO_VERSION}-alpine as builder

RUN apk add --no-cache ca-certificates git

# set the working directory
WORKDIR $GOPATH/src/github.com/marcsj/jotfly-server

# go modules
ENV GO111MODULE=on
ENV GOFLAGS="-mod=vendor"

# copy context
COPY . ./

# build executable to /app
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app \
    cmd/server/main.go

# leave building environment
FROM scratch

COPY --from=builder /usr/local/share/ca-certificates /usr/local/share/ca-certificates
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /app /app

EXPOSE 50051

# run the compiled binary
ENTRYPOINT ["/app"]