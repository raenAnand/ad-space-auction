# Use the official Go image as the base image
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o supply

EXPOSE 8000

CMD ["./supply"]
