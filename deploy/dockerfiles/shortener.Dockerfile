FROM node:22-alpine AS frontend
WORKDIR /src
COPY url-shortener/frontend ./frontend
RUN cd frontend && rm -rf node_modules && npm ci && npm run build

FROM golang:1.25-alpine AS build
WORKDIR /src
COPY url-shortener/go.mod url-shortener/go.sum ./
RUN go mod download
COPY url-shortener .
RUN CGO_ENABLED=0 go build -o /out/shortener ./cmd/URL-shortener

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/shortener /app/shortener
COPY url-shortener/go.mod /app/go.mod
COPY portfolio/deploy/scripts/shortener-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
RUN mkdir -p /data
EXPOSE 8282
ENTRYPOINT ["/entrypoint.sh"]
