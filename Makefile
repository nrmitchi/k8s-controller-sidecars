IMAGE=lemonadehq/controller-sidecars

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
	docker build -t ${IMAGE} .

push:
	docker push ${IMAGE}