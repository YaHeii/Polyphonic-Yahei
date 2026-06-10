FROM golang:1.25.9-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/admin-api ./service/api/admin

FROM alpine:3.22

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /out/admin-api /app/admin-api
COPY service/api/admin/etc/admin-api.yaml /app/etc/admin-api.yaml

RUN mkdir -p /app/runtime/resource

CMD ["/app/admin-api", "-f", "/app/etc/admin-api.yaml"]
