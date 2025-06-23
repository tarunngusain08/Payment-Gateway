# Deploying Payment Gateway to Google Kubernetes Engine (GKE)

## Prerequisites
- Google Cloud SDK installed
- kubectl configured
- Helm installed
- Docker installed
- Access to a GKE cluster

## Deployment Steps

### 1. Build and Push Docker Image
```bash
# Build the Docker image
docker build -t gcr.io/[PROJECT_ID]/payment-gateway:latest .

# Push to Google Container Registry
docker push gcr.io/[PROJECT_ID]/payment-gateway:latest
```

### 2. Configure Helm Values
Update `helm/payment-gateway/values.yaml`:
```yaml
image:
  repository: gcr.io/[PROJECT_ID]/payment-gateway
  tag: latest
  pullPolicy: Always
```

### 3. Deploy using Helm
```bash
# Create namespace (optional)
kubectl create namespace payment-gateway

# Install the Helm chart
helm install payment-gateway ./helm/payment-gateway -n payment-gateway
```

### 4. Verify Deployment
```bash
# Check deployment status
kubectl get deployments -n payment-gateway

# Check pods
kubectl get pods -n payment-gateway

# Check services
kubectl get svc -n payment-gateway
```

### 5. Access the Application
By default, the service is deployed as ClusterIP. To expose it:

```bash
# Create an Ingress or change service type to LoadBalancer in values.yaml:
service:
  type: LoadBalancer
  port: 8000
```

## Configuration

The application configuration is managed through a ConfigMap created by the Helm chart. To modify the configuration:

1. Update the `config.yaml` file in `internal/config/`
2. Upgrade the Helm release:
```bash
helm upgrade payment-gateway ./helm/payment-gateway -n payment-gateway
```

## Scaling

To scale the deployment:
```bash
# Update replicas in values.yaml
helm upgrade payment-gateway ./helm/payment-gateway -n payment-gateway --set replicaCount=3
```

## Monitoring and Logs

```bash
# View logs
kubectl logs -f deployment/payment-gateway -n payment-gateway

# View pod metrics
kubectl top pods -n payment-gateway
```

## Cleanup

To remove the deployment:
```bash
helm uninstall payment-gateway -n payment-gateway
```
