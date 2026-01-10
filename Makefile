# For development and testing
# See docker-compose.yml for prod

.PHONY: pg_start pg_stop redis_start redis_stop be fe

# Backend development server
backend_dev:
	cd backend && exec go run ./

# Frontend development server
frontend_dev:
	cd frontend && exec npm run dev

# Start a new Postgres container (for development)
pg_start:
	docker rm -f pg-dev 2>/dev/null || true
	docker run --name pg-dev -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=dev -e POSTGRES_DB=dev -p 127.0.0.1:5432:5432 -d postgres:latest

# Stop and remove the Postgres container
pg_stop:
	docker stop pg-dev && docker rm pg-dev || true

# Start a new Redis container (for development)
redis_start:
	docker rm -f redis-dev 2>/dev/null || true
	docker run --name redis-dev -p 127.0.0.1:6379:6379 -d redis:latest

# Stop and remove the Redis container
redis_stop:
	docker stop redis-dev && docker rm redis-dev || true

# (Re)start all development containers
restart_all:
	pg_stop redis_stop pg_start redis_start

# Copy the default environment
env:
	cp env.example .env

# Run unit tests
unit: 
	cd backend && exec go test --count=1 -v ./unit

# Run integration tests (excluded from CI)
integration: restart_all
	cd backend && exec go test --count=1 -v ./integration
