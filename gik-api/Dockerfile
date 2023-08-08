FROM golang:1.18.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ /app

RUN go build -o /gik-api

EXPOSE 3001

CMD ["/gik-api"]