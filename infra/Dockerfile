FROM golang:alpine as builder
WORKDIR /go/example-golang
COPY . .

RUN ls -al
RUN go mod download
RUN go build ./cmd/main.go

FROM alpine
COPY --from=builder /go/example-golang/main /main

EXPOSE 9090
ENTRYPOINT ["/main"]
