FROM golang:1.20-alpine AS base
SHELL ["/bin/ash", "-eo", "pipefail", "-c"]

WORKDIR /build

COPY . .
RUN go mod vendor

## Linting stage
FROM base AS lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53 \
    && golangci-lint run --timeout 5m

## Test stage
FROM base AS test
RUN go test -v ./...

## Build stage
FROM base as build
RUN go build -o /dist/receipt-processor

## Final stage
FROM scratch AS final
COPY --from=build /dist/receipt-processor .
EXPOSE 8080
CMD ["./receipt-processor"]