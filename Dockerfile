FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN go build -o opsvg

FROM alpine:latest

# Create a non-root user
RUN adduser -D -u 1000 appuser

# Create app directory with full permissions
RUN mkdir -p /app && \
    chmod 777 /app

WORKDIR /app

# Copy files with proper ownership
COPY --from=builder --chown=appuser:appuser /build/opsvg .
COPY --from=builder --chown=appuser:appuser /build/templates ./templates
COPY --from=builder --chown=appuser:appuser /build/static ./static
COPY --from=builder --chown=appuser:appuser /build/icons ./icons

# Ensure all files are readable and executable
RUN chmod -R 755 /app

# Switch to non-root user
USER appuser

EXPOSE 8080
CMD ["./opsvg"] 