FROM golang:latest

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
CMD ["air"]