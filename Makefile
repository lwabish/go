SHIP_TARGET="root@infra2.t:/root/wbw/lwabish"
REGISTRY="ccr.ccs.tencentyun.com/lwabish/lwabish"
IMAGE=$(REGISTRY):latest
KUBE_CONTEXT="infra"
KUBE_NS="wbw"

default: install-mac gen-doc

build-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/lwabish-linux-amd64

build-mac:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/lwabish-darwin-amd64

install-mac: build-mac
	@mv -f bin/lwabish-darwin-amd64 ${GOPATH}/bin/lwabish

install-linux: build-linux
	@mv -f bin/lwabish-linux-amd64 ${GOPATH}/bin/lwabish

gen-doc:
	@go run main.go -g

ship: build-linux
	@scp bin/lwabish-linux-amd64 $(SHIP_TARGET)

image: build-linux
	@docker build -t $(IMAGE) .
	@docker push $(IMAGE)
	@docker rmi $(IMAGE)

install-kube: image
	kubectl config use-context $(KUBE_CONTEXT)
	helm upgrade -i -n $(KUBE_NS) lwabish ./chart --set image=$(IMAGE)

lint:
	goimports-reviser -format -rm-unused ./...