# This project uses my Rust Redis clone
# It is functionally similar enough that they
# they're interchangable for simple use cases like this

FROM rust:1.89 AS builder
RUN apt-get update && apt-get install -y git && rm -rf /var/lib/apt/lists/*
WORKDIR /app

RUN git clone https://github.com/thomascpowell/redis.git .
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/target/release/redis .

EXPOSE 6379
CMD ["./redis"]
