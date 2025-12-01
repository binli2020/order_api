#!/usr/bin/env bash
set -e

BASE_URL="http://localhost:8080"
# BASE_URL="https://orderfoodonline.deno.dev/api"

echo "📦 Test: List products"
curl -s "${BASE_URL}/product" | jq .

echo "📦 Test: Get product by ID"
curl -s "${BASE_URL}/product/1" | jq .

echo "🧾 Test: Place order"
curl -s -X POST "${BASE_URL}/order" \
  -H "Content-Type: application/json" \
  -d '{
        "couponCode": "HAPPYHRS",
        "items": [
          { "productId": "1", "quantity": 2 },
          { "productId": "3", "quantity": 1 }
        ]
      }' | jq .
