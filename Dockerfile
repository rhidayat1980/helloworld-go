FROM golang:1.13 as builder 

WORKDIR /app

COPY . /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o app

FROM alpine:3

RUN apk add --no-cache ca-certificates

COPY --from=builder /app/app /app

CMD ["/app"]