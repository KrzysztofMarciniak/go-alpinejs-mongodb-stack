#!/bin/bash
set -e

echo "[*] Killing all running containers..."
docker kill $(docker ps -q) 2>/dev/null || true

echo "[*] Removing all containers..."
docker rm $(docker ps -a -q) 2>/dev/null || true

echo "[*] Removing all images..."
docker rmi -f $(docker images -q) 2>/dev/null || true

echo "[*] Removing all volumes..."
docker volume rm $(docker volume ls -q) 2>/dev/null || true

echo "[*] Removing all networks (except default)..."
docker network rm $(docker network ls | grep -v "bridge\|host\|none" | awk '{print $1}') 2>/dev/null || true

echo "[*] System prune..."
docker system prune -af --volumes

echo "[+] Docker environment is clean."
