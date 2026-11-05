FROM golang:1.25-alpine AS build
WORKDIR /src
COPY websocket/go.mod websocket/go.sum ./
RUN go mod download
COPY websocket .
RUN CGO_ENABLED=0 go build -o /out/websocket .

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/websocket /app/websocket
COPY websocket/static /app/static
COPY portfolio/deploy/scripts/websocket-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
RUN mkdir -p /data
EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]
