FROM golang:1.18
WORKDIR /go/src/mks-login-api-docker
COPY . .
RUN go build -o /bin/server cmd/web/main.go
CMD ["/bin/server"]
