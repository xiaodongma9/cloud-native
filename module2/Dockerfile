FROM golang:1.17 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o module2 .


FROM alpine:3.11
COPY --from=builder /build/module2 /
EXPOSE 8080
ENTRYPOINT ["/module2"]
