FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod init skills-service && go mod tidy
RUN go build -o app

CMD ["./app"]
