FROM golang:1.17 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o print-headers cmd/main.go

FROM alpine:latest
EXPOSE 8080
WORKDIR /app
COPY --from=build /build/print-headers .
CMD ["./print-headers"]
