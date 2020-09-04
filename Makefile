IMAGE=nrmitchi/k8s-controller-sidecars:dev

build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

docker:
	docker build -t ${IMAGE} --build-arg BRANCH=dev -f dev/Dockerfile .
	docker push ${IMAGE}
