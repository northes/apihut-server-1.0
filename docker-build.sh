NAME=ccr.ccs.tencentyun.com/apihut/server
TAG=1.0.2
docker build -t ${NAME}:${TAG} .
docker push ${NAME}:${TAG}
#minikube load ${NAME}:${TAG}