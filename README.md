# Payment-Gateway

## Project Overview

Payment-Gateway is a modular, extensible payment gateway service written in Go. It supports multiple gateway integrations (e.g., GatewayA with JSON, GatewayB with SOAP/XML), transaction management, and callback handling. The codebase is designed for resilience, observability, testability, and ease of extension.

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

- **Resilience Patterns:**  
  - **Exponential Backoff Retries:** All gateway calls use exponential backoff with configurable retry limits to handle transient failures.
  - **Circuit Breaking:** Circuit breaker pattern is implemented for each gateway to prevent cascading failures and allow recovery.
  - **Timeouts:** All gateway and worker pool operations are context-aware and use timeouts to avoid indefinite blocking.
  - **Worker Pool:** A worker pool is used to process deposit and withdrawal requests asynchronously, supporting high throughput and context cancellation.
  - **Goroutine Leak Prevention:** All asynchronous operations are context-aware, ensuring goroutines are not leaked and are properly cleaned up on cancellation.
  - **Idempotency Caching:** Caching is used to ensure idempotency for transaction requests, preventing duplicate processing.

- **Testing:**  
  - Extensive unit tests for all services, handlers, and repositories.
  - Use of GoMock for mocking interfaces.
  - Gateway tests use `httptest.Server` to simulate real gateway responses, including timeouts and error scenarios.
  - DTOs have both JSON and XML struct tags to ensure correct marshaling/unmarshaling for each gateway.

- **Error Handling:**  
  Custom error types in `pkg/error` provide clear, domain-specific error messages and codes.

- **Extensibility:**  
  Adding a new gateway requires implementing the `PaymentGateway` interface and registering it in the gateway pool.

- **XML/JSON Compatibility:**  
  DTOs use struct tags for both XML and JSON, ensuring correct parsing for each gateway protocol.

- **Thread Safety:**  
  The in-memory repository uses `sync.Map` for safe concurrent access.

- **Mocking and Test Coverage:**  
  All interfaces are mocked for unit testing, and all major code paths are covered by tests.

- **Error Propagation:**  
  Errors are propagated and handled at the appropriate layer, with HTTP handlers returning meaningful status codes.

---

## Setup & Running Instructions

### Prerequisites

- Go 1.20+
- Docker & Docker Compose (for containerized development)
- Make (for build/test automation)

### Local Development

1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-org/Payment-Gateway.git
   cd Payment-Gateway
   ```

2. **Build the project:**
   ```sh
   make build
   ```

3. **Run the service locally:**
   ```sh
   make run
   ```
   Or, using Docker Compose:
   ```sh
   docker-compose up --build
   ```

4. **Configuration:**
   - Edit `internal/config/config.yaml` to adjust gateway URLs, ports, and resilience settings.

---

## How to Run Tests

- **Unit tests:**
  ```sh
  make test
  ```

- **Load tests:**
  ```sh
  make load_test
  ```
  See [Benchmarking Results](docs/Benchmarking-Results.md) for detailed performance analysis.

---

## Sample API Payloads

See [`docs/DummyCurls&Responses.md`](docs/DummyCurls&Responses.md) for ready-to-use curl commands and example responses.

---

## Tradeoffs and Assumptions

- **In-memory storage:** The repository uses an in-memory store for demo purposes. For production, swap with a persistent backend.
- **Gateway simulation:** Mock endpoints are provided for GatewayA and GatewayB; real integrations would require secure credentials and error handling.
- **Resilience:** Timeout, retry, and circuit breaker patterns are implemented for gateway calls. These are configurable via YAML.
- **Extensibility:** Adding a new gateway requires implementing the `PaymentGateway` interface and registering it in the pool.
- **Observability:** All handlers, services, and gateways use structured logging with trace and request IDs.
- **Testing:** The project is designed for high test coverage and easy mocking.

---

## Documentation
- [High Level Design](docs/High-Level-Design.md)
- [Build & Development Guide](docs/Build-Steps.md)
- [Sample API Requests & Responses](docs/Dummy-Curls&Responses.md)
- [Benchmarking Results](docs/Benchmarking-Results.md)
- [Assumptions](docs/Assumptions.md)
- [GCP Deployment Guide](docs/GCP-Deployment.md)
- [Minikube Deployment Guide](docs/Minikube-Deployment.md)

---
