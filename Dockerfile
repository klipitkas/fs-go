FROM golang:1.13-alpine as base

WORKDIR /app
COPY . .
RUN go get ./... && go build -o fs-go .

FROM alpine:3.10

VOLUME /data
WORKDIR /app
COPY --from=base /app/fs-go .
EXPOSE 8080

ENTRYPOINT ["/app/fs-go", "-dir", "/data"]