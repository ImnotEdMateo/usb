# Build stage
FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go build -o ubs .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=build /app/ubs .

# Expose the port specified by the environment variable
EXPOSE ${UBS_PORT}

# Set environment variable
ENV UBS_PORT=8080

CMD ["./ubs"]
