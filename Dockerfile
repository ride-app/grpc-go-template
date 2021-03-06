# syntax=docker/dockerfile:1

# Create .netrc file for private go module
FROM bufbuild/buf:1.4.0 as buf

ARG BUF_USERNAME ""

SHELL ["/bin/ash", "-o", "pipefail", "-c"]
RUN --mount=type=secret,id=BUF_TOKEN \
  buf registry login --username=$BUF_USERNAME --token-stdin < /run/secrets/BUF_TOKEN

# Build go binary
FROM golang:1.18-alpine as build

WORKDIR /go/src/app

COPY --from=buf /root/.netrc /root/.netrc
ENV GOPRIVATE=go.buf.build

COPY go.mod go.sum /
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app

# Run
FROM gcr.io/distroless/static:nonroot

WORKDIR /

COPY --from=build /go/bin/app .
COPY .env .env

EXPOSE 50051
CMD ["/app"]