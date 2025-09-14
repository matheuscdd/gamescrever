FROM golang:1.25

WORKDIR /api

RUN go install github.com/cespare/reflex@v0.3.1

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["reflex", "-r", "\\.go$", "-s", "--", "go", "run", "."]
