FROM node:22-alpine AS shortener-frontend
WORKDIR /src
COPY url-shortener/frontend/package*.json ./frontend/
RUN cd frontend && npm ci
COPY url-shortener/frontend ./frontend
RUN cd frontend && npm run build

FROM nginx:1.27-alpine
COPY portfolio/deploy/nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=shortener-frontend /src/frontend/dist /usr/share/nginx/html/shortener
