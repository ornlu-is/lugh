FROM golang:1.19 AS builder

ADD . /banshee
WORKDIR /banshee

RUN make build

FROM debian:bookworm-slim
COPY --from=builder /banshee/bin/banshee /usr/bin/banshee

ENTRYPOINT ./usr/bin/banshee
