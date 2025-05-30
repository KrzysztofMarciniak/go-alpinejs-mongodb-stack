# Stage 1: deps & dev
FROM node:alpine AS dev
WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm install
COPY . .
CMD ["npm", "run", "dev"]

# Stage 2: production build
FROM node:alpine AS builder
WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm ci
COPY . .
RUN npm run build

# Stage 3: serve
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html

