FROM node:22-alpine AS build
WORKDIR /src
COPY portfolio/frontend ./frontend
RUN cd frontend && rm -rf node_modules && npm ci && npm run build

FROM nginx:1.27-alpine
COPY --from=build /src/frontend/dist /usr/share/nginx/html
COPY portfolio/deploy/nginx/portfolio-frontend.conf /etc/nginx/conf.d/default.conf
