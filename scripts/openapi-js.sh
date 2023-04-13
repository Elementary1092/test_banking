#!/bin/bash
set -e

readonly service="$1"
readonly dir="$2"

docker run -t -i -p 8246:8080 -e SWAGGER_JSON=/$(service).yml -v $(pwd)/docs/$(service)/$(service).yml:/$(service).yml swaggerapi/swagger-ui