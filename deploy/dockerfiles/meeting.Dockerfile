FROM golang:1.25-alpine AS build
WORKDIR /src
COPY meeting-room-mvp/go.mod meeting-room-mvp/go.sum ./
RUN go mod download
COPY meeting-room-mvp .
RUN CGO_ENABLED=0 go build -o /out/meeting ./cmd/server

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/meeting /app/meeting
COPY meeting-room-mvp/internal /app/internal
COPY meeting-room-mvp/web /app/web
EXPOSE 8080
ENTRYPOINT ["/app/meeting"]
