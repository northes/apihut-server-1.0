FROM golang:1.17 AS builder

WORKDIR /build

COPY . .

RUN export CGO_ENABLED=0 && go build -o bin/app .

FROM alpine

WORKDIR /apihut

COPY --from=builder /build/bin .
COPY etc etc
COPY templates templates
COPY data data
COPY static static

EXPOSE 8080

ENTRYPOINT ["./app"]