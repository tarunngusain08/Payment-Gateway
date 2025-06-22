# Payment-Gateway

## Development Workflow

### 1. Using the Makefile

The Makefile provides shortcuts for common development tasks. Typical targets include:

- **Build the project:**
  ```sh
  make build
  ```
  Compiles the Go binaries and prepares the application for running.

- **Run tests:**
  ```sh
  make test
  ```
  Runs all Go tests in the repository, showing verbose output and failing if any test fails.

- **Lint the code:**
  ```sh
  make lint
  ```
  Runs linters (such as `golangci-lint`) to check code quality and style.

- **Format the code:**
  ```sh
  make fmt
  ```
  Formats all Go files using `gofmt` or similar tools.

- **Clean build artifacts:**
  ```sh
  make clean
  ```
  Removes binaries and temporary files generated during build.

#### Analysis

- The Makefile ensures a consistent workflow for building, testing, and maintaining code quality.
- It automates repetitive tasks, reducing manual errors and saving time.
- Running `make test` before pushing code helps catch regressions early.

---

### 2. Using Docker Compose

Docker Compose is used to orchestrate the Payment-Gateway service and its dependencies (such as databases or mock gateways) in containers.

#### Common Commands

- **Start all services:**
  ```sh
  docker-compose up
  ```
  Builds (if needed) and starts all services defined in `docker-compose.yml`.

- **Start in detached mode:**
  ```sh
  docker-compose up -d
  ```
  Runs containers in the background.

- **Stop all services:**
  ```sh
  docker-compose down
  ```
  Stops and removes all containers, networks, and volumes created by `up`.

- **Rebuild images:**
  ```sh
  docker-compose build
  ```

- **View logs:**
  ```sh
  docker-compose logs -f
  ```

- **Run a one-off command in a service:**
  ```sh
  docker-compose run <service> <command>
  ```

#### Example Workflow

1. **Build and test locally:**
   ```sh
   make build
   make test
   ```

2. **Start services for integration testing:**
   ```sh
   docker-compose up
   ```

3. **Run integration tests (if any):**
   ```sh
   make test-integration
   ```

4. **Stop and clean up:**
   ```sh
   docker-compose down
   ```

#### In-Depth Analysis

- **Isolation:** Each service runs in its own container, ensuring a clean and reproducible environment.
- **Dependency Management:** Easily spin up databases, caches, or other services required by the Payment-Gateway.
- **Portability:** The same Compose file can be used across different machines and CI environments.
- **Scaling:** You can scale services (e.g., multiple gateway containers) for load testing:
  ```sh
  docker-compose up --scale gateway=3
  ```
- **Logs and Debugging:** Centralized logs make it easy to debug issues across services.

---

### 3. Typical Development Cycle

1. Edit code and write tests.
2. Run `make test` to ensure all tests pass.
3. Use `make build` to compile the application.
4. Use Docker Compose to run the full stack locally.
5. Iterate as needed, using `docker-compose logs` and `make` targets for feedback.

---

### 4. Troubleshooting

- If containers fail to start, check logs with `docker-compose logs`.
- If tests fail, run them with `make test` and review the output.
- Ensure Docker is running and you have sufficient resources allocated.

---

### 5. Additional Tips

- Update dependencies as needed using Go modules.
- Use `.env` files or Compose environment variables for configuration.
- For production, consider using Compose overrides or Kubernetes for orchestration.

---