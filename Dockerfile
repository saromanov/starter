FROM golang:1.15-alpine

ADD . /project
WORKDIR /project

RUN go build -o /bin/starter ./cmd/main.go
ENTRYPOINT [ "/bin/starter" ]

