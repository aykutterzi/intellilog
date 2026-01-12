#!/bin/bash

echo "1. Logging in..."
TOKEN=$(curl -s -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "password"}' | jq -r .token)

if [ "$TOKEN" == "null" ]; then
  echo "Login failed"
  exit 1
fi
echo "Token received."

echo "2. Check components health..."
curl -s http://localhost:8080/health | jq .

echo "3. Posting Normal Log..."
curl -s -X POST http://localhost:8080/api/logs \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "level": "INFO",
    "service": "user-service",
    "message": "User logged in successfully"
  }' | jq .

echo "4. Posting Anomaly Log (High Latency)..."
curl -s -X POST http://localhost:8080/api/logs \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "level": "WARNING",
    "service": "order-service",
    "message": "Database query latency 2500ms detected"
  }' | jq .

echo "5. Posting Anomaly Log (OOM)..."
curl -s -X POST http://localhost:8080/api/logs \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "level": "ERROR",
    "service": "payment-service",
    "message": "Out of memory exception in worker pool"
  }' | jq .

echo "6. Retrieving Logs..."
curl -s -X GET http://localhost:8080/api/logs \
  -H "Authorization: Bearer $TOKEN" | jq .
