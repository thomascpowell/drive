# For development and testing
# See docker-compose.yml for prod

.PHONY: pg_start pg_stop redis_start redis_stop be fe


# Local development servers
backend_dev:
	cd backend && exec go run ./

frontend_dev:
	cd frontend && exec npm run dev


# Docker containers (for development)
pg_start:
	docker rm -f pg-dev 2>/dev/null || true
	docker run --name pg-dev -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=dev -e POSTGRES_DB=dev -p 127.0.0.1:5432:5432 -d postgres:latest

pg_stop:
	docker stop pg-dev && docker rm pg-dev || true

redis_start:
	docker rm -f redis-dev 2>/dev/null || true
	docker run --name redis-dev -p 127.0.0.1:6379:6379 -d redis:latest

redis_stop:
	docker stop redis-dev && docker rm redis-dev || true


# Default environment
env:
	cp env.example .env


# Testing
unit: 
	cd backend && exec go test --count=1 -v ./unit

integration: pg_stop redis_stop pg_start redis_start
	cd backend && exec go test --count=1 -v ./integration
