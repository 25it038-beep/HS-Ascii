# HS ASCII LIVE
### By HS AI Solutions

A production-ready, publicly accessible ASCII animation streaming platform with 95+ continuously animated ASCII experiences.

## Features
- Real HTTP streaming with ANSI sequences
- 95+ animations across 7 categories
- Concurrent viewers, context cancellation
- REST API for metadata
- Premium React frontend with Obsidian Black + Champagne Gold theme
- Docker deployment for Render
- Extensible animation engine

## Architecture
backend/
- engine: frame, registry, renderer, engine
- handlers: animation, api, health
- animations: procedural generators per category
- middleware: cors, ratelimit, logging
- config

## Local Installation
```bash
cd backend
go mod tidy
go run main.go
```

## Endpoints
- `GET /<animation>` - stream animation
- `GET /random` - random animation
- `GET /list` - list animations
- `GET /help` - usage
- `GET /health` - health
- `GET /api/animations` - JSON catalogue
- `GET /api/stats` - stats

Examples:
```bash
curl https://hs-ascii-live.onrender.com/rick
curl https://hs-ascii-live.onrender.com/matrix?speed=fast&color=green&fps=20
```

## Animation List
Characters, Space, Nature, Technology, Abstract, Symbols, HS AI Exclusive - 95 total.

## Deployment
Backend: Render via `render.yaml` + Dockerfile
Frontend: Vercel

## Testing
```bash
go test ./...
```

## License
HS AI Solutions
