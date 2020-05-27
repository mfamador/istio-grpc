# GitOps workflow demo to install Istion and gRPC services

## Install k3d without Taefik

MacOS
```
brew install k3d
```

[Other OS](https://github.com/rancher/k3d)

### Create a k3d cluster with 4 workers
```
k3d create --server-arg --no-deploy --server-arg traefik --workers 4

export KUBECONFIG="$(k3d get-kubeconfig --name='k3s-default')"
```
### If you need another worker
```
k3d add-node
```

## Install Istio 1.6

[Install Istio 1.6](https://istio.io/docs/setup/install/istioctl/)

```
istioctl install
```

```
kubectl label namespace default istio-injection=enabled  
 ```

```
docker build --tag grpc-greeter-go-server server   
docker build --tag grpc-greeter-go-client client   

docker push marcoamador/grpc-greeter-go-server:1
docker push marcoamador/grpc-greeter-go-client:1
```

```
kubectl apply -f manifests/server
kubectl apply -f manifests/client
```

```
kubectl logs -f deploy/greeting
```

```
kubectl logs -f deploy/greeting-client
```


