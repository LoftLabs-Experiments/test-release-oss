# Build program
FROM golang:1.24 AS builder

WORKDIR /testapp-dev
ARG TARGETOS
ARG TARGETARCH
ARG BUILD_VERSION=dev
ARG TELEMETRY_PRIVATE_KEY=""

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# Copy the go source
COPY cmd/testapp cmd/testapp
COPY cmd/testctl cmd/testctl
COPY pkg/ pkg/
COPY config/ config/

ENV GO111MODULE=on
ENV DEBUG=true

# create and set GOCACHE now, this should slightly speed up the first build inside of the container
RUN mkdir -p /.cache /.config
ENV GOCACHE=/.cache
ENV GOENV=/.config

# Build cmd
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} GO111MODULE=on go build -ldflags "-X github.com/LoftLabs-Experiments/test-release-oss/pkg/telemetry.SyncerVersion=$BUILD_VERSION -X github.com/LoftLabs-Experiments/test-release-oss/pkg/telemetry.telemetryPrivateKey=$TELEMETRY_PRIVATE_KEY" -o /testapp cmd/testapp/main.go

ENTRYPOINT ["go", "run", "cmd/testapp/main.go", "start"]

# we use alpine for easier debugging
FROM alpine:3.22

# install runtime dependencies
RUN apk add --no-cache ca-certificates zstd tzdata

# Set root path as working directory
WORKDIR /

COPY --from=builder /testapp .

ENTRYPOINT ["/testapp", "start"]
