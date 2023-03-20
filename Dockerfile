FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct

WORKDIR /App
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go


EXPOSE 8080
ENTRYPOINT ["./main"]

