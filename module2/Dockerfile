
FROM golang:1.16-alpine
WORKDIR /app

COPY *.go ./
RUN go mod init module2
RUN go build -o /module2

EXPOSE 8080

CMD [ "/module2" ]