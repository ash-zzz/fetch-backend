FROM golang:1.23.2-bookworm

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD [ "/godocker" ]