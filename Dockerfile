FROM golang:latest

WORKDIR /app
COPY . /app
RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
ADD ./database/my.cnf /etc/mysql/conf.d/my.cnf
CMD ["air"]
