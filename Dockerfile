FROM golang:1.18.3 AS builder
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build main.go

FROM scratch

COPY --from=builder /app/main .

CMD ["./main"]





