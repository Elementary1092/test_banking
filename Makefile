PROJ_DIR=$(PWD)
BUILD_DIR=$(PROJ_DIR)/build
MAIN_FILE=cmd/main.go

include .env

.PHONY: mocks_gen
mocks_gen:
	mkdir -p "$(PROJ_DIR)/internal/domain/customer/command/mocks/"
	mockgen -source=$(PROJ_DIR)/internal/domain/customer/command/dao.go -destination=$(PROJ_DIR)/internal/domain/customer/command/mocks/dao_mock.go -package=mocks
	mkdir -p "$(PROJ_DIR)/internal/domain/customer/query/mocks/"
	mockgen -source=$(PROJ_DIR)/internal/domain/customer/query/dao.go -destination=$(PROJ_DIR)/internal/domain/customer/query/mocks/dao_mock.go -package=mocks
	mkdir -p "$(PROJ_DIR)/internal/domain/account/query/mocks/"
	mockgen -source=$(PROJ_DIR)/internal/domain/account/query/dao.go -destination=$(PROJ_DIR)/internal/domain/account/query/mocks/dao_mock.go -package=mocks
	mkdir -p "$(PROJ_DIR)/internal/domain/account/command/mocks/"
	mockgen -source=$(PROJ_DIR)/internal/domain/account/command/dao.go -destination=$(PROJ_DIR)/internal/domain/account/command/mocks/dao_mock.go -package=mocks

.PHONY: api_gen
api_gen:
	@./scripts/openapi-gen.sh api internal/adapters/http api

.PHONY: api_swagger
api_swagger:
	docker run -d -t -i -p 8246:8080 -e SWAGGER_JSON=/api.yml -v $(PROJ_DIR)/docs/api/api.yml:/api.yml swaggerapi/swagger-ui

.PHONY: run
run:
	go run $(MAIN_FILE)

.PHONY: build
build:
	go build -o $(BUILD_DIR)/app $(MAIN_FILE)