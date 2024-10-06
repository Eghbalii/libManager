FROM golang:1.21-alpine AS build
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# Build the Go executables
RUN go build -o main ./main.go

#Runtime stage
FROM alpine:3.15
COPY --from=build /app/main /app/main

CMD ["app/main"]

EXPOSE 8080