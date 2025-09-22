# Drive
Fullstack file storage server with Dropbox-like functionality.

[![Tests](https://github.com/thomascpowell/drive/actions/workflows/tests.yml/badge.svg)](https://github.com/thomascpowell/drive/actions/workflows/tests.yml)

## Features
- File management operations (upload, download, delete)
- Clean UI with secure authentication via JWTs
- Files are private by default, with optional share link generation

## Stack
- Frontend: SvelteKit (TypeScript) with Vite
- Backend: Go (Gin, GORM)
- Stores: SQLite, Redis
- Deployment: Docker, Nginx

## Requirements
- For Development: `Node`, `Go`, and `Redis`
- For Production: Any Docker environment

> [!NOTE]
> This project uses my [Redis clone](https://github.com/thomascpowell/redis/), which is interchangable with the official Redis image for the purposes of this project. The backend still uses the official `go-redis` client library. Any RESP-compliant TCP server will work for development and testing.


## Usage
### Development
```sh
# Start any RESP-compatible TCP server
# (Only required for some features)
redis-server

# Start the backend
cd backend
go run main.go

# Start the frontend
cd frontend
npm install
npm run dev
```

### Testing
```sh
cd backend

# Run unit tests
go test -v ./tests

# Run Redis integration tests
# (Requires a Redis server running)
go test -v ./integration
```

### Production
```sh
docker-compose up --build
```
