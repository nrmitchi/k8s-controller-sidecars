FROM golang:1.10 AS build
WORKDIR /go/src/github.com/lemonade-hq/k8s-controller-sidecars
RUN git clone --single-branch -b master \
  https://github.com/lemonade-hq/k8s-controller-sidecars.git /go/src/github.com/lemonade-hq/k8s-controller-sidecars/
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o main .

RUN apt-get update && apt-get install -y upx
RUN upx main

RUN mkdir -p /empty

FROM scratch
COPY --from=build /go/src/github.com/lemonade-hq/k8s-controller-sidecars/main /
COPY --from=build /empty /tmp
CMD ["/main"]
