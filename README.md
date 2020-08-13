# loggerheading

Log the headers received by a service for debugging purposes via Istio traffic mirroring. This assumes the K8s cluster you're using already has Istio installed.

#### set up test service

```
kubectl create ns wai-1

kubectl label namespace wai-1 istio-injection=enabled --overwrite

git clone https://github.com/theemadnes/gke-whereami.git

kubectl apply -n wai-1 -f gke-whereami/k8s/configmap.yaml

kubectl apply -n wai-1 -f gke-whereami/k8s/ksa.yaml

kubectl apply -n wai-1 -f gke-whereami/k8s/deployment.yaml

kubectl apply -f wai-1/

curl --header "Host: wai-1.example.com" http://34.72.250.238
```

#### build loggerheading container

```
pack build --builder gcr.io/buildpacks/builder:v1 --publish gcr.io/${PROJECT_ID}/loggerheading
```

#### deploy loggerheading as a service

```
kubectl apply -f k8s/

kubectl get pods -n wai-1 --selector=mirror=target

watch -n 2 'curl --header "Host: wai-1.example.com" http://34.72.250.238'

kubectl logs -n wai-1 -l mirror=target -c loggerheading -f
```