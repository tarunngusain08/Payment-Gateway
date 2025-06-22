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

3. **Stop and clean up:**
   ```sh
   docker-compose down
   ```

---

### 3. Performance Testing and Profiling

#### Benchmarking
Run benchmarks for all packages:
```sh
make bench
```

To benchmark specific packages:
```sh
go test -bench=. -benchmem ./internal/gateway/...
go test -bench=. -benchmem ./internal/service/...
go test -bench=. -benchmem ./internal/handler/...
```

#### CPU Profiling
Generate CPU profiles:
```sh
# All packages
make profile

# Specific packages
go test -cpuprofile cpu.prof ./internal/gateway/...
go test -cpuprofile cpu.prof ./internal/service/...
go test -cpuprofile cpu.prof ./internal/handler/...
```

Analyze CPU profile:
```sh
go tool pprof cpu.prof
```

Common pprof commands:
- `top`: Show top CPU consumers
- `web`: Generate graph visualization (requires graphviz)
- `list <function>`: Show source code with CPU samples

#### Memory Profiling
Generate memory profiles:
```sh
# All packages
make profile

# Specific packages
go test -memprofile mem.prof ./internal/gateway/...
go test -memprofile mem.prof ./internal/service/...
go test -memprofile mem.prof ./internal/handler/...
```

Analyze memory profile:
```sh
go tool pprof mem.prof
```

#### Trace Profiling
Generate execution traces:
```sh
# Single package trace
go test -trace trace.out ./internal/gateway/...

# View trace
go tool trace trace.out
```

#### Profile Visualization
Install graphviz for visual profile analysis:
```sh
# macOS
brew install graphviz

# Ubuntu
apt-get install graphviz
```

Then use:
```sh
go tool pprof -web cpu.prof    # CPU profile graph
go tool pprof -web mem.prof    # Memory profile graph
```
