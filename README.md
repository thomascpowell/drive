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

## Usage
### Development
```sh
# Start a Redis server
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
# Run unit tests
go test -v ./tests

# Run Redis integration tests
go test -v ./integration
```

### Production
```sh
docker-compose up --build
```
