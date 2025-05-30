FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /src
# Install Air
RUN go install github.com/air-verse/air@latest


COPY go.mod ./
RUN go mod download

COPY . .

# No build here—Air will rebuild on the fly

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

