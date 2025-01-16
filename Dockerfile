# Build stage
FROM golang:alpine AS build

WORKDIR /app
COPY . .
RUN go build -o usb .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app
COPY --from=build /app/usb .
COPY --from=build /app/static /app/static
RUN mkdir /app/uploads && chown -R appuser:appgroup /app/uploads

USER appuser

# OPTIONAL: Hardcore a default port
# EXPOSE 8080 

HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD wget --spider --quiet http://localhost:${PORT:-8080}/ || exit 1

CMD ["/bin/sh", "-c", "./usb --port ${PORT:-8080}"]
