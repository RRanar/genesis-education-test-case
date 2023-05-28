FROM golang:1.17-alpine as build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/swayne275/joke-web-server

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /joke-web-server .

FROM scratch

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /joke-web-server /joke-web-server

EXPOSE 5000

ENTRYPOINT ["/joke-web-server"]