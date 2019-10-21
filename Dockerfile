FROM golang:1.13-alpine as base

WORKDIR /app

COPY . .

RUN go get ./... && \
    go build -o fs-go .

FROM alpine:3.10

WORKDIR /data

COPY --from=base /app/fs-go /app/fs-go

ENTRYPOINT ["/app/fs-go"]

EXPOSE 8080
