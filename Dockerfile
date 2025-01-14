# Build stage
FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go build -o ubs .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=build /app/ubs .
COPY --from=build /app/static /app/static
RUN mkdir /app/uploads && chown -R appuser:appgroup /app/uploads

USER appuser

# Optional: Hardcode a default port
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD wget --spider --quiet http://localhost:8080/ || exit 1

CMD ["./ubs"]
