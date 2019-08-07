FROM golang:1.12 as builder
WORKDIR /go/api
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:latest
EXPOSE 8080
RUN apk --no-cache add ca-certificates
WORKDIR /api
COPY --from=builder /go/api/ .
RUN pwd
CMD ["./secret_test"]
