FROM golang:1.19.8-alpine AS builder

ENV GO111MODULE=on

RUN apk update && apk add git

RUN apk add make

RUN git clone --depth 1 https://github.com/Elementary1092/test

WORKDIR /test-app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make build

FROM alpine:3.17.3

WORKDIR /

COPY --from=builder /test-app/build /build

EXPOSE 8080

CMD ["/build/app"]
