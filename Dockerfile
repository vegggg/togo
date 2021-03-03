# pull official base image
FROM golang as builder
ENV GO111MODULE=on
# Working directory
WORKDIR /app
# Copy files
COPY go.mod .
COPY go.sum .
# Install app dependencies
RUN go mod download
# Add src app
COPY . .
# Build app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./app

# final stage
FROM gcr.io/distroless/base
# Copy binary from builder
COPY --from=builder /app/app /app
# Run server command
ENTRYPOINT ["/app"]
# Export grpc, metric, http ports
EXPOSE 5050