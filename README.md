# go-chart
## This is a simple go application that uses helm to deploy two services in kubernetes
Required tools:
- minikube
- helm
- kubectl
- docker
- go

## start minikube

```bash
minikube start
```

## create namespace

```bash
kubectl create namespace go-ns
```

## install helm

```bash
helm install go-chart ./my-chart --namespace go-ns
```



## go build service-a and service-b

```bash
cd service-a
go build -o service-a main.go
cd ../service-b
go build -o service-b main.go

```

## build docker images with minikube docker:

```bash
eval $(minikube docker-env)
cd microservice-arch
docker build -t service-a:1.0.0 ./service-a
docker build -t service-b:1.0.0 ./service-b
```

## if image is already built using docker, copy already built images to minikube docker:

```bash
docker build -t service-a:1.0.0 -f service-a/Dockerfile .
docker build -t service-b:1.0.0 -f service-b/Dockerfile .

minikube image load service-a:1.0.0
minikube image load service-b:1.0.0
```

## deploy service-a and service-b

```bash
helm upgrade -i go-chart ./my-chart --namespace go-ns --create-namespace
```


## run using docker-compose

```bash
docker-compose up
```