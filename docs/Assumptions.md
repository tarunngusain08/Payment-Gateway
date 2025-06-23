# Design & Implementation Assumptions

## 1. In-Memory Storage
- **Assumption:** Transaction data is stored in-memory for demo and development.
- **Reasoning:** Simplifies setup and enables fast prototyping. For production, a persistent database should be used.

## 2. Gateway Simulation
- **Assumption:** Mock gateways (A: JSON, B: XML/SOAP) are used for integration and load testing.
- **Reasoning:** Allows safe, repeatable testing without relying on real payment providers.

## 3. Worker Pool Sizing
- **Assumption:** Number of workers matches or slightly exceeds logical CPU cores.
- **Reasoning:** Balances CPU utilization and throughput; avoids excessive context switching and resource contention.

## 4. Buffer Size
- **Assumption:** Buffer size is set to 3–10x the number of workers.
- **Reasoning:** Absorbs request bursts while minimizing latency and risk of deadline exceeded.

## 5. Context Deadlines
- **Assumption:** All gateway calls and worker pool tasks use context with deadlines/timeouts.
- **Reasoning:** Prevents resource leaks and ensures timely failure in case of slow or unresponsive gateways.

## 6. Idempotency & Caching
- **Assumption:** Idempotency is enforced via caching for transaction and callback requests.
- **Reasoning:** Prevents duplicate processing and ensures safe retries.

## 7. Extensibility
- **Assumption:** New gateways can be added by implementing the `PaymentGateway` interface.
- **Reasoning:** Promotes modularity and future growth.

## 8. Observability
- **Assumption:** Structured logging and error propagation are used throughout.
- **Reasoning:** Facilitates debugging, monitoring, and production readiness.

## 9. Load Testing
- **Assumption:** Load tests simulate real-world traffic patterns, including random callback delays.
- **Reasoning:** Provides realistic performance and resilience insights.

## 10. Kubernetes/Cloud Scaling
- **Assumption:** Production deployments will use Kubernetes with horizontal scaling, redundancy, and autoscaling.
- **Reasoning:** Ensures high availability, resilience, and scalability in cloud environments. Allows for rolling updates, self-healing, and efficient resource utilization.

---
