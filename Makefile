# Commands for local development

.PHONY: pg_start pg_stop redis_start redis_stop be fe

pg_start:
	# uses Docker for convenience, but binds to loopback
	docker rm -f pg-dev 2>/dev/null || true
	docker run --name pg-dev -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=dev -e POSTGRES_DB=dev -p 127.0.0.1:5432:5432 -d postgres:latest

pg_stop:
	docker stop pg-dev && docker rm pg-dev || true

redis_start:
	docker rm -f redis-dev 2>/dev/null || true
	docker run --name redis-dev -p 127.0.0.1:6379:6379 -d redis:latest

redis_stop:
	docker stop redis-dev && docker rm redis-dev || true

be:
	cd backend && exec go run ./

fe:
	cd frontend && exec npm run dev

env:
	# quickly create a .env
	cp env.example .env
