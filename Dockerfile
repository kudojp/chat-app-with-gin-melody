# Builder Container
FROM golang:latest AS builder
RUN apt update
WORKDIR /go/src/app
COPY . .
# build for alopine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

# Exec Container
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/app/main .
COPY --from=builder /go/src/app/view ./view
CMD ["./main"]