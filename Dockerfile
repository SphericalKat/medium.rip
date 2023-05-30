# frontend
FROM node:20-slim as frontend
RUN corepack enable
RUN corepack prepare pnpm@latest --activate
WORKDIR /app
COPY frontend/package.json frontend/pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile
COPY frontend/ ./
RUN pnpm build

# backend
FROM golang:1.20-alpine as backend
WORKDIR /app
COPY . .
RUN rm -rf frontend
COPY --from=frontend /app/ ./frontend
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o medium

# final
FROM debian:bullseye-slim
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=backend /app/medium /usr/local/bin/medium
CMD [ "medium" ]