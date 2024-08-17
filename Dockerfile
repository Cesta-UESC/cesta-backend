FROM golang:1.22.4-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

RUN go install github.com/air-verse/air@latest

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

CMD ["air", "-c", ".air.toml"]
