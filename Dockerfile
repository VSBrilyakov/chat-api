FROM golang:1.25

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# build go app
RUN go mod download
RUN go build -o chat-app ./cmd/main.go

CMD ["./chat-app"]