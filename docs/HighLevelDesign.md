# High-Level Design: Payment-Gateway Microservice

## 1. Overview

The Payment-Gateway microservice is a modular, extensible, and resilient service for processing payment transactions (deposits and withdrawals) via multiple external payment gateways. It is designed for high throughput, reliability, and easy integration of new gateways and protocols.

---

## 2. Functional Requirements

- Accept deposit and withdrawal requests via REST API.
- Route requests to the appropriate payment gateway (JSON/SOAP/XML).
- Support asynchronous processing with worker pool and context cancellation.
- Ensure idempotency for transaction requests.
- Handle gateway callbacks to update transaction status.
- Provide transaction status query endpoint.
- Expose health and readiness endpoints for monitoring.
- Support configuration-driven gateway registration and resilience settings.

---

## 3. Non-Functional Requirements

- **Scalability:**  
  - Should handle hundreds to thousands of concurrent requests per second with horizontal scaling.
  - Worker pool and async processing allow efficient CPU utilization.
- **Reliability:**  
  - Circuit breaker, retries, and timeouts for external gateway calls.
  - Goroutine leak prevention and context-aware cancellation.
- **Observability:**  
  - Structured logging, error propagation, and metrics.
- **Extensibility:**  
  - Easy to add new gateways and protocols.
- **Security:**  
  - Authentication middleware and input validation.
- **Testability:**  
  - High unit/integration test coverage with mocks and simulated gateways.

---

## 4. Rough Scale Estimation

- **Throughput:**  
  - With 10 workers (default), can process ~100-500 transactions/sec on a single node (depends on gateway latency).
  - Horizontal scaling: Add more nodes and increase worker pool size.
- **Latency:**  
  - End-to-end latency is typically gateway-limited (e.g., 100ms-2s per transaction).
  - Internal processing adds minimal overhead due to async design.
- **Memory/CPU:**  
  - In-memory store is lightweight; memory usage grows with transaction volume but is bounded by TTL and cleanup.
  - CPU usage scales with number of workers and request volume.

---

## 5. API Definitions & Routes

### Deposit

- **POST /deposit**
  - **Request:**  
    ```json
    {
      "account": "string",
      "amount": 100.0
    }
    ```
  - **Response:**  
    ```json
    {
      "transactionId": "string",
      "status": "pending|success|failed"
    }
    ```

### Withdrawal

- **POST /withdrawal**
  - **Request:**  
    ```json
    {
      "account": "string",
      "amount": 100.0
    }
    ```
  - **Response:**  
    ```json
    {
      "transactionId": "string",
      "status": "pending|success|failed"
    }
    ```

### Gateway Callbacks

- **POST /callback/gateway-a**  
  - Handles JSON callback from GatewayA.
- **POST /callback/gateway-b**  
  - Handles XML/SOAP callback from GatewayB.

### Mock Gateway Endpoints (for local testing)

- **POST /mock-gateway-a/deposit**
- **POST /mock-gateway-a/withdrawal**
- **POST /mock-gateway-b/deposit**
- **POST /mock-gateway-b/withdrawal**

---

## 6. Route Summary

| Method | Path                        | Description                        |
|--------|-----------------------------|------------------------------------|
| POST   | /deposit                    | Initiate deposit                   |
| POST   | /withdrawal                 | Initiate withdrawal                |
| POST   | /callback/gateway-a         | GatewayA callback (JSON)           |
| POST   | /callback/gateway-b         | GatewayB callback (XML/SOAP)       |
| POST   | /mock-gateway-a/deposit     | Mock GatewayA deposit endpoint     |
| POST   | /mock-gateway-a/withdrawal  | Mock GatewayA withdrawal endpoint  |
| POST   | /mock-gateway-b/deposit     | Mock GatewayB deposit endpoint     |
| POST   | /mock-gateway-b/withdrawal  | Mock GatewayB withdrawal endpoint  |
