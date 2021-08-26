FROM golang:1.14 AS builder
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build
COPY . .
RUN go build -o library-rest-api -mod vendor .

FROM alpine:3.10
COPY --from=builder /build/library-rest-api /usr/local/bin
RUN chmod a+x /usr/local/bin/library-rest-api
COPY /migrations/ /migrations/

CMD ["library-rest-api"]