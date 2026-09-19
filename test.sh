#!/bin/bash

BASE_URL="http://localhost:3000"

echo "=== Register user ==="

API_KEY=$(curl -s -X POST "$BASE_URL/register?name=Ali")

echo "API Key: $API_KEY"
echo


echo "=== Valid API key ==="

curl -i -s \
  -H "X-API-Key: $API_KEY" \
  "$BASE_URL/hello"

echo
echo


echo "=== Invalid API key ==="

curl -i -s \
  -H "X-API-Key: wrong-key" \
  "$BASE_URL/hello"

echo
echo


echo "=== Rate limit test ==="

for i in {1..91}
do
    echo "Request $i:"
    curl -i -s \
      -H "X-API-Key: $API_KEY" \
      "$BASE_URL/hello"
    echo
done

