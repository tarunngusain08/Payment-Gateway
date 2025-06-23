APP_NAME=payment-gateway
PKG=./...
MOCKS_DIR=pkg/mocks

.PHONY: all build run test bench profile mocks clean load_test

all: build

build:
	go build -o $(APP_NAME) ./cmd

run: build
	./$(APP_NAME)

test:
	go test -v $(PKG)

bench:
	go test -bench=. -benchmem $(PKG)

profile:
	go test -cpuprofile cpu.prof -memprofile mem.prof $(PKG)

mocks:
	mockgen -source=internal/gateway/interface.go -destination=$(MOCKS_DIR)/mock_gateway.go -package=mocks
	mockgen -source=internal/service/interface.go -destination=$(MOCKS_DIR)/mock_service.go -package=mocks
	mockgen -source=internal/repository/transaction_repository.go -destination=$(MOCKS_DIR)/mock_transaction_repository.go -package=mocks

clean:
	rm -f $(APP_NAME) cpu.prof mem.prof *.test *.out
	go clean

load_test:
	go test -v -run ^TestLoadSimulator$$ ./load_test.go
