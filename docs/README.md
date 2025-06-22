# Payment-Gateway

## Project Overview

This project implements a modular, extensible payment gateway service in Go. It supports multiple gateway integrations (e.g., GatewayA with JSON, GatewayB with SOAP/XML), transaction management, and callback handling. The codebase is designed for testability, maintainability, and ease of extension.

---

## Project Structure

- **cmd/**: Entry point for the application (main.go, wiring).
- **internal/**
  - **constants/**: Enumerations and constant values (transaction types, statuses).
  - **dtos/**: Data Transfer Objects for requests/responses, including XML/JSON struct tags for gateway compatibility.
  - **gateway/**: Gateway implementations (e.g., GatewayA for JSON, GatewayB for SOAP/XML), with interface abstraction.
  - **handler/**: HTTP handlers for transaction and callback endpoints, including gateway-specific callback handlers.
  - **middleware/**: HTTP middleware (auth, logging, etc.).
  - **models/**: Core business models (Transaction, DepositRequest, WithdrawalRequest).
  - **repository/**: In-memory repository for transactions, with thread-safe operations.
  - **service/**: Business logic, including transaction processing, gateway pool, and callback services.
- **pkg/**
  - **error/**: Custom error types for domain-specific error handling.
  - **mocks/**: Auto-generated mocks for interfaces, used in unit tests.
- **docs/**: Documentation, including this README.
- **Makefile**: Automation for build, test, lint, and formatting.
- **docker-compose.yml**: Orchestrates the application and dependencies for local development/testing.

---

## Implementation Highlights

- **Gateway Abstraction:**  
  The `PaymentGateway` interface allows seamless integration of new gateways. Each gateway (A/B) implements its own protocol (JSON, SOAP/XML) and error handling.

- **Callback Handling:**  
  Separate handlers for each gateway (`GatewayACallbackHandler`, `GatewayBCallbackHandler`) parse incoming callback data in the expected format (JSON or XML), validate, and update transaction status.

- **Transaction Management:**  
  The `TransactionService` coordinates transaction creation, processing via gateways, and status updates. The repository uses a thread-safe in-memory store for demo purposes.

- **Testing:**  
  - Extensive unit tests for all services, handlers, and repositories.
  - Use of GoMock for mocking interfaces.
  - Gateway tests use `httptest.Server` to simulate real gateway responses, including timeouts and error scenarios.
  - DTOs have both JSON and XML struct tags to ensure correct marshaling/unmarshaling for each gateway.

- **Error Handling:**  
  Custom error types in `pkg/error` provide clear, domain-specific error messages and codes.

- **Extensibility:**  
  Adding a new gateway requires implementing the `PaymentGateway` interface and registering it in the gateway pool.

---

## Notable Implementation Details

- **XML/JSON Compatibility:**  
  DTOs use struct tags for both XML and JSON, ensuring correct parsing for each gateway protocol.
- **Thread Safety:**  
  The in-memory repository uses `sync.Map` for safe concurrent access.
- **Mocking and Test Coverage:**  
  All interfaces are mocked for unit testing, and all major code paths are covered by tests.
- **Error Propagation:**  
  Errors are propagated and handled at the appropriate layer, with HTTP handlers returning meaningful status codes.

---

## Troubleshooting

- If containers fail to start, check logs with `docker-compose logs`.
- If tests fail, run them with `make test` and review the output.
- Ensure Docker is running and you have sufficient resources allocated.

---

## Extending the Service

To add a new payment gateway:
1. Implement the `PaymentGateway` interface in `internal/gateway/`.
2. Add the new gateway to the gateway pool in the service layer.
3. Add new handlers or callback logic as needed.
4. Write tests for the new gateway and its integration.

---

## Summary

This project demonstrates a robust, extensible, and testable payment gateway service in Go, with clear separation of concerns, strong test coverage, and modern development tooling.
