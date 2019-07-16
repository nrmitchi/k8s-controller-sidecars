export TAG ?= 1.0.0
IMAGE=registry.riskxint.com/library/controller-sidecars

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

docker:
	docker build -t ${IMAGE}:${TAG} -f Dockerfile .
	docker push ${IMAGE}:${TAG}
