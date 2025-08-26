FROM golang:1.22-alpine AS builder

RUN apk add --no-cache make git

WORKDIR /go/src/app
COPY . . 
RUN make build

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["./kbot"]
