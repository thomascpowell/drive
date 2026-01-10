# Drive
File storage server with Dropbox-like functionality.

[![Tests](https://github.com/thomascpowell/drive/actions/workflows/tests.yml/badge.svg)](https://github.com/thomascpowell/drive/actions/workflows/tests.yml)

## Features
- File management operations (upload, download, delete)
- Clean UI with secure authentication via JWTs
- Files are private by default, with optional share link generation

## Stack
- Frontend: SvelteKit (TypeScript) with Vite
- Backend: Go (Gin, GORM)
- Stores: Postgres, Redis
- Deployment: Docker, Nginx

## Usage
- Requires: `Node`, `Go` and any Docker enviroment
- See `Makefile` for local development and running tests
- See `docker-compose.yml` for production
