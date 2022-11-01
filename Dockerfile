# The base go image
FROM golang:latest

ENV PORT=8080

RUN mkdir /app

# Copy files to the /app directory
COPY . /app

WORKDIR /app

RUN go mod download && go mod verify

# Building a backend.exe in the /app directory
RUN go build -o backend .

EXPOSE $PORT

ENTRYPOINT [ "/app/backend" ]

# CMD go run main.go