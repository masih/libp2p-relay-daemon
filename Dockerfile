FROM golang:1.19-bullseye as BUILD

ARG GO_LIBP2P_RELAY_DAEMON_VERSION="0.3.0"
ENV CGO_ENABLED=0
RUN go install github.com/libp2p/go-libp2p-relay-daemon/cmd/libp2p-relay-daemon@v${GO_LIBP2P_RELAY_DAEMON_VERSION}

FROM gcr.io/distroless/static-debian11
COPY --from=BUILD /go/bin/libp2p-relay-daemon /usr/bin/

ENTRYPOINT ["/usr/bin/libp2p-relay-daemon"]
