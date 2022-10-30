# The base go image
FROM golang:latest

RUN mkdir /app

# Copy files to the /app directory
COPY . /app

WORKDIR /app

RUN go mod download && go mod verify

# Building a backend.exe in the /app directory
RUN go build -o backend .

EXPOSE 8080

ENTRYPOINT [ "/app/backend" ]

# # Run backend.exe
# CMD go run main.go