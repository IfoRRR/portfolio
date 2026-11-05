FROM golang:1.25-alpine AS build
WORKDIR /src
COPY vk-quiz/go.mod vk-quiz/go.sum ./
RUN go mod download
COPY vk-quiz .
RUN CGO_ENABLED=0 go build -o /out/quiz ./cmd/quiz

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/quiz /app/quiz
COPY vk-quiz/static /app/static
RUN mkdir -p /data
EXPOSE 8093
ENTRYPOINT ["/app/quiz"]
