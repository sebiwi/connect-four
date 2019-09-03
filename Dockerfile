FROM golang:1.9

RUN mkdir -p /go/src/github.com/sebiwi/connect-four
WORKDIR /go/src/github.com/sebiwi/connect-four
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static"

FROM scratch
COPY --from=0 /go/src/github.com/sebiwi/connect-four/connect-four /connect-four
CMD ["/connect-four"]
