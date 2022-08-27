FROM golang:1

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
COPY pkg/ ./pkg/
CMD ["go", "run", "."]
