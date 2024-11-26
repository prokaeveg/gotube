FROM golang:1.23.2

WORKDIR /usr/src/app
ENV CGO_ENABLED 0
RUN go install github.com/air-verse/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем проект
RUN go build -o gotube

CMD ["air", "-c", ".air.toml"]