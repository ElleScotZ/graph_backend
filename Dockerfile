FROM golang:latest

ENV PORT=8080

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod download && go mod verify

RUN go build -o backend .

EXPOSE $PORT

ENTRYPOINT [ "/app/backend" ]