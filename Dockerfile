# The base go image
FROM golang:latest

RUN mkdir /app

# Copy files to the /app directory
COPY . /app

WORKDIR /app

RUN go mod download

# Building a backend.exe in the /app directory
RUN go build -o backend .

ENTRYPOINT [ "/app/backend" ]

# EXPOSE 8080

# Run backend.exe
CMD [ "/app/backend" ]