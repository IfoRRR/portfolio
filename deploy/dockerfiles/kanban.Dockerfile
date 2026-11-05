FROM golang:1.25-alpine AS build
WORKDIR /src
COPY vk-kanban/go.mod vk-kanban/go.sum ./
RUN go mod download
COPY vk-kanban .
RUN CGO_ENABLED=0 go build -o /out/kanban ./cmd/cases

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/kanban /app/kanban
COPY vk-kanban/static /app/static
RUN mkdir -p /data
EXPOSE 8094
ENTRYPOINT ["/app/kanban"]
