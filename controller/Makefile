.PHONY: manifests generate build image run all

all: manifests generate build image run

manifests:
	controller-gen rbac:roleName=conjob-manager-role crd paths="./..." output:dir=manifests/
	mv manifests/role.yaml manifests/0002-role.yml
	mv manifests/kubecon.appliedcomputing.io_conjobs.yaml manifests/0001-conjob-crd.yml

generate:
	controller-gen object paths="./..."

build: manifests generate
	CGO_ENABLED=0 go build -trimpath -o conjob-ctrl ./ctrl.go

image:
	docker build -t localhost:5000/conjob-controller:latest . && \
	docker push localhost:5000/conjob-controller:latest

run:
	kubectl apply -f manifests
