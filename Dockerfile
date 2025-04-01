FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN go build -o opsvg

FROM alpine:latest

# Create a non-root user
RUN adduser -D -u 1000 appuser

# Create app directory and set permissions
WORKDIR /app
COPY --from=builder /build/opsvg .
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/static ./static
COPY --from=builder /build/icons ./icons

# Set proper permissions
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

EXPOSE 8080
CMD ["./opsvg"] 