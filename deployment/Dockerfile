# syntax=docker/dockerfile:1

ARG GO_VERSION=1.24.0
ARG BIN_FILE=/bin/server
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

RUN apt-get update && \
    apt-get install -y protobuf-compiler

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go install tool

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download -x

ARG TARGETARCH

ENV OUTPUT_FILE=/bin/server
ENV GOARCH=$TARGETARCH

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=.,rw \
    go tool task generate

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=.,rw \
    go tool task build-bin

FROM alpine:latest AS final

RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

COPY --from=build /bin/server /bin/

EXPOSE 50051

ENTRYPOINT [ "/bin/server" ]
