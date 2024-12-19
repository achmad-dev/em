FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN apk add --no-cache dumb-init

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main ./cmd/server


FROM alpine:latest


WORKDIR /root/


COPY --from=builder /app/main .
COPY --from=builder ["/usr/bin/dumb-init", "/usr/bin/dumb-init"]


ARG PORT
EXPOSE ${PORT}

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["./main"]