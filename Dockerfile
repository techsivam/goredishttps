# Build stage
FROM golang:1.19-alpine AS build

ENV REDIS_HOST=redis
ENV REDIS_PORT=6379
ENV REDIS_PASSWORD=
ENV REDIS_DB=0

WORKDIR /src
COPY . .

RUN go mod init techsivam16/dockertest
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Final stage
FROM alpine:3.14

# Set the working directory
WORKDIR /app

# Copy the binary and certificate files
COPY --from=build /src/app /app/app
COPY cert.pem /app/cert.pem
COPY key.pem /app/key.pem

ENTRYPOINT ["/app/app"]
