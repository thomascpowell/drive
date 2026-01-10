FROM golang:1.24.3-alpine

WORKDIR /app
RUN mkdir -p ./data
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o backend .

EXPOSE 8080
CMD ["./backend"]
