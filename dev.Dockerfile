FROM golang:1.18.3
WORKDIR /app

# Install air for live reload
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]