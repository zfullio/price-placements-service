FROM golang:alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o price-placement-service /build/cmd/server/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /build/price-placement-service ./
RUN chmod +x price-placement-service
ENTRYPOINT ["./price-placement-service", "--env"]