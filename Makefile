PROJ_DIR=$(PWD)
BUILD_DIR=$(PROJ_DIR)/_build
MAIN_FILE=$(PROJ_DIR)/cmd/main.go

mocks_gen:
	mkdir -p "$(PROJ_DIR)/internal/domain/customer/command/mocks/"
	touch $(PROJ_DIR)/internal/domain/customer/command/mocks/mocks.go
	mockgen -source=$(PROJ_DIR)/internal/domain/customer/command/dao.go -destination=$(PROJ_DIR)/internal/domain/customer/command/mocks/mocks.go -package=mocks
	mkdir -p "$(PROJ_DIR)/internal/domain/customer/query/mocks/"
	touch $(PROJ_DIR)/internal/domain/customer/query/mocks/mocks.go
	mockgen -source=$(PROJ_DIR)/internal/domain/customer/query/dao.go -destination=$(PROJ_DIR)/internal/domain/customer/query/mocks/mocks.go -package=mocks