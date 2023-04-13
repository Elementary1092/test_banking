PROJ_DIR=$(PWD)
BUILD_DIR=$(PROJ_DIR)/_build
MAIN_FILE=$(PROJ_DIR)/cmd/main.go

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