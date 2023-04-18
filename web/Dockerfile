# Stage 1 - Install dependencies and build the app in a build environment
FROM rust:slim AS build-env

# Install Trunk
RUN cargo install trunk --locked

# Install Rust WASM target
RUN rustup target add wasm32-unknown-unknown

# Install curl and protobuf compiler
RUN apt-get update && \
    apt-get install -y curl protobuf-compiler

# Install Node.js and npm
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt-get install -y nodejs

# Create a new directory for your app
WORKDIR /app

RUN npm install -D tailwindcss

# Add node_modules to path
ENV PATH="/app/node_modules/.bin:${PATH}"

# Copy your project's files into the container
COPY . .

COPY --from=ghcr.io/shifty11/cosmos-notifier-api:latest /. /app/api/

# TODO: Make this configurable
ENV GRPC_ENDPOINT_URL https://rust.decrypto.online
ENV DISCORD_OAUTH2_URL https://discord.com/api/oauth2/authorize?client_id=955835724714872942&redirect_uri=https%3A%2F%2Frust.decrypto.online%3A8080&response_type=code&scope=identify

# Build your Sycamore project using Trunk
RUN trunk build --release

# Stage 2 - Create the run-time image
FROM caddy:2.5.2-alpine
COPY --from=build-env /app/dist /var/www/cosmos-notifier-sycamore
COPY Caddyfile /etc/caddy/Caddyfile
