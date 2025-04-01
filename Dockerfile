FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go build -o opsvg

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/opsvg .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/icons ./icons

EXPOSE 8080
CMD ["./opsvg"] 