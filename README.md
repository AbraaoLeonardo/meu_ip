# meu_ip

A minimal web project that returns each visitor's public IP. Backend is written in Go (Golang) and frontend uses Vue.js.

## Features
- Simple REST endpoint that returns the client's public IP.
- Lightweight Vue single-page app that displays the IP.
- Handles reverse-proxy headers (X-Forwarded-For).

## Architecture
- Backend: Go HTTP server exposing `/api/hello`.
- Frontend: Vue.js app that calls the API and shows the IP.
- Deployable as standalone services or in Docker.

## Tech stack
- Backend: Go (net/http)
- Frontend: Vue 3 (Vite)
- Optional: Docker

## API
GET /api/hello
Response:
```
{ "ip": "203.0.113.45" }
```
Notes:
- Server read `RemoteAddr`.

## Backend (Go) — outline
- Return JSON `{ "Content": "<client-ip>" }`.
- Example run:
```
go run ./cmd/server
```
Environment variables:
- PORT (default 8080)

## Frontend (Vue) — outline
- Small component that fetches `/api/ip` and displays it.
- Example (Vite):
```
npm install
npm run dev
```
Fetch example:
```js
fetch('/api/ip')
    .then(r => r.json())
    .then(data => { ip.value = data.ip })
```

## Deployment
- Serve frontend assets from a static server or the Go backend.
- In Docker, build two images (backend/frontend)
- The deploy can be maded by terraform on GCP and AWS.

## Quick start (local)
1. Start backend:
     - go run backend/main.go
2. Start frontend (from /web):
     - npm install
     - npm run dev
Visit http://localhost:8081 (or configured port) to see your public IP.