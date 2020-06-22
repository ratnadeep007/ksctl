# Cross Cloud Kubernetes Control

A cross cloud kubernetes deploying command line.

### Development requirement

- Golang 1.14
- upx (for compressing executable)

### Linux
```bash
sudo apt install upx -y
```
using snap
```bash
sudo snap install upx
```

#### macOS
```zsh
brew install upx
```

### Building
```bash
./build.sh
```

### Deploying in DigitalOcean
```yaml
# example.yaml
name: demo-cluster-yaml
provider: do
region: blr1
nodes:
  - type: s-1vcpu-2gb
    count: 2
  - type: s-2vcpu-2gb
    count: 2
```

```bash
./ksctl create cluster --config example.yaml # create cluster
kubectl get node --kubeconfig <cluster_name>-kubeconfig # show nodes to new k8s cluster
./ksctl list --provider do # list all cluster for give provider
./ksctl delete --config example.yaml # delete cluster
```

### Progress

- [X] DigitalOcean Deploy cluster 
- [ ] AWS EKS Deploy Cluster
- [ ] Azure EKS Deploy Cluster
- [ ] GCP GKE Deploy Cluster
- [ ] Linode Deploy Cluster
- [ ] Update Clusters
- [ ] Deploying Databases
