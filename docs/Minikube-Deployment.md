# Deploying Payment Gateway on Minikube

## Prerequisites
- Minikube installed
- kubectl configured
- Helm installed
- Docker installed

## Deployment Steps

### 1. Start Minikube
```bash
# Start Minikube
minikube start

# Enable ingress addon (optional)
minikube addons enable ingress
```

### 2. Build and Load Docker Image
```bash
# Point shell to minikube's docker daemon
eval $(minikube docker-env)

# Build the Docker image
docker build -t payment-gateway:latest .
```

### 3. Configure Helm Values
Update `helm/payment-gateway/values.yaml`:
```yaml
image:
  repository: payment-gateway
  tag: latest
  pullPolicy: Never  # Important: Use Never for local images
```

### 4. Deploy using Helm
```bash
# Create namespace (optional)
kubectl create namespace payment-gateway

# Install the Helm chart
helm install payment-gateway ./helm/payment-gateway -n payment-gateway
```

### 5. Verify Deployment
```bash
# Check deployment status
kubectl get deployments -n payment-gateway

# Check pods
kubectl get pods -n payment-gateway

# Check services
kubectl get svc -n payment-gateway
```

### 6. Access the Application

Method 1: Port forwarding
```bash
kubectl port-forward svc/payment-gateway 8000:8000 -n payment-gateway
```

Method 2: Minikube service URL
```bash
minikube service payment-gateway -n payment-gateway --url
```

## Development Workflow

### Hot Reload for Development
1. Make changes to your code
2. Rebuild Docker image:
```bash
eval $(minikube docker-env)
docker build -t payment-gateway:latest .
```
3. Restart deployment:
```bash
kubectl rollout restart deployment payment-gateway -n payment-gateway
```

## Debugging

```bash
# View logs
kubectl logs -f deployment/payment-gateway -n payment-gateway

# Shell into pod
kubectl exec -it deployment/payment-gateway -n payment-gateway -- /bin/sh

# View minikube dashboard
minikube dashboard
```

## Common Issues and Solutions

1. Image Pull Issues
```bash
# Verify image is in minikube's docker
minikube ssh
docker images | grep payment-gateway
```

2. Service Access Issues
```bash
# Check service status
minikube service list
```

## Cleanup

```bash
# Remove helm release
helm uninstall payment-gateway -n payment-gateway

# Stop minikube
minikube stop

# Delete minikube cluster
minikube delete
```
