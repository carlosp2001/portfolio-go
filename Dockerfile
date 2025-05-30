# --------- Build Stage ---------
FROM golang:1.24-alpine AS builder
WORKDIR /app

RUN apk add --no-cache nodejs npm make
RUN npm install -g sass

# Cache dependencies
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

RUN make sass

# Build with static linking (no C dependencies)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./backend

# --------- Runtime Stage ---------
FROM alpine:latest
WORKDIR /app

# Copy binary
COPY --from=builder /app/app .
COPY --from=builder /app/frontend ./frontend
RUN rm -rf /frontend/scss

# Expose port if needed
EXPOSE 8080

## Healthcheck (optional)
#HEALTHCHECK CMD wget --no-verbose --tries=1 --spider http://localhost:8080/health || exit 1

# Run the app
ENTRYPOINT ["./app"]