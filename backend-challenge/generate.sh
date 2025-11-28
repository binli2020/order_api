#!/usr/bin/env bash

echo "=== Generating OpenAPI models and Chi server code ==="

# Exit immediately on error
set -e

# Generate models
oapi-codegen -generate types -package generated spec/openapi.yaml > internal/generated/models.gen.go

# Generate Chi server code
oapi-codegen -generate chi-server -package generated spec/openapi.yaml > internal/generated/server.gen.go

echo "Models generated: internal/generated/models.gen.go"
echo "Server generated: internal/generated/server.gen.go"
echo "=== Generation complete ==="
