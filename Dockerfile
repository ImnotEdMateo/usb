# Build stage
FROM golang:alpine AS build

RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

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
COPY --from=build /app/utils /app/utils
COPY --from=build /app/handlers /app/handlers
COPY --from=build /app/config /app/config
RUN mkdir /app/uploads && chown -R appuser:appgroup /app/uploads

USER appuser

# Expose the port specified by the environment variable
EXPOSE ${UBS_PORT}

CMD ["./ubs"]
