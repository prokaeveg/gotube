FROM golang:1.23.2

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем проект
RUN go build -o gotube

CMD ["air", "-c", ".air.toml"]