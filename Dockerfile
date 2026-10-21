FROM golang:1.26

WORKDIR /app

RUN go install github.com/air-verse/air@v1.62.0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]