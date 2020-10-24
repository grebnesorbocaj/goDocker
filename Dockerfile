FROM golang:alpine as builder
WORKDIR /build
ADD server.go /build/
RUN go build server.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/server /app/

CMD ["./server"]
