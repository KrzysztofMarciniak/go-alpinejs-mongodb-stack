#!/bin/bash
echo "Deploying to production"
docker compose up nginx api mongo
