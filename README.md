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

### Progress

- [X] DigitalOcean Deploy cluster 
- [ ] AWS EKS Deploy Cluster
- [ ] Azure EKS Deploy Cluster
- [ ] GCP GKE Deploy Cluster
- [ ] Linode Deploy Cluster
- [ ] Update Clusters
- [ ] Deploying Databases
