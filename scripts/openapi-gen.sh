#!/bin/bash
set -e # exit script at the first error

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"

mkdir -p "$output_dir"
mkdir -p "$output_dir/client/$service"
oapi-codegen -generate types -o "$output_dir/api_types.gen.go" -package "$package" "docs/api/$service.yml"
oapi-codegen -generate chi-server -o "$output_dir/api.gen.go" -package "$package" "docs/api/$service.yml"
oapi-codegen -generate types -o "$output_dir/client/$service/api_types_gen.go" -package "$service" "docs/api/$service.yml"
oapi-codegen -generate client -o "$output_dir/client/$service/openapi_client_gen.go" -package "$service" "docs/api/$service.yml"