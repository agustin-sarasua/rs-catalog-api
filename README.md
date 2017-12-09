# rs-catalog-api

Local Minikube
$ minikube start --vm-driver="virtualbox" --insecure-registry="$REG_IP":80
$ minikube dashboard

$ docker build -t gcr.io/[PROJECT ID]/rs-catalog-api:v1 .
$ kubectl create -f deployment.yaml
$ kubectl create -f service.yaml
$ echo $(minikube service rs-catalog-api-service --url)
$ curl $(minikube service rs-catalog-api-service --url)/catalog/UY/MVD

Cloud

$ docker build -t gcr.io/[PROJECT ID]/rs-catalog-api .
$ gcloud docker -- push gcr.io/[PROJECT ID]/rs-catalog-api:latest
