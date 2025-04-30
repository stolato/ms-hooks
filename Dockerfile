FROM golang:1.24.2

# Set destination for COPY
WORKDIR /app

COPY . .

ENV GOPROXY=https://goproxy.cn

RUN go mod download

RUN env GOOS=linux GOARCH=arm go build -o main cmd/main.go

EXPOSE 80

CMD ["./main"]