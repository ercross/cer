FROM golang:1.23 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /cer .

# Use a minimal base image for the runtime stage
FROM alpine:latest

# Add a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
WORKDIR /root/
COPY --from=builder /cer .

# Set permissions for the binary
RUN chmod +x /cer && chown appuser:appgroup /cer

# Use the non-root user
USER appuser
EXPOSE 8080
ENTRYPOINT ["./cer"]

