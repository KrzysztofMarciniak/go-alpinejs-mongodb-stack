#!/bin/bash
echo "Deploying to production"
docker compose -f docker-compose.prod.yml up --build -d
