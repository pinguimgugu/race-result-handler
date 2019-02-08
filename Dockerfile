FROM golang:1.11

RUN go get github.com/stretchr/testify/assert
RUN go get github.com/stretchr/testify/suite
RUN go get github.com/stretchr/testify/mock
RUN go get github.com/golang/dep/cmd/dep