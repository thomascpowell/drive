# Drive
Fullstack file storage server with Dropbox-like functionality.

[![Tests](https://github.com/thomascpowell/drive/actions/workflows/tests.yml/badge.svg)](https://github.com/thomascpowell/drive/actions/workflows/tests.yml)

## Features
- File management operations (upload, download, delete)
- Clean UI with secure authentication (JWT)
- Planned: Link-based file sharing

## Stack
- Frontend: SvelteKit (TypeScript) with Vite
- Backend: Go (Gin, GORM, SQLite)
- Deployment: Docker, Nginx

## Usage
```sh
# Backend
cd backend
go run *.go

# Frontend
cd frontend
npm install
npm run dev

# Docker
docker-compose up --build
```
