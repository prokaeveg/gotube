FROM golang:1.23.2

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Компилируем проект
RUN go build -o gotube

RUN chmod a+x ./gotube

CMD ["./gotube"]