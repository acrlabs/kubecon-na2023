.PHONY: manifests generate build image run all

all: manifests generate build image run

manifests:
	controller-gen rbac:roleName=conjob-manager-role crd paths="./..." output:dir=manifests/

generate:
	controller-gen object paths="./..."

build: manifests generate
	CGO_ENABLED=0 go build -trimpath -o conjob-ctrl ./ctrl.go

image:
	docker build -t localhost:5000/conjob-controller:latest . && \
	docker push localhost:5000/conjob-controller:latest

run:
	kubectl apply -f manifests
