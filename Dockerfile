FROM golang:alpine

RUN apk update && apk add git
COPY /src/ /go/src/

WORKDIR /go/src/

RUN go mod download \
&& GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main server.go

CMD [ "./main" ]

EXPOSE 8080

ENV GO111MODULE=on
