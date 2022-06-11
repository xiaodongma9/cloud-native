
FROM golang:1.16-alpine as first
WORKDIR /app

COPY *.go ./
RUN go mod init module2
RUN go build -o /module2

EXPOSE 8080


FROM scratch
COPY --from=first ./module2 /
EXPOSE 8080
ENTRYPOINT ["/module2"]