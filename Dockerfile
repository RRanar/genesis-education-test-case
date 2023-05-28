FROM golang:1.17-alpine as build-stage

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/RRanar/genesis-education-test-case

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o /genesis-education-test-case .

FROM scratch

COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build-stage /genesis-education-test-case /genesis-education-test-case

EXPOSE 5000

ENTRYPOINT ["/genesis-education-test-case"]