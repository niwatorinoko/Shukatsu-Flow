# ---- build stage ----
FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/api ./cmd/api

# ---- runtime stage ----
FROM alpine:3.20

WORKDIR /app
COPY --from=build /bin/api /app/api

EXPOSE 8080
CMD ["/app/api"]