FROM golang:1.25.2
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates

WORKDIR /opt/lktr

COPY . .d
RUN go mod tidy
RUN go build -o bin/lktr ./cmd/lktr

FROM debian:bookworm
RUN apt-get update && \
    apt-get install -y ca-certificates && \
    update-ca-certificates
COPY --from=0 /opt/lktr/bin/lktr /opt/lktr

ENTRYPOINT ["/opt/lktr"]