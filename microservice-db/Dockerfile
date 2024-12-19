FROM golang:1.23

WORKDIR /app

COPY . .

RUN go mod download
RUN go mod tidy

RUN go build -o main ./cmd/api/main.go

EXPOSE 80

CMD [ "./main" ]
