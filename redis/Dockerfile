# This project uses my Rust Redis clone
# They are interchangable in this project
# The backend uses the offical redis-go client

FROM rust:1.89 AS builder
WORKDIR /app
RUN git clone https://github.com/thomascpowell/redis.git .
RUN cargo build --release

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/target/release/redis .

EXPOSE 6379
CMD ["./redis", "0.0.0.0:6379"]
