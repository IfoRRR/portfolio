FROM node:22-alpine AS frontend
WORKDIR /src
COPY portfolio/frontend/package*.json ./frontend/
RUN cd frontend && npm ci
COPY portfolio/frontend ./frontend
RUN cd frontend && npm run build

FROM golang:1.25-alpine AS build
WORKDIR /src
COPY portfolio/backend/go.mod portfolio/backend/go.sum ./backend/
RUN cd backend && go mod download
COPY portfolio/backend ./backend
RUN cd backend && CGO_ENABLED=0 go build -o /out/portfolio-api .

FROM alpine:3.22
WORKDIR /app
COPY --from=build /out/portfolio-api /app/portfolio-api
COPY --from=build /src/backend/config.json /app/config/config.json
COPY --from=build /src/backend/assets /app/assets
COPY --from=build /src/backend/assets /app/default-resume
COPY --from=frontend /src/frontend/dist /app/frontend
COPY portfolio/deploy/scripts/portfolio-entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
EXPOSE 8080
ENTRYPOINT ["/entrypoint.sh"]
